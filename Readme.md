# Project Name: Services backend core-service

### Descripción

Proyecto para la construccion de servicios backend para la creacion de ordenes de entrega

### Construcción 🛠️
* **Language:** Golang
* **Framework:** Fiber

## Requerimientos
- Docker
- Git
- Terminal(Cmder,cmd)

## Instalación

Pasos:

1. Clone el proyecto.
2. Clone el archivo ```.env.example``` con el nombre ```.env``` ubicado en la ruta: ```core-service/services/delivery-service/deployments```
3. Clone el archivo ```docker-compose.example.yml``` con el nombre ```docker-compose.yml``` ubicado en la ruta: ```core-service/services/delivery-service/deployments```
4. Ubicarse en la ruta: ```core-service/services/delivery-service/deployments``` con el terminal(Cmder,cmd) y construya las imagenes ejecutando el siguiente comando: ```docker-compose up -d --build```

En caso que requiera detener los docker ejecute el siguiente comando:
- Detener docker: ```docker-compose down```
