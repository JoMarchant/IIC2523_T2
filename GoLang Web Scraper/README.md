# GoLang Scraper
\
Scraper para productos de Ripley.cl


## <br/>Requisitos
- GoQuery 

Para instalarlo se ejecuta
```powershell
go install github.com/PuerkitoBio/goquery@latest
```

Si aparece un error del tipo
```powershell
go: go.mod file not found in current directory or any parent directory.
        'go get' is no longer supported outside a module.
        To build and install a command, use 'go install' with a version,
        like 'go install example.com/cmd@latest'
        For more information, see https://golang.org/doc/go-get-install-deprecation
        or run 'go help get' or 'go help install'.
```

tienes que ejecutar 
```powershell
go env -w GO111MODULE=off
```

## Ejecución

Para correr el scraper, se ejecuta el comando,
```powershell
go run main.go
```
Luego se te pedirán dos inputs, el primero para saber que productos se va a escrapear, y el segundo son la cantidad de páginas que se van a escrapear. Al terminar la ejecución, tendrás un archivo csv con el mismo nombre del producto que buscaste. 