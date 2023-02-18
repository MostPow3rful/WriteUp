package telegram

import (
	"github.com/JesusKian/WriteUp/src/config"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Send(_title, _link, _pubDate string) {
	var (
		err  error          = nil
		resp *http.Response = &http.Response{}
		body []byte         = []byte{}
		url  string         = fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.EnvData.TelegramApi)
		text string         = fmt.Sprintf(`	      📰 %s
		
		📅 %s
		
		🔗 %s`, _title, _pubDate, _link)
	)

	body, _ = json.Marshal(map[string]interface{}{
		"chat_id":    config.EnvData.ChannelName,
		"text":       text,
		"parse_mode": "Markdown",
	})

	resp, err = http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		config.SetLog("E", "telegram.Send() -> Couldn't Send Message To Telegram Channel")
		config.SetLog("D", err.Error())
	}

	resp.Body.Close()
}
