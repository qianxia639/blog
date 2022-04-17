package utils

import (
	"bytes"
	"text/template"

	"github.com/qianxia/blog/global"
)

func Loadtemplate(title string, pageNum, pageSize int) (bytes.Buffer, error) {

	tmplFile := "./source/elasticsearch/search_title.json.tmpl"
	t, err := template.ParseGlob(tmplFile)
	if err != nil {
		global.QX_LOG.Errorf("create template failed, err: %v", err)
		return bytes.Buffer{}, err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, map[string]interface{}{
		"title": title,
		"te":    string([]rune(title)[0]),
		"size":  pageSize,
		"from":  (pageNum - 1) * pageSize,
	})
	return buf, err
}
