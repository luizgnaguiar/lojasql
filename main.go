package main

import (
	"github/luizgnaguiar/lojasql/routes"
	"net/http"

	_ "github.com/lib/pq"
)

//var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)

}
