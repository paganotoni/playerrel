package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/paganotoni/app/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
