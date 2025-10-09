# requests.hurl

# 1. Obtener todos los productos
GET http://localhost:8080/obras
HTTP 200
[Asserts]
jsonpath "$.length()" == 2
jsonpath "$[0].name" == "Laptop"

# 2. Obtener el producto con id 2
GET http://localhost:8080/obra/2
HTTP 200
[Asserts]
jsonpath "$.name" == "Smartphone"