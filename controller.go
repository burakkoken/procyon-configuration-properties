package main

import (
	web "github.com/procyon-projects/procyon-web"
)

type GreetController struct {
	greetProperties GreetProperties
}

func NewGreetController(greetProperties GreetProperties) GreetController  {
	return GreetController {
		greetProperties,
	}
}

func (controller GreetController) RegisterHandlers(registry web.HandlerRegistry)  {
	registry.Register(
		web.Get(controller.Welcome, web.Path("/welcome")),
		web.Get(controller.Bye, web.Path("/bye")),
		)
}

func (controller GreetController) Welcome(ctx *web.WebRequestContext)  {
	ctx.SetModel(controller.greetProperties.WelcomeMessage)
}

func (controller GreetController) Bye(ctx *web.WebRequestContext)  {
	ctx.SetModel(controller.greetProperties.ByeMessage)
}
