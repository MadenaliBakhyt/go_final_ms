package main

import (
	"context"
	"e-learning/internal/data"
	"e-learning/internal/validator"
	pb "e-learning/pb/teachers-service"
	"e-learning/pkg/client"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
	"net/http"
	"time"
)

func (app *application) addTeacher(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Subject string `json:"subject"`
		Salary  int64  `json:"salary"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	teacher := &data.Teacher{
		Name:    input.Name,
		Surname: input.Surname,
		Subject: input.Subject,
		Salary:  input.Salary,
	}
	v := validator.New()

	if data.ValidateTeacher(v, teacher); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	createTeacher, err := client.GetTeachersClient().CreateTeacher(context.Background(),
		&pb.Teacher{Name: teacher.Name, Surname: teacher.Surname, Subject: teacher.Subject, Salary: teacher.Salary})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/teachers/%d", createTeacher.Id))
	err = app.writeJSON(w, http.StatusCreated, createTeacher, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) getTeacher(w http.ResponseWriter, r *http.Request) {
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
		"teachers_client", // name
		false,             // durable
		true,              // delete when unused
		false,             // exclusive
		false,             // noWait
		nil,               // arguments
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

	bytes, err := proto.Marshal(&pb.TeacherId{
		Id: id,
	})
	if err != nil {
		app.logError(r, err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",         // exchange
		"teachers", // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			ReplyTo:     q.Name,
			Body:        bytes,
		})
	if err != nil {
		app.logError(r, err)
	}

	for d := range msgs {
		res := &pb.Teacher{}

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

func (app *application) deleteTeacher(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	_, err = client.GetTeachersClient().DeleteTeacher(context.Background(), &pb.TeacherId{Id: id})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"response": "Teacher deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getTeachers(w http.ResponseWriter, r *http.Request) {
	teachers, err := client.GetTeachersClient().GetTeachers(context.Background(), &pb.GetTeachersRequest{})
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, teachers, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
