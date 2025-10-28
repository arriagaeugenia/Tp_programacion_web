# requests.hurl

# 1. Obtener todos las obras
GET http://localhost:8080/obras
HTTP 200


# 2. Obtener la obra con id 8
GET http://localhost:8080/obras/8
HTTP 200
[Asserts]
jsonpath "$.id" == 8

# 3. Crear una nueva obra
POST http://localhost:8080/obras
Content-Type: application/json

{
  "titulo": "Guernica4",
  "descripcion": "Obra famosa de Picasso",
  "artista": "Picasso",
  "precio": "100.0",
  "vendida": "false"
}

HTTP 201
[Asserts]
jsonpath "$.id" isInteger
jsonpath "$.titulo" == "Guernica4"

[Captures]
obra_id: jsonpath "$.id"

# 4. Update obra creada
PUT http://localhost:8080/obras
Content-Type: application/json
{
  "id": {{obra_id}},
  "descripcion": "Una obra muy famosa del estilo cubismo pintada por Picasso"
}
HTTP 200

# 5. Obtener obra actualizada para verificar la actualizacion
GET http://localhost:8080/obras/{{obra_id}}
HTTP 200
[Asserts]
jsonpath "$.id" == {{obra_id}}

# 6. Eliminar la obra creada
DELETE http://localhost:8080/obras/{{obra_id}}
HTTP 200

# 7. Chequear que se borro la obra
GET http://localhost:8080/obras/{{obra_id}}
HTTP 404