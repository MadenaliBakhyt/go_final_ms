package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/courses", app.getCourses)
	router.HandlerFunc(http.MethodPost, "/courses", app.requirePermission("write", app.addCourses))
	router.HandlerFunc(http.MethodGet, "/courses/:id", app.getCourse)
	router.HandlerFunc(http.MethodDelete, "/courses/:id", app.requirePermission("write", app.deleteCourse))
	router.HandlerFunc(http.MethodGet, "/teachers", app.getTeachers)
	router.HandlerFunc(http.MethodPost, "/teachers", app.requirePermission("write", app.addTeacher))
	router.HandlerFunc(http.MethodGet, "/teachers/:id", app.getTeacher)
	router.HandlerFunc(http.MethodDelete, "/teachers/:id", app.requirePermission("write", app.deleteTeacher))
	router.HandlerFunc(http.MethodDelete, "/users/:id", app.requirePermission("write", app.deleteUser))
	router.HandlerFunc(http.MethodPost, "/register", app.registerUser)
	router.HandlerFunc(http.MethodPost, "/authenticate", app.createToken)

	return app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router))))
}
