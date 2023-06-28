package utils

import (
	"bytes"
	"text/template"
)

func GetTempString(temp string, data any) (string, error) {
	tem, err := template.New(temp).Parse(temp)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	tem.Execute(&b, data)
	return b.String(), nil
}
