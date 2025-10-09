# TPE Parte 2: Base de Datos

Este módulo gestiona la **base de datos** de la aplicación.  
Utiliza **sqlc** para generar código Go a partir de consultas SQL.


## Estructura

```
base_de_datos/
├── db/               # Definiciones de queries y modelos generados
│   ├── queries/      # Consultas SQL
│   ├── schema/       # Esquema de la base de datos
│   └── sqlc/         # Código Go generado automáticamente
├── go.mod            # Dependencias del módulo
├── go.sum
├── main.go           # Entrada principal para probar conexión/queries
└── sqlc.yaml         # Configuración de sqlc
```

---
## Requisitos

Tener instalado Docker desde la pagina: [text](https://docs.docker.com/compose/install/)
O en la terminal, instalarlo con el comando: 

```bash
sudo apt update && sudo apt install docker.io docker-compose -y
```
Esta instalacion funciona para Ubuntu y Debian

## Uso

### 1. Configurar la base de datos

Para crear la base de datos y el usuario en Docker se debe ejecutar el siguiente comando:

```bash
docker compose up --build
```

### 2. Probar el módulo

Ejecutar: 

```bash
go run main.go
```

---

## Notas

- Los modelos y funciones en go se encuentran en `db/sqlc`.  
- Los queries personalizados están en `db/queries/`. 
- Solo es necesario volver a ejecutar `sqlc generate` si se modifican las consultas (`db/queries/`) o el esquema (`db/schema/`). 

---
