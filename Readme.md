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

## Consumo de la Api

Pasos:

1. Consumir el endpoint ```healthCheck``` que verifica la salud de la Api Rest.
   - Nota: El docker de base de datos toma unos segundos mientras se inicia el servicio.
2. Consumir el endpoint ```sign up``` que permite dar de alta un usuario para el consumo de los recursos. 
3. Consumir el endpoint ```sign in``` para inciar session y obtener el token de acceso.
   - Nota: Para eliminar la session del token activo debe consumir el endpoint ```sign out```
   - Copiar el token que se encuentra en la etiqueta ```Access``` y enviarlo en el header de los servicios rest ```Authorization: Bearer```
4. Consumir los servicios de: creaciòn de orden, consulta, actualizaciòn status y cancelaciòn.
   - ver documentacion: ```https://documenter.getpostman.com/view/10015938/2s7YYoBn2P```