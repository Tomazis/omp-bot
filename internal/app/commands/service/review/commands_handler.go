package review

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *ServiceReviewCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "get":
		c.Get(message)
	case "list":
		c.List(message)
	case "delete":
		c.Delete(message)
	case "new":
		c.New(message)
	case "edit":
		c.Edit(message)
	default:
		c.Default(message)
	}
}
