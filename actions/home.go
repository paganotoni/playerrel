package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/paganotoni/playerrel/models"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	var player models.Player
	err := tx.Eager().Order("created_at desc").First(&player)
	if err != nil {
		return errors.WithStack(err)
	}

	c.Set("player", player)

	return c.Render(200, r.HTML("index.html"))
}

func CreateHandler(c buffalo.Context) error {
	tx := c.Value("tx").(*pop.Connection)

	player := models.Player{
		Firstname: "John",
		Lastname:  "Smith",
	}

	err := tx.Create(&player)

	tx.Create(&models.ContactInfo{
		ParentName:  "Ana Smith",
		PhoneNumber: "3993939393",
		PlayerID:    player.ID,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
