package service

import "fmt"

var Reviews = []Review{
	{Title: "Самый лучший сервис", Author: "Anton L.", Text: "Сервис на кончиках пальцев", Score: 10},
	{Title: "PandB", Author: "Big B.", Text: "ヾ(＠⌒ー⌒＠)ノ", Score: 8},
	{Title: "Bruh", Author: "Carl J.", Text: "C'mon", Score: 3},
	{Title: "ナイス！！！", Author: "田中　森", Text: "最後まで見てた。(>_<)", Score: 9},
}

type Review struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
	Score  uint8  `json:"score"`
}

func (r *Review) String() string {
	payload := fmt.Sprintf("#############\n"+
		"%s\n"+
		"Author: %s\n\n"+
		"%s\n\n"+
		"Stars: %d/10⭐", r.Title, r.Author, r.Text, r.Score)
	return payload
}
