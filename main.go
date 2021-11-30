package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-platform-advisor/commands"
	"github.com/jfrog/jfrog-cli-platform-advisor/inputs"
)

func main() {
	inputs.GetConfig()
	// plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "hello-frog"
	app.Description = "Easily greet anyone."
	app.Version = "v0.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{
		commands.GetHelloCommand()}
}
