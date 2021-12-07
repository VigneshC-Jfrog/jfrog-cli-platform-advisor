package main

import (
	"github.com/jfrog/jfrog-cli-core/v2/plugins"
	"github.com/jfrog/jfrog-cli-core/v2/plugins/components"
	"github.com/jfrog/jfrog-cli-platform-advisor/controller"
)

func main() {
	plugins.PluginMain(getApp())
}

func getApp() components.App {
	app := components.App{}
	app.Name = "platform-advisor"
	app.Description = "Advises JFrog platform best practises"
	app.Version = "v1.1.0"
	app.Commands = getCommands()
	return app
}

func getCommands() []components.Command {
	return []components.Command{controller.GetAdvisory()}
}
