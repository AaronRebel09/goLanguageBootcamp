package handlers

import (
	"github.com/gin-gonic/gin"
)

//Definimos una paseudobase de datos donde consultaremos la información.
var empleados = map[string]string{
	"644" : "Empleado A",
	"755" : "Empleado B",
	"777" : "Empleado C",
}

//Definimos nuestra estructura de información
type Empleado struct {
	// Una etiqueta de struct se cierra con caracteres de acento grave `
	Nombre string `form:"name" json:"name"`
	Id string `form:"id" json:"id"`
	Activo string `form:"active" json:"activa" bidding:"requiere"`
}

//Este handler se encargará de responder a /.
func PaginaPrincipal(ctxt *gin.Context) {
	ctxt.String(200, "¡Bienvenido a la Empresa Gophers!")
}

//Este handler verificará si la id que pasa el cliente existe en nuestra base de datos.
func BuscarEmpleado(ctxt *gin.Context){
	empleado, existe := empleados[ctxt.Param("id")]
	if existe {
		ctxt.String(200,"Información del empleado %s, nombre: %s",ctxt.Param("id"),empleado)
	} else{
		ctxt.String(404,"Información del empleado ¡No existe!")
	}
}

/*func BuscarEmpleado(ctxt *gin.Context) {
	// ShouldBind verifica el Content-Type para seleccionar un
	// mecanismo de binding (vinculación) de forma automática.
	var empleado Empleado
	//Nuestro objetivo aquí es asignar los campos de nuestra estructura con los datos que recibimos del request.
	if ctxt.Bind(&empleado) == nil {
		log.Println("====== Bind Por Query String ======")
		log.Println(empleado.Id)
		log.Println(empleado.Nombre)
		ctxt.String(200, "(Query String) - Empleado: %s, Id: %s\n", empleado.Nombre, empleado.Id)
	}
	//Por query string como arriba (es decir, form) o por JSON.
	if ctxt.BindJSON(&empleado) == nil {
		log.Println("====== Bind Por JSON ======")
		log.Println(empleado.Id)
		log.Println(empleado.Nombre)
		ctxt.String(200, "(Query JSON) - Empleado: %s, Id: %s\n", empleado.Nombre, empleado.Id)
	}
}*/

/*//Esta función solo mostrará aquellos empleados activos o inactivos, dependiente del parámetro active.
func FiltrarEmpleados(ctxt *gin.Context) {
	empleados := GenerarListaEmpleados()
	var filtrados []*Gopher
	cantidadFiltrados := 0
	for _, e := range empleados {
		if ctxt.Query("active") == e.Activo {
			filtrados = append(filtrados, e)
			cantidadFiltrados++
		}
	}
	if cantidadFiltrados > 0 {
		ctxt.String(200, "¡Filtrado exitoso!\n")
		for _, e := range filtrados {
			log.Println("====== Filter Por Query String ======")
			log.Println(e.Id)
			log.Println(e.Nombre)
			log.Println(e.Activo)
		}
	} else {
		ctxt.String(404, "¡Oh no! No se pudo hacer el filtrado\n")
	}
}

func GenerarListaEmpleados() interface{} {

}*/
