package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bobTheBuilder7/bunrouter"
	"github.com/bobTheBuilder7/bunrouter/extra/reqlog"
)

func main() {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware(
			reqlog.FromEnv("BUNDEBUG"),
		)),
	).Compat()

	router.GET("/", indexHandler)

	router.WithGroup("/api", func(g *bunrouter.CompatGroup) {
		g.GET("/users/:id", debugHandler)
		g.GET("/users/current", debugHandler)
		g.GET("/users/*path", debugHandler)
	})

	log.Println("listening on http://localhost:9999")
	log.Println(http.ListenAndServe(":9999", router))
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	if err := indexTemplate().Execute(w, nil); err != nil {
		panic(err)
	}
}

func debugHandler(w http.ResponseWriter, req *http.Request) {
	params := bunrouter.ParamsFromContext(req.Context())

	_ = bunrouter.JSON(w, bunrouter.H{
		"route":  params.Route(),
		"params": params.Map(),
	})
}

var indexTmpl = `
<html>
  <h1>Welcome</h1>
  <ul>
    <li><a href="/api/users/123">/api/users/123</a></li>
    <li><a href="/api/users/current">/api/users/current</a></li>
    <li><a href="/api/users/foo/bar">/api/users/foo/bar</a></li>
  </ul>
</html>
`

func indexTemplate() *template.Template {
	return template.Must(template.New("index").Parse(indexTmpl))
}
