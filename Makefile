#Levantar el servidor y la base de datos
up:
	@echo "Levantando la base de datos..."
	cd base_de_datos && sudo docker compose up --build -d 

	@echo "Levantando el servidor Go..."
	cd servidor && go run .  > logs.txt 2>&1 &

test:
	@echo "Ejecutando pruebas Hurl..."
	hurl --test requests.hurl 

down:
	@echo "Deteniendo la base de datos..."
	cd base_de_datos && sudo docker-compose down

	@echo "Deteniendo servidor Go..."
	cd servidor && -pkill -f "go run" || true