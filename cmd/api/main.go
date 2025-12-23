package main

import "github.com/PrinceNarteh/kanban-api/internals/utils"

func main() {
	utils.NewLogger()

	app := NewApplication()
	r := app.mount()
	utils.Logger.Fatal(app.run(r))
}
