package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type nodo struct {
	nombre string
	apellido string
	apodo string
	favoritos string
	Siguiente *nodo
	Anterior *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	var cadena strings.Builder
	fmt.Fprintf(&cadena, "digraph Tarea3{\n")
	fmt.Fprintf(&cadena, "node[shape=record];\n")
	Grafico(&cadena, &li)
	fmt.Println("fin")
}

func Grafico(s *strings.Builder, lista *lista) {
	imp := lista.cabeza
	for imp != nil {
		siguiente := imp.Siguiente
		anterior := imp.Anterior
		fmt.Fprintf(s, "node%p[label=\"{{%v|%v}|%v}|%v\"];\n", &(*imp), imp.nombre, imp.apellido, imp.apodo, imp.favoritos)
		if imp == lista.cabeza {
			if imp.Siguiente!=nil {
				fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*siguiente))
			}
		}else if imp.Siguiente !=nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*anterior))
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*siguiente))
		}else if imp.Siguiente == nil {
			fmt.Fprintf(s, "node%p->node%p;\n", &(*imp), &(*anterior))
		}
		imp = imp.Siguiente
	}
	fmt.Fprintf(s, "}")
	guardarArchivo(s.String(), "Tarea3.dot")
}

func guardarArchivo(cadena string, nombreArchivo string) {
	f, err := os.Create(nombreArchivo)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(cadena)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpdf", "./"+nombreArchivo).Output()
	mode := int(0777)
	ioutil.WriteFile(nombreArchivo+".pdf", cmd, os.FileMode(mode))
}