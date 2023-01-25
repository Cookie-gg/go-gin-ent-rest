package main

import (
	"go-gin-ent-rest/infra"
	"go-gin-ent-rest/router"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	client := infra.InitDB()
	defer client.Close()

	router := router.NewRouter(client)

	http.ListenAndServe(":3000", router)
}
