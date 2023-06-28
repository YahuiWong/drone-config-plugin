package tests

import (
	"os"
	"testing"
	"text/template"
)

type helloValue struct {
	hello string
	world string
}

func TestTemplateText(t *testing.T) {
	tem, _ := template.New("hello .world").Parse("hello {{ .world }}")
	value := helloValue{world: "world", hello: "hello"}
	tem.Execute(os.Stdout, map[string]interface{}{
		"world": "world",
	})
	tem.Execute(os.Stdout, value)
}
