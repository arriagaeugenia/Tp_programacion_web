#!/bin/bash

#Ingresar valores por consola
#read -p "Ingresa el titulo de la obra" titulo
#read -p "Ingresa la descripción" desc
#read -p "Ingresa el artista" artista
#read -p "Ingresa el precio" precio


# Crear un nuevo producto
curl -X POST http://localhost:8080/obras \
     -H "Content-Type: application/json" \
     -d '{"Titulo": "a", "Descripcion": "a", "Artista": "a", "Precio": 2.0, "Vendida": false}'
