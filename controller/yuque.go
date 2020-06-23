package controller

import (
	"bastion/service/yuque"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Page struct {
}

func (p *Page) Docs(c *gin.Context) {
	docs, e := yuque.GetAllDocs()
	if e != nil {
		panic(e)
	}

	var filterDocs []*yuque.Doc
	for _, item := range docs {

		// https://www.yuque.com/yuque/developer/docserializer
		//• public - 是否公开 [1 - 公开, 0 - 私密]
		//• status - 状态 [1 - 正常, 0 - 草稿]
		if item.Public == 1 && item.Status == 1 {
			filterDocs = append(filterDocs, item)
		}
	}
	c.HTML(http.StatusOK, "list.html", map[string]interface{}{
		"title": "南的博客",
		"data":  filterDocs,
	})
	//app.Success(c, filterDocs)
	return
}

func (p *Page) DocContent(c *gin.Context) {
	id := c.Param("id")
	doc, e := yuque.GetDocDetail(id)
	if e != nil {
		panic(e)
	}
	c.HTML(http.StatusOK, "detail.html", map[string]interface{}{
		"title":   doc.Data.Title,
		"content": doc.Data,
	})
	//app.Success(c, doc)
	return
}

func (p *Page) Stat(c *gin.Context) {
	c.HTML(http.StatusOK, "stat.html", nil)
	return
}

func (p *Page) MovieForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
	return
}

func (p *Page) Errors(c *gin.Context) {
	c.HTML(http.StatusOK, "error.html", nil)
	return
}