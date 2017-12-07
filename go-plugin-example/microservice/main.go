package main

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"plugin"
	"github.com/jpfaria/golang-tests/go-plugin-example/core"
	"errors"
)

type result struct {
	Message string
}

func main() {

	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", handle)

	e.Logger.Fatal(e.Start(":7071"))
}

func handle(c echo.Context) error {

	// load module
	// 1. open the so file to load the symbols
	mod := "/Users/jpfaria/tmp/plugin.so"

	plug, err := plugin.Open(mod)

	if err != nil {
		return err
	}

	// 2. look up a symbol (an exported function or variable)
	// in this case, variable Greeter
	symPlugin, err := plug.Lookup("Plugin")

	if err != nil {
		return err
	}

	// 3. Assert that loaded symbol is of a desired type
	// in this case interface type Greeter (defined above)
	var p core.Plugin
	p, ok := symPlugin.(core.Plugin)
	if !ok {
		return errors.New("unexpected type from module symbol")
	}

	// 4. use the module
	message := p.Exec()

	return c.JSON(http.StatusOK, result{Message: message})
}
