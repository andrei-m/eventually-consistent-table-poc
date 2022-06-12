package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type server struct {
	autoDB AutoDB
}

func (s *server) handlePostRandomizedAuto(c *gin.Context) {
	auto := s.autoDB.NewRandomizedAuto()
	c.JSON(http.StatusOK, auto)
}

func (s *server) handleGetAutos(c *gin.Context) {
	c.JSON(http.StatusOK, s.autoDB.GetAutos())
}

func GetRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})
	s := server{}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	r.GET("/autos", s.handleGetAutos)
	r.POST("/randomized_auto", s.handlePostRandomizedAuto)
	return r
}
