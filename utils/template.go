package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func Loadtemplate(title string) (bytes.Buffer, error) {

	tmplFile := "./source/elasticsearch/search_title.json.tmpl"
	t, err := template.ParseGlob(tmplFile)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return bytes.Buffer{}, err
	}

	// t = template.Must(template.New("").ParseGlob(tmplFile))

	var buf bytes.Buffer
	err = t.Execute(&buf, map[string]string{
		"title": title,
		"te":    string([]rune(title)[0]),
	})
	return buf, err
}
