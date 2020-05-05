package main

import (
	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo"
)

type (
	User struct {
		username	string	`json:"username"`
		password	string	`json:"password"`
	}
	TemplateRenderer struct {
		templates *template.Template
	}
)

var (
	users map [string]*User
)

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}
	return t.templates.ExecuteTemplate(w, name, data)
}


func main()  {
	e := echo.New()
	e.Static("/static", "echo-angular-example")
	renderer :=  &TemplateRenderer{
		templates: template.Must(template.ParseGlob("echo-angular-example/*.html")),
	}
	e.Renderer = renderer
	e.POST("/users/create", CreateUser)
	e.GET("/users/create", ShowCreateUser)
	e.GET("/users/:username", ShowUser)
	e.GET("/users/get/:username", GetUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func ShowUser(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func GetUser(c echo.Context) error {
	username := c.Param("username")
	user := users[username]
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, nil)
}

func ShowCreateUser(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	if username != "" && password != "" {
		u := &User{
			username : username,
			password: password,
		}
		return c.JSON(http.StatusCreated, u)
	}
	return c.JSON(http.StatusBadRequest, "Bad request")

}
