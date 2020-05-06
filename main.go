package main

import (
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type (
	// properties of model must in Sentence case
	User struct {
		Username	string	`json:"username"`
		Password	string	`json:"password"`
	}
	TemplateRenderer struct {
		templates *template.Template
	}
)

var (
	Users = map [string]*User{}
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
	e.Static("/template", "echo-angular-example/template")
	renderer :=  &TemplateRenderer{
		templates: template.Must(template.ParseGlob("echo-angular-example/*.html")),
	}
	e.Renderer = renderer
	e.POST("/users/create", CreateUser)
	e.GET("/users/create", ShowCreateUser)
	e.GET("/users/:username", ShowUser)
	e.GET("/users/get/:username", GetUser)
	e.GET("/users/get", GetUsers)
	e.GET("/users/", ShowUsers)

	e.Logger.Fatal(e.Start(":1323"))
}

func ShowUser(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func ShowUsers(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}

func GetUser(c echo.Context) error {
	username := c.Param("username")
	user := Users[username]
	if user != nil {
		return c.JSON(http.StatusOK, user)
	}
	return c.JSON(http.StatusNotFound, nil)
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, Users)
}

func ShowCreateUser(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}

func CreateUser(c echo.Context) error {
	user := &User{}
	c.Bind(user)
	if user.Username != "" && user.Password != "" {
		u := &User{
			Username : user.Username,
			Password: user.Password,
		}
		Users[user.Username] = user
		return c.JSON(http.StatusCreated, u)
	}
	return c.JSON(http.StatusBadRequest, "Bad request. Username or Password is empty.")

}
