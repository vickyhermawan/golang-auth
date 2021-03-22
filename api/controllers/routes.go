package controllers

import "github.com/vickyhermawan/rest-api/api/middlewares"

func (s *Server) initializeRoutes() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	//Auth Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")

}
