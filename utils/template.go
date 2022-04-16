package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func Loadtemplate(title string, pageNum, pageSize int) (bytes.Buffer, error) {

	tmplFile := "./source/elasticsearch/search_title.json.tmpl"
	t, err := template.ParseGlob(tmplFile)
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return bytes.Buffer{}, err
	}

	// t = template.Must(template.New("").ParseGlob(tmplFile))

	var buf bytes.Buffer
	err = t.Execute(&buf, map[string]interface{}{
		"title": title,
		"te":    string([]rune(title)[0]),
		"size":  pageSize,
		"from":  (pageNum - 1) * pageSize,
	})
	return buf, err
}
