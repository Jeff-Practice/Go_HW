package main

import (
	"hw/control"
	"net/http"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
)

func main() {
	context.Background()
	router := gin.Default()
	http.HandleFunc("ss", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	router.GET("/role", control.Get)

	router.GET("/role/:id", control.GetOne)

	router.POST("/role", control.Post)

	router.PUT("/role/:id", control.Put)

	router.DELETE("/role/:id", control.Delete)

	router.Run(":8080")
}
