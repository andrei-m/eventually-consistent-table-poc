package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type server struct {
	autoDB AutoDB
	saleDB SaleDB
}

func (s *server) handlePostRandomizedAuto(c *gin.Context) {
	auto := s.autoDB.NewRandomizedAuto()
	c.JSON(http.StatusOK, auto)
}

func (s *server) handleGetAutos(c *gin.Context) {
	c.JSON(http.StatusOK, s.autoDB.GetAutos())
}

func (s *server) handlePostRandomizedSale(c *gin.Context) {
	autoIdStr, ok := c.GetPostForm("autoId")
	if !ok {
		c.JSON(http.StatusBadRequest, "autoId is required in POST params")
		return
	}
	autoID, err := strconv.Atoi(autoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("autoId must be an integer; got %s", autoIdStr))
		return
	}
	if !s.autoDB.AutoIDExists(autoID) {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("autoId %d does not exist", autoID))
		return
	}
	sale := s.saleDB.NewRandomizedSale(autoID)
	c.JSON(http.StatusOK, sale)
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
	r.POST("/randomized_sale", s.handlePostRandomizedSale)
	return r
}
