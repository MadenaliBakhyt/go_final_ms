package main

import (
	"context"
	"e-learning/internal/data"
	"e-learning/internal/validator"
	pb "e-learning/pb/courses-service"
	"e-learning/pkg/client"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (app *application) addCourses(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name    string `json:"name"`
		Credits int64  `json:"credits"`
		Price   int64  `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	course := &data.Course{
		Name:    input.Name,
		Credits: input.Credits,
		Price:   input.Price,
	}
	v := validator.New()

	if data.ValidateCourse(v, course); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	createCourse, err := client.GetCoursesClient().CreateCourse(context.Background(),
		&pb.Course{Name: course.Name, Credits: course.Credits, Price: course.Price})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/courses/%d", createCourse.Id))
	err = app.writeJSON(w, http.StatusCreated, createCourse, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getCourse(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		app.logError(r, err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		app.logError(r, err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"courses", // name
		false,     // durable
		true,      // delete when unused
		false,     // exclusive
		false,     // noWait
		nil,       // arguments
	)
	if err != nil {
		app.logError(r, err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		app.logError(r, err)
	}

	bytes, err := proto.Marshal(&pb.CourseId{
		Id: id,
	})
	if err != nil {
		app.logError(r, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",        // exchange
		"courses", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        bytes,
		})
	if err != nil {
		app.logError(r, err)
	}

	for d := range msgs {
		res := &pb.Course{}

		err = proto.Unmarshal(d.Body, res)
		if err != nil {
			app.logError(r, err)
		}
		err = app.writeJSON(w, http.StatusOK, res, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}
		return
	}
}

func (app *application) deleteCourse(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	_, err = client.GetCoursesClient().DeleteCourse(context.Background(), &pb.CourseId{Id: id})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"response": "Course deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := client.GetCoursesClient().GetCourses(context.Background(), &pb.GetCoursesRequest{})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, courses, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
