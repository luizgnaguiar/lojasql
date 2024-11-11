package routes

import (
	"github/luizgnaguiar/lojasql/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)

}
