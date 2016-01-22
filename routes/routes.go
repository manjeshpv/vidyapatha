package routes

import (
	"github.com/vidyapatha/api/todo/routes"
	"github.com/vidyapatha/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
