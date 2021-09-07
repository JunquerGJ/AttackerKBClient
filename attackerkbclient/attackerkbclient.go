package attackerkbclient

const BaseUrl = "https://api.attackerkb.com/v1/"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}
