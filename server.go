package main;

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

const (
	staticFiles = "/home/frick/views"
	layoutDir = "/home/frick/views"
	layoutFile = "html/layouts/default"
	homeFile = "html/pages/home"

)

func main() {
	m := martini.Classic()
	renderOptions := render.Options{
		Directory: staticFiles,
		Layout: layoutFile,
		Extensions: []string{".html",},
		Charset: "UTF-8",
	}

	m.Use(render.Renderer(renderOptions))
	m.Use(martini.Static(staticFiles))
	m.Get("/", func(r render.Render) {
		r.HTML(200, homeFile, "")
	})


	m.Run()
	


}
