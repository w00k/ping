# Servicio Ping

Servicio **ping** que llama al servicio **pong**, para entregar una respuesta. Ambos servicios serán deployados en kubernetes. 

## Pre-requisiros

Tener instalado: 
* Golang. 
* Descargar el proyecto pong desde https://github.com/w00k/pong
* Descargar las dependencias. 

En la ruta donde está el proyecto 
```bash 
go mod download
```

* Agergar en las variables de entorno la URL_PONG con valor **http://localhost:8081/pong**. 

Como ejemplo en Windows, abrir cmd y ejecutar. 
```bash 
set URL_PONG=http://localhost:8081/pong
```

## Ejecutar

Solo para ejecutar 
```bash 
go run ping.go
```

Para compilar en Windows
```bash 
go build -o path/ping.exe ping.go
```
