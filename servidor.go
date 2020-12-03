package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)
//// mapa de materias ////
var materias = make(map[string]map[string]float64)
////  mapa alumnos ////
var alumnos = make(map[string]map[string]float64)

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}
func agregar(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		//// alumno - calificaion
		alumno := make(map[string]float64)
		//// materia - calificacion
		materia := make(map[string]float64)
		cal, err := strconv.ParseFloat(req.FormValue("promedio"), 64)
		if err != nil {
			fmt.Println(err)
		}
		cont :=0
		al := 0
		nom:=req.FormValue("nombre")
		mat := req.FormValue("materia")
		fmt.Println(req.FormValue("nombre"),req.FormValue("materia"),req.FormValue("promedio"))
		alumno[nom] = cal
		materia[mat] = cal
		for i,k := range alumnos{
			if i == nom{
				al = 1
				for i,_ := range k{
					if mat == i{
						fmt.Println("calificacion ya registrada")
						return
					}
				}
			}
		}
		for i,_ := range materias{
			if i == mat{
				cont = 1
				materias[mat][nom]=cal
			}
		}
		if cont == 0{		
			materias[mat] = alumno
		}
		if al == 0{
			alumnos[nom]= materia
		}else{
			alumnos[nom][mat]= cal
		} 
		fmt.Println(alumnos,materias)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			nom,mat,cal,
		)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("agregar.html"),
		)
	}

}
func promedio_alumno(res http.ResponseWriter, req *http.Request)  {
	switch req.Method {
	case "POST":
		var prom,num float64
		num = 0.0
		prom = 0.0
		nombre := req.FormValue("nombre")
		for i,mat:=range alumnos{
			if nombre == i{
				for _,cal:= range mat{
					prom += cal
					num +=1
				}
			}
		}
		prom = prom/num
		fmt.Println("Premedio Alumno",prom)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedio.html"),
			prom,
		)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioAlumno.html"),
		)
	}
}
func promedio_materia(res http.ResponseWriter, req *http.Request)  {
	switch req.Method {
	case "POST":
		var prom,num float64
	num = 0.0
	prom = 0.0
	materia:= req.FormValue("materia")
	for i,al:=range materias{
		if i == materia{
			for _,cal:= range al{
				prom += cal
				num +=1
			}
		}
	}
	prom = prom/num
	fmt.Println("Promedio Materia",prom)
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("promedio.html"),
		prom,
	)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioMateria.html"),
		)
	}
	
}
func promedio_general(res http.ResponseWriter, req *http.Request)  {
	var prom,num float64
	num = 0.0
	prom = 0.0
	for _,mat:=range alumnos{
		for _,cal:= range mat{
			prom += cal
			num +=1
		}
	}
	prom = prom/num
	fmt.Println("Promedio General ",prom)
	res.Header().Set(
		"Content-Type",
		"text/html"   ,
	)
	fmt.Fprintf(
		res,
		cargarHtml("promedio.html"),
		prom,
	)
}
func root(res http.ResponseWriter, req *http.Request) {
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("index.html"),
		)
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/agregar", agregar)
	http.HandleFunc("/promedio_alumno", promedio_alumno)
	http.HandleFunc("/promedio_materia",promedio_materia)
	http.HandleFunc("/promedio_general",promedio_general)
	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)
}
