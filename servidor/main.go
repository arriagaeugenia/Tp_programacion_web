package main

import (
	"database/sql" // para sql.Open()
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	sqlc "galeriadearte.com/base_de_datos/db/sqlc"
)

var dbQueries *sqlc.Queries

func main() {

	//LOGICA DE NEGOCIO
	// Conexión a la base de datos
	connStr := "postgresql://milibianeuge:programacionweb@localhost:5432/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Rutas
	http.HandleFunc("/obras", obrasHandler)
	http.HandleFunc("/obras/", obraHandler)
	http.HandleFunc("/inicio", inicio)

	dbQueries = sqlc.New(db)

	staticDir := "./html"
	// Servidor de archivos
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/", fs)

	// Puerto
	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	fmt.Printf("Sirviendo archivos desde: %s\n", staticDir)

	// Levantar servidor
	err3 := http.ListenAndServe(port, nil)
	if err3 != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}

}
func inicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Carpeta de archivos estáticos
	fmt.Print(w, "./html")
}

func obrasHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getObras(w, r)
	case http.MethodPost:
		createObra(w, r)
	case http.MethodPut:
		updateObra(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func obraHandler(w http.ResponseWriter, r *http.Request) {
	// Extraer ID del path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.Error(w, "Invalid product ID",
			http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		getObra(w, r, id)
	case http.MethodDelete:
		deleteObra(w, r, id)
	default:
		http.Error(w, "Method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func getObras(w http.ResponseWriter, r *http.Request) {
	//llamo a listar obras
	w.Header().Set("Content-Type", "application/json")
	obras, err := dbQueries.ListObras(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir a JSON y enviar respuesta
	json.NewEncoder(w).Encode(obras)
}

func createObra(w http.ResponseWriter, r *http.Request) {
	//creo variable de tipo Obra
	var nuevaObra sqlc.CreateObraParams

	// Decodificar JSON del cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&nuevaObra)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Inserto en la base de datos
	createdObra, err := dbQueries.CreateObra(r.Context(), nuevaObra)
	if err != nil {
		http.Error(w, "Failed to create obra", http.StatusInternalServerError)
		return
	}

	// Enviar respuesta con la obra creada
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdObra)
}

func getObra(w http.ResponseWriter, r *http.Request, id int) {
	obraBuscada, err := dbQueries.GetObraById(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Obra not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obraBuscada)
}

func updateObra(w http.ResponseWriter, r *http.Request) {
	//id de la obra a actualizar se encuentra dentro del Body.
	//creo variable de tipo Obra
	var obraActualizada sqlc.UpdateObraParams

	err := json.NewDecoder(r.Body).Decode(&obraActualizada)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err2 := dbQueries.UpdateObra(r.Context(), obraActualizada)
	if err2 != nil {
		http.Error(w, "Failed to update obra", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Obra actualizada exitosamente")

}

func deleteObra(w http.ResponseWriter, r *http.Request, id int) {
	//borramos la obra
	err := dbQueries.DeleteObra(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "Failed to delete obra", http.StatusInternalServerError)
		return
	}
	// Respuesta de éxito
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Obra eliminada exitosamente")
}
