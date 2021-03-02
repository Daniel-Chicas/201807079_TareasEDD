package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type NodoArbol struct{
	Num int
	Factor int
	Izq *NodoArbol
	Der *NodoArbol
}

type Arbol struct{
	raiz *NodoArbol
}

func NuevoArbol() *Arbol{
	return &Arbol{nil}
}

func NuevoNodo (valor int) *NodoArbol{
	return &NodoArbol{valor, 0, nil, nil}
}
/*
func RotacionII(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n.Izq = n1.Der
	n1.Der = n
	if n1.Factor == -1 {
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func RotacionDD(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n.Der = n1.Izq
	n1.Izq = n
	if n1.Factor == 1{
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func RotacionDI(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n2 := n1.Izq
	n.Der = n2.Izq
	n2.Izq = n
	n1.Izq = n2.Der
	n2.Der = n1
	if n2.Factor == 1 {
		n.Factor = -1
	}else{
		n.Factor = 0
	}
	if n2.Factor == -1{
		n1.Factor = 1
	}else{
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func RotacionID(n *NodoArbol, n1 *NodoArbol) *NodoArbol{
	n2 := n1.Der
	n.Izq = n2.Der
	n2.Der = n
	n1.Der = n2.Izq
	n2.Izq = n1
	if n2.Factor == 1 {
		n1.Factor = -1
	}else{
		n1.Factor = 0
	}
	if n2.Factor == -1 {
		n.Factor = 1
	}else{
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}
 */

 func insertar(ra *NodoArbol, fact int, ya *bool) *NodoArbol{
	//var n1 *NodoArbol
	 if ra == nil {
		 ra = NuevoNodo(fact)
		 *ya = true
	 }else if fact<ra.Num{
	 	izq := insertar(ra.Izq, fact, ya)
	 	ra.Izq = izq
		 if *ya == true {
			 switch ra.Factor {
				case 1:
			 		ra.Factor = 0
			 		*ya = false
			 		break
			 	case 0:
			 		ra.Factor = -1
			 		break
			 	case -1:
			 		/*
			 		n1 = ra.Izq
					if n1.Factor==-1 {
						ra = RotacionII(ra, n1)
					}else{
						ra = RotacionID(ra, n1)
					}
			 		 */
					*ya = false
			 }
		 }
	 }else if fact > ra.Num{
	 	der := insertar(ra.Der, fact, ya)
	 	ra.Der = der
	 	if *ya{
			switch ra.Factor {
				case 1:
					/*
					n1 = ra.Der
					if n1.Factor == 1 {
						ra = RotacionDD(ra, n1)
					}else{
						ra = RotacionDI(ra, n1)
					}
					 */
					*ya = false
					break
				case 0:
					ra.Factor = 1
					break
				case -1:
					ra.Factor = 0
					*ya = false
			}
		}
	 }
	 return ra
 }

func (this *Arbol) Insertar(val int){
	b := false
	a := &b
	this.raiz = insertar(this.raiz, val, a)
}

func main(){
	a := NuevoArbol()
	for i := 0; i < 10; i++ {
		a.Insertar(rand.Intn(10))
	}
	fmt.Println("PREORDEN:")
	PreOrden(a.raiz)
	fmt.Println("\n")
	fmt.Println("INORDEN")
	inOrden(a.raiz)
	fmt.Println("\n")
	fmt.Println("POSORDEN")
	posOrden(a.raiz)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func PreOrden(raiz *NodoArbol) {
	if raiz!=nil {
		fmt.Print(strconv.Itoa(raiz.Num)+",")
		PreOrden(raiz.Izq)
		PreOrden(raiz.Der)
	}
}

func inOrden(raiz *NodoArbol) {
	if raiz!=nil {
		inOrden(raiz.Izq)
		fmt.Print(strconv.Itoa(raiz.Num)+",")
		inOrden(raiz.Der)
	}
}

func posOrden(raiz *NodoArbol) {
	if raiz!=nil {
		posOrden(raiz.Izq)
		posOrden(raiz.Der)
		fmt.Print(strconv.Itoa(raiz.Num)+",")
	}
}