package plugin

import (
	"github.com/jpfaria/golang-tests/go-plugin-example/core"
)

func NewPlugin() core.Plugin {
	return &plugin{}
}

type plugin struct {
}

func (u *plugin) Exec() string {
	return "executou aqui"
}

// exported
var Plugin plugin


/*
cd /Users/jpfaria/Projects/Go/src/github.com/jpfaria/golang-tests/go-plugin-example/plugin
go build -buildmode=plugin -o /Users/jpfaria/tmp/plugin.so plugin.go
*/
