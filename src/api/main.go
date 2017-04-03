package main


import (
	"github.com/labstack/echo"
	"net/http"
)

type Todo struct {
	Id int `json:"id"`
	Text string `json:"text"`
	Status bool `json:"status"`
}

func All(store map[int]Todo) echo.HandlerFunc {
	return func(c echo.Context) error {
		vs := []Todo{}
		for _, v := range store {
			vs = append(vs, v)
		}
		return c.JSON(http.StatusOK, vs)
	}
}

func Get(store interface{}, my interface{}) echo.HandlerFunc {
	return nil
}

func Create(store interface{}, my interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo := new(Todo)
		if err := c.Bind(todo); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, todo)
	}
}

func Edit(store interface{}, my interface{}) echo.HandlerFunc {
	return nil
}

func Delete(store interface{}, my interface{}) echo.HandlerFunc {
	return nil
}



func main() {
	e := echo.New()
	store := map[int]Todo{}

	e.GET("/todos/", All(store))
	e.GET("/todos/:id", Get(&store, &Todo{}))
	e.POST("/todos/", Create(&store, &Todo{}))
	e.PATCH("/todos/:id", Edit(&store, &Todo{}))
	e.DELETE("/todos/:id", Delete(&store, &Todo{}))

	e.Logger.Fatal(e.Start(":3000"))
}

