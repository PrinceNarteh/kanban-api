package main

import "github.com/PrinceNarteh/kanban-api/internals/utils"

func main() {
	utils.NewLogger()

	app := NewApplication()
	r := app.mount()
	app.initRoutes(r)
	utils.Logger.Fatal(app.run(r))
}
