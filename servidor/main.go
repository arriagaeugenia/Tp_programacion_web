package main

import (
	"fmt"
	"net/http"
)

func main() {
	staticDir := "./html"
	// Servidor de archivos
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	// Puerto
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)

	// Levantar servidor
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}
func inicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos
	fmt.Print(w, "./html")
}
