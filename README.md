# TPE Parte 3: Logica de negocio

Este módulo implementa la **logica de negocio** de la aplicación conectando la base de datos con el servidor.

---

## Estructura

```

base_de_datos/
├── db/               # Definiciones de queries y modelos generados
│   ├── queries/      # Consultas SQL
│   ├── schema/       # Esquema de la base de datos
│   └── sqlc/         # Código Go generado automáticamente
servidor/
├── html/             # Archivos estáticos 
│   ├── index.html
│   └── Imagenes/
└── main.go           # Código principal del servidor HTTP
go.mod            
go.sum
Makefile              # Archivo para levantar, testear y bajar la pagina
request.hurl          # Herramienta Hurl para testear los endpoints de la pagina.

```

---
## Requisitos
### Docker para linux:
Tener instalado Docker desde la pagina: https://docs.docker.com/compose/install/
O en la terminal, instalarlo con el comando: 

```bash
sudo apt update && sudo apt install docker.io docker-compose -y
```
Esta instalacion funciona para Ubuntu y Debian

### Hurl para linux:
    Abrir la terminal y ejecutar:

```bash
$ INSTALL_DIR=/tmp
$ VERSION=7.0.0
$ curl --silent --location https://github.com/Orange-OpenSource/hurl/releases/download/$VERSION/hurl-$VERSION-x86_64-unknown-linux-gnu.tar.gz | tar xvz -C $INSTALL_DIR
$ export PATH=$INSTALL_DIR/hurl-$VERSION-x86_64-unknown-linux-gnu/bin:$PATH
```

## Uso


### 1. Ejecutar el servidor y la base de datos
Desde la carpeta `TP_PROGRAMACION_WEB/` correr: 

```bash
make up
```
---

### 2. Testear funcionamiento de endpoints

```bash
make test
```
---

### 3. Bajar el servidor y la base de datos

```bash
make down
```
---

## Acceder a la aplicación
Abrir en el navegador:  http://localhost:8080


## Integrantes
- Milagros Lopez
- Maria Eugenia Arriaga
- Bianca Rizzalli

