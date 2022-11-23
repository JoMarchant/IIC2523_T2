# API GoLang de productos

## Descripción

API de productos desarrollada en GoLang, se pueden crear productos, listarlos, actualizarlos y eliminarlos.

## Instalación

Para instalar el proyecto:
1.  Se debe clonar el repositorio
2.  Se **debe tener gcc instalado** para compilar la libreria de sqlite3
3.  Se debe tener instalado GoLang
4.  Se debe ejecutar el comando `go get` para instalar las dependencias
5.  Se debe ejecutar el comando `go run main.go` para ejecutar el proyecto

```bash
go run main.go
```
### Consideraciones

-  Por defecto corre en el puerto 8080

## Endpoints

### Crear producto

```http
  POST /products
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name` | `string` | Nombre del producto |
| `description` | `string` | Descripción del producto |
| `price` | `int` | Valor del producto |
| `exp_date` | `string` | Fecha de expiración del producto |

### Listar productos

```http
  GET /products
```

### Ver producto

```http
  GET /products/:id
```

**Nota:** El id del producto se debe enviar en la url
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | Id del producto |


### Actualizar producto

```http
  PATCH /products/:id
```
**Nota:** El id del producto se debe enviar en la url
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | Id del producto |
| `name` | `string` | Nombre del producto |
| `description` | `string` | Descripción del producto |
| `price` | `int` | Valor del producto |
| `exp_date` | `string` | Fecha de expiración del producto |

### Eliminar producto

```http
  DELETE /products/:id
```
**Nota:** El id del producto se debe enviar en la url
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id` | `int` | Id del producto |
