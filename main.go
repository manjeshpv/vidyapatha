package main

import (
	"fmt"
	routes "github.com/manjeshpv/vidyapatha/routes"
	httprouter "github.com/julienschmidt/httprouter"
	"net/http"
)

const port string = ":3333"

func main() {
	fmt.Printf("Running at %v\n", port)

	r := httprouter.New()

	routes.Init(r)

	http.ListenAndServe(port, r)
}
