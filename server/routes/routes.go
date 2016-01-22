package routes

import (
	"github.com/vidyapatha/server/api/todo/routes"
	"github.com/vidyapatha/server/common/static"
	"github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
	static.Init(router)
}
