package review

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/service/review"
)

type ReviewCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)
}

type ServiceReviewCommander struct {
	bot           *tgbotapi.BotAPI
	reviewService service.ReviewService
}

func NewReviewCommander(bot *tgbotapi.BotAPI, service service.ReviewService) *ServiceReviewCommander {
	return &ServiceReviewCommander{
		bot:           bot,
		reviewService: service,
	}
}
