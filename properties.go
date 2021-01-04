package main

type GreetProperties struct {
	WelcomeMessage 	string	`json:"welcome" yaml:"welcome" default:"Hi"`
	ByeMessage 		string	`json:"bye" yaml:"bye" default:"See you later."`
}

func NewGreetProperties() *GreetProperties {
	return &GreetProperties{}
}

func (properties *GreetProperties) GetConfigurationPrefix() string  {
	return "message"
}

