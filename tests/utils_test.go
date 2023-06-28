package tests

import (
	"bytes"
	"os"
	"testing"
	"text/template"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin/config"
)

type helloValue struct {
	Hello string
	World string
}

func TestTemplateText(t *testing.T) {
	tem, _ := template.New("hello .world").Parse("hello {{ .World }}")
	value := helloValue{World: "world", Hello: "hello"}
	tem.Execute(os.Stdout, map[string]interface{}{
		"world": "world",
	})
	tem.Execute(os.Stdout, value)
}
func getTempString(temp string, data any) (string, error) {
	tem, err := template.New(temp).Parse(temp)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	tem.Execute(&b, data)
	return b.String(), nil
}
func TestPathTemplate(t *testing.T) {
	configReq := config.Request{Build: drone.Build{
		Number: 121,
	},
		Repo: drone.Repo{
			HTTPURL: "htttp://www.baidu.com",
		}}
	s, _ := getTempString("hello {{ .Build.Number }}", configReq)
	t.Log(s)
}
