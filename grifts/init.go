package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/paganotoni/playerrel/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
