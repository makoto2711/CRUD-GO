package main

import(
	"net/http"
	"fmt" 
	"text/template"
)


var plantillas= template.Must(template.ParseGlob("plantillas/*"))



func main() {
	http.HandleFunc("/", Inicio)
	http.HandleFunc("/crear", Crear)
	
	fmt.Println("Servidor corriendo")
	
	http.ListenAndServe(":8080", nil)
}

func Inicio(w http.ResponseWriter, r *http.Request) {

	plantillas.ExecuteTemplate(w, "inicio", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {

	plantillas.ExecuteTemplate(w, "crear", nil)
}










