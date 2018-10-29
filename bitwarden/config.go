package bitwarden

type Configuration struct {
	Session string  `env:"BW_SESSION"`
	ItemID string  `env:"BW_ITEMID"`
	LogPath string
	EnableLog bool
}
