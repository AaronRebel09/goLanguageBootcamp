package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//Esta función nos permite ver la anatomía de un mensaje Request de una
func Ejemplo(context *gin.Context) {
	//El body, header y method están contenidos en el contexto de gin.
	contenido := context.Request.Body
	header := context.Request.Header
	metodo := context.Request.Method
	fmt.Println("¡He recibido algo!")
	fmt.Printf("\tMetodo: %s\n", metodo)
	fmt.Printf("\tContenido:\n")
	for key, value := range header {
		fmt.Printf("\t\t%s -> %s\n", key, value)
	}
	fmt.Printf("\tCotenido:%s\n", contenido)
	fmt.Println("¡Yay!")
	context.String(200, "¡Lo recibí!") //Respondemos al cliente con 200 OK y un mensaje.
}

func main() {
	// Creates a gin router with default middleware
	router := gin.Default()
	// A handler for GET request on /example
	router.GET("/hello-world", func(c *gin.Context) {
		// get context
		Ejemplo( c )
		c.JSON(200, gin.H{
			"message": "Hello World!",
		}) // gin.H is a shortcut for map[string]interface{}
	})

	//Cada vez que llamamos a GET y le pasamos una ruta, definimos un nuevo endpoint.
	router.GET("/", HandlerRaiz)
	router.GET("/gophers", HandlerGophers)
	router.GET("/gophers/get", HandlerGetGopher)
	router.GET("/gophers/info", HandlerGetInfo)
	router.GET("/about", HandlerAbout)

	server := gin.Default()
	server.GET("/", HandlerRaiz)
	//Ahora podemos atender peticiones a /gophers/, /gophers/get o /gophers/info de una
	//forma más elegante.
	gopher := server.Group("/gophers")
	{
		gopher.GET("/", HandlerGophers)
		gopher.GET("/get", HandlerGetGopher)
		gopher.GET("/info", HandlerGetInfo)
	}
	server.GET("/about", HandlerAbout)
	server.Run(":8081")

	router.Run() // listen and serve on port 8080
}

func HandlerGetGopher(context *gin.Context) {
	
}


func HandlerGetInfo(context *gin.Context) {
	
}

func HandlerAbout(context *gin.Context) {
	
}

func HandlerGophers(context *gin.Context) {
	
}

func HandlerRaiz(context *gin.Context) {
	
}
