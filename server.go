package main;

import(
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)


func main() {
	m := martini.Classic()
	renderOptions := render.Options{
		Directory: "/Users/warren/workbench/views",
		Layout: "html/layouts/default",
		Extensions: []string{".html",},
		Charset: "UTF-8",
	}

	m.Use(render.Renderer(renderOptions))
	m.Use(martini.Static("/Users/warren/workbench/views"))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "html/pages/home", "")
	})


	m.Run()
	


}
