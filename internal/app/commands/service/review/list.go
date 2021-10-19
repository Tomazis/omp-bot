package review

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"log"
)

type CallbackListData struct {
	Cursor int `json:"cursor"`
	Limit  int `json:"limit"`
}

func (c *ServiceReviewCommander) GetList(cursor int, limit int) (string, error) {
	outputMsgText := ""
	reviews, err := c.reviewService.List(uint64(cursor), uint64(limit))
	if err != nil {
		return outputMsgText, err
	}
	for _, r := range reviews {
		outputMsgText += r.String()
		outputMsgText += "\n\n"
	}

	return outputMsgText, nil
}

func (c *ServiceReviewCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := "Here all the reviews: \n\n"
	cursor, limit := 0, 2

	res, err := c.GetList(cursor, limit)
	if err != nil {
		log.Printf("fail to get list of reviews with idx %d: %v", cursor, err)
		return
	}
	outputMsgText += res

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Cursor: cursor,
		Limit:  limit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "service",
		Subdomain:    "review",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String()),
		),
	)

	c.bot.Send(msg)
}

func (c *ServiceReviewCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("fail to get callback json")
		return
	}
	cursor, limit := parsedData.Cursor, parsedData.Limit

	outputMsgText, err := c.GetList(cursor+limit, limit)

	if err != nil {
		log.Printf("fail to get list of reviews with idx %d and limit %d: %v", cursor, limit, err)
		return
	}
	serializedDataNext, err := json.Marshal(CallbackListData{
		Cursor: cursor + limit,
		Limit:  limit,
	})
	if err != nil {
		log.Printf("%v", err)
		return
	}
	serializedDataBack, err := json.Marshal(CallbackListData{
		Cursor: cursor - limit,
		Limit:  limit,
	})

	if err != nil {
		log.Printf("%v", err)
		return
	}

	callbackPathNext := path.CallbackPath{
		Domain:       "service",
		Subdomain:    "review",
		CallbackName: "list",
		CallbackData: string(serializedDataNext),
	}

	callbackPathBack := path.CallbackPath{
		Domain:       "service",
		Subdomain:    "review",
		CallbackName: "list",
		CallbackData: string(serializedDataBack),
	}

	delMsg := tgbotapi.NewDeleteMessage(callback.Message.Chat.ID, callback.Message.MessageID)

	msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Back", callbackPathBack.String()),
			tgbotapi.NewInlineKeyboardButtonData("Next", callbackPathNext.String()),
		),
	)
	c.bot.Send(msg)

	c.bot.Send(delMsg)
}
