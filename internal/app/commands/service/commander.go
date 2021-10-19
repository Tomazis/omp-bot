package service

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/service/review"
	"github.com/ozonmp/omp-bot/internal/app/path"
	service "github.com/ozonmp/omp-bot/internal/service/service/review"
	"log"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type ServiceCommander struct {
	bot             *tgbotapi.BotAPI
	reviewCommander Commander
}

func NewServiceCommander(bot *tgbotapi.BotAPI) *ServiceCommander {
	return &ServiceCommander{
		bot:             bot,
		reviewCommander: review.NewReviewCommander(bot, service.NewDummyReviewService()),
	}
}

func (c *ServiceCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "review":
		c.reviewCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("DemoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *ServiceCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "review":
		c.reviewCommander.HandleCommand(message, commandPath)
	default:
		log.Printf("DemoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
