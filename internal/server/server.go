package server

import "github.com/labstack/echo"

type ServerStruct struct {
	engine *echo.Echo
}

func (s *ServerStruct) StartServer() {
	s.engine.Start("localhost:8080")
}

func NewHTTPServer() *ServerStruct {
	e := echo.New()

	return &ServerStruct{
		engine: e,
	}
}
