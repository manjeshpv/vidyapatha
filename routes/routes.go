package routes

import (
	todoroutes "github.com/manjeshpv/vidyapatha/api/todo/routes"
//	"github.com/manjeshpv/vidyapatha/common/static"
	httprouter "github.com/julienschmidt/httprouter"
)

func Init(router *httprouter.Router) {
	todoroutes.Init(router)
//	static.Init(router)
}
