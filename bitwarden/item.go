package bitwarden


type ItemResponse struct {
	Login        ItemResponseLogin   `json:"login"`
}

type ItemResponseLogin struct {
	Username        string   `json:"username"`
	Password        string   `json:"password"`
}

