package main
import(
	//"encoding/json"
	//"fmt"
	//"go/types"
	//"io/ioutil"
	//"log"
)
type almacen struct{
	datos []dato
}

type dato struct {
	indice string `json: "Indice"`
	departamentos []departamento
}

type departamento struct {
	nombreDepartamento string `json: "Nombre"`
	tiendas []tienda
}

type tienda struct {
	nombreTienda string `json: "Nombre"`
	descripcion string `json: "Descripcion"`
	contacto string `json: "Contacto"`
	calificacion string `json: "Calificacion"`
}
