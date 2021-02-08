package listas

type Nodo struct {
	OrigenM string
	DestinoM string
	CuerpoM string
	Siguiente *Nodo
	Anterior *Nodo
}

type Datos struct{
	General []Mensajes `json:"Mensajes"`
}

type Mensajes struct {
	Origen string `json:"Origen"`
	Destino string `json:"Destino"`
	Msg []Contenido `json:"Msg"`
}

type Contenido struct {
	Fecha  string `json:"Fecha"`
	Texto string `json:"Texto"`
}

func (L *Lista) Insertar(nuevo *Nodo) string{
	if L.Cabeza == nil{
		L.Cabeza = nuevo
		L.Final = nuevo
	}else{
		L.Final.Siguiente = nuevo
		nuevo.Anterior = L.Final
		L.Final = nuevo
	}
	return ""
}

type Lista struct {
	Cabeza *Nodo
	Final *Nodo
}

func (L *Lista) Imprimir() string{
	var aux string
	imp := L.Cabeza
	for imp != nil{
		//aux += imp.CuerpoM+"\n"
		return aux
		imp = imp.Siguiente
	}
	return ""
}

func (L *Nodo) RetornarNodo() string{
	return "Origen: "+L.OrigenM+"\t Destino: "+L.DestinoM+"\t Cuerpo: "+L.CuerpoM+"\n"
}

func (L *Lista) Retornar() string {
	var cadena string
	aux := L.Cabeza
	for aux != nil {
		cadena += aux.RetornarNodo()+"\n"
		aux = aux.Siguiente
	}
	return cadena
}