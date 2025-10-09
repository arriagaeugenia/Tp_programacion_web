# TPE Parte 1: Servidor Web

Este módulo implementa el **servidor HTTP** de la aplicación.  
Se encarga de servir la interfaz web.

---

## Estructura

```
servidor/
├── html/             # Archivos estáticos 
│   ├── index.html
│   └── Imagenes/
├── go.mod            
├── go.sum
└── main.go           # Código principal del servidor HTTP
```

---

## Uso

### 1. Ejecutar el servidor
Desde la carpeta `servidor/` correr: 

```bash
go run main.go
```

### 2. Acceder a la aplicación
Abrir en el navegador:  [http://localhost:8080]

---

## Interfaz

- La página principal se encuentra en `html/index.html`.  
- Los recursos (imágenes) están en `html/Imagenes/`.

## Integrantes
- Milagros Lopez
- Maria Eugenia Arriaga
- Bianca Rizzalli

