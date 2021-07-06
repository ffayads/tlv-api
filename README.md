# TLV-API

## Requerimientos

- Instalacion y ajuste del ambiente de desarrollo Go https://www.gopherguides.com/courses/preparing-your-environment-for-go-development
- Clonar el proyecto en $GOPATH/src

## Dependencias

Las dependencias estan agregadas con el go.mod, en caso de agregar una nueva dependencia realizar un go get

## Arrancar el proyecto

Golang es un proyecto compilado por lo cual se debe ejecutar la siguiente sentencia para que el proyecto se ejecute como desarrollador:

	$ go run main.go

Para ejecutar el proyecto y crear el archivo compilado como ambiente de produccion se debe ejecutar la siguiente sentencia:

	$ go build main.go -o stock-exchange

Al ingresar a [http://127.0.0.1:5000](http://127.0.0.1:5000) podra ver el proyecto corriendo, para cambiar el puerto puede agregar la variable de entorno APIPORT, ejemplo:

````
export APIPORT=5000
````

## Ejecucion de endpoints

Se desarrollo un enpoint que recibe un arreglo de bytes que un tiene un TLV, este endpoint se consume de la siguiente manera:
````
curl --location --request POST 'http://localhost:5000/getTlv' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [49 ,49 ,65 ,66 ,51 ,57 ,56 ,55 ,54 ,53 ,85 ,74 ,49 ,65 ,48 ,53 ,48 ,50 ,48 ,48 ,78 ,50 ,51]
}'
````


