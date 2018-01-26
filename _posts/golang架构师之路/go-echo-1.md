#golang echo web framework

## install 
go get -u github.com/labstack/echo/...

## hello world
```go
package main

import (
	"net/http"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world china")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

## 路径参数
```go
e.GET("/users/:id",controller.GetUser)
func GetUser(c echo.Context) error {
    return c.String(http.StatusOK,c.Param("id"))
}
``` 
browser http://localhost:1323/`users/300`  you will see `300`

## 查询参数
