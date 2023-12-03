package bitwarden

type Configuration struct {
	Session   string `json:"Session"`
	ItemID    string `json:"ItemID"`
	LogPath   string `json:"LogPath"`
	EnableLog bool   `json:"EnableLog"`
}
