package kernel

import (
	"CAS/context"
	"CAS/exception"
	"CAS/middleware"
)

var Middleware []context.HandlerFunc

func Load() {

	Middleware = []context.HandlerFunc{
		exception.Exception,
		//middleware.Session,
		middleware.CorsMiddle,
	}

}
