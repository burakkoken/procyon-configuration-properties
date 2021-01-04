package main

import (
	"github.com/procyon-projects/procyon"
	core "github.com/procyon-projects/procyon-core"
)

func main()  {
	procyon.NewProcyonApplication().Run()
}

func init()  {
	// controller
	core.Register(NewGreetController)
	// properties
	core.Register(NewGreetProperties)
}