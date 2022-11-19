# Instrucciones para ejecutar la aplicación

- Primero, para instalar las dependencias: `go get .`
- Luego, para ejecutar la aplicación: `go run main.go`

# Consideraciones

- Se le agregó un atributo `id` a los productos que es un entero. Este entero no es autogenerado, por lo que se debe pasar al momento de crear un producto y no se valida que sea único.
- Por un tema de simplicidad la fecha de expiración es un string y el valor del producto un entero. No se valida que los tipos calcen al momento de crear o actualizar un producto.
- No se utilizó ninguna forma de almacenamiento, por lo que los productos viven en memoria. Una vez que se detiene la aplicación se perderán todos los productos creados.
- La aplicación inicia sin productos creados.
- La aplicación corre en el puerto 3000.

# Endpoints

## GET: `localhost:3000/products`

Retorna un arreglo con todos los productos.

## GET: `localhost:3000/products/{id}`

Retorna el producto con ese id en caso de existir.

## POST: `localhost:3000/products`

Debe enviar un body con el siguiente formato para crear un producto:
```
{
    "id": 0, 
    "name": "Nombre del producto", 
    "description": "Descripción del producto", 
    "value": 26,
    "date": "8/11/2023"
}
```

## DELETE `localhost:3000/products/{id}`

Elimina el producto con dicho id en caso de existir.

## PUT `locahost:3000/products/{id}`

Actualiza el producto con dicho id en caso de existir. El formato del body es igual al del post. No es necesario enviar los atributos que no serán actualizados.