package templates

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

type TemplatesMap map[string][]string

var (
	templatesMap = TemplatesMap{
		"index":          []string{"index.html"},
		"login":          []string{"login.html"},
		"sign_up":        []string{"sign_up.html"},
		"connect_school": []string{"connect_school.html"},
	}
)

func PrepareRouter(r *gin.Engine) {
	r.HTMLRender = loadTemplates(r, templatesMap)
}

func loadTemplates(r *gin.Engine, tmplMap TemplatesMap) multitemplate.Render {
	templatesDir := "web/templates/"
	staticDir := "/web/static"
	r.Static(staticDir, "."+staticDir)
	templates := multitemplate.New()
	for name, tmpls := range tmplMap {
		var paths []string

		// Default templates
		paths = append(paths, templatesDir+"base.html", templatesDir+"navbar.html")

		// Add path prefix to files' names
		for _, p := range tmpls {
			paths = append(paths, templatesDir+p)
		}

		templates.AddFromFiles(name, paths...)
	}

	return templates
}
