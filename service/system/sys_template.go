package system

import (
	"bytes"
	"text/template"

	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
)

type TemplateService struct {
	buf bytes.Buffer
}

var TemplateServices = new(TemplateService)

// 加载模板文件
func (*TemplateService) loadtemplate(tmplFile string) *template.Template {
	tmpl, err := template.ParseGlob(tmplFile)
	if err != nil {
		global.QX_LOG.Error("create template failed,err: %v\n", err)
		return nil
	}
	return tmpl
}

// 根据title进行搜索
func (t *TemplateService) SearchBlogByTitle(title string, pageNo, pageSize int) (bytes.Buffer, error) {
	tmpl := t.loadtemplate(config.Path("elasticsearch/search_title.json.tmpl"))
	err := tmpl.Execute(&t.buf, map[string]interface{}{
		"title": title,
		"te":    string([]rune(title)[0]),
		"size":  pageSize,
		"from":  (pageNo - 1) * pageSize,
	})
	return t.buf, err
}

// 根据title OR time进行过搜索
func (t *TemplateService) SearchBlogByTitleOrTime(title, startDate, endDate string, pageSize, pageNo int, userId uint64) (bytes.Buffer, error) {
	tmpl := t.loadtemplate(config.Path("elasticsearch/search_title_time.json.tmpl"))
	err := tmpl.Execute(&t.buf, map[string]interface{}{
		"userId":    userId,
		"title":     title,
		"startDate": startDate,
		"endDate":   endDate,
		"size":      pageSize,
		"from":      (pageNo - 1) * pageSize,
	})
	return t.buf, err
}

// 根据title AND time进行过搜索
func (t *TemplateService) SearchBlogByTitleAndTime(title, startDate, endDate string, pageSize, pageNo int, userId uint64) (bytes.Buffer, error) {
	tmpl := t.loadtemplate(config.Path("elasticsearch/search_title_time_all.json.tmpl"))
	err := tmpl.Execute(&t.buf, map[string]interface{}{
		"userId":    userId,
		"title":     title,
		"startDate": startDate,
		"endDate":   endDate,
		"size":      pageSize,
		"from":      (pageNo - 1) * pageSize,
	})
	return t.buf, err
}
