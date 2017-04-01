package main


import (
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name string `json:"name"`
}


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hell	o, World")
	})
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

