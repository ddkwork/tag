package main

import (
	"tag"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("tag", func(w *unison.Window) {
		tag.New().Layout(w.Content())
	})
}
