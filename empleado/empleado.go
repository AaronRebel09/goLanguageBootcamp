package main

import (
	"github.com/AaronRebel09/go-simple-module/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.GET("/",handlers.PaginaPrincipal)
	server.GET("/empleados/:id",handlers.BuscarEmpleado)
	server.Run(":8080")

}
