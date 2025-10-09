// package basededatos
package main

import (
	"context"      // para context.Background()
	"database/sql" // para sql.Open()
	"fmt"          // para fmt.Println()
	"log"
	"strconv"

	sqlc "galeriadearte.com/base_de_datos/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "postgresql://milibianeuge:programacionweb@localhost:5432/db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}
	defer db.Close()
	queries := sqlc.New(db)
	ctx := context.Background()

	createdObra, err := queries.CreateObra(ctx, // Create
		sqlc.CreateObraParams{
			Titulo:      "La noche estrellada",
			Descripcion: sql.NullString{String: "Una pintura famosa", Valid: true},
			Artista:     "Vincent van Gogh",
			Precio:      strconv.Itoa(1000000),
			Vendida:     sql.NullBool{Bool: false, Valid: true},
		})
	if err != nil {
		log.Fatalf("failed to create obra: %v", err)
	}
	fmt.Printf("Created obra: %+v\n", createdObra)

	obra, err := queries.GetObraById(ctx, createdObra.ID) // Read One
	if err != nil {
		log.Fatalf("failed to get obra: %v", err)
	}
	fmt.Printf("Retrieved obra: %+v\n", obra)

	//Actualiza la obra de "La noche estrellada"
	err = queries.UpdateObra(ctx, sqlc.UpdateObraParams{ // Update
		ID:          1,
		Titulo:      "La noche estrellada!",
		Descripcion: sql.NullString{String: "Una pintura famosa y muy hermosa", Valid: true},
		Artista:     "Vincent van Gogh",
		Precio:      "1200000",
		Vendida:     sql.NullBool{Bool: false, Valid: true},
	})
	if err != nil {
		log.Fatalf("failed to update obra: %v", err)
	}
	fmt.Println("Obra actualizada")
}
