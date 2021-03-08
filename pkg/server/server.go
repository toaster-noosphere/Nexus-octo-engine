package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	http *http.Server
}

func New(addr string) *Server {
	s := &Server{
		http: &http.Server{
			Addr: addr,
		},
	}
	s.http.Handler = s.setupRouter()

	return s
}

func (s *Server) Run() {
	errs := make(chan error, 1)

	defer func() {
		if err := s.http.Close(); err != nil {
			log.Fatal(fmt.Errorf("server stopped with error: %w", err))
		}
	}()

	go func() {
		log.Printf("server started")
		errs <- s.http.ListenAndServe()
	}()

	err := <-errs
	if err != nil {
		log.Fatal("server exited with error: %w", err)
	}
}

func (s *Server) setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/helloworld", s.helloworld)

	return r
}

func (s *Server) helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
}
