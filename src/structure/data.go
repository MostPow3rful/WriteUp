package structure

type Data struct {
	MysqlUsername  string `json:"username"`
	MysqlPassword  string `json:"password"`
	TelegramApi    string `json:"telegram_api"`
	DiscordWebhook string `json:"discord_webhook"`
	ChannelName    string `json:"channel_name"`
}
