package wrapper

import "net/http"

type GPTClient struct {
	Client  *http.Client
	ApiBase string //"http://localhost:4891/v1"
	Token   string //"not needed for a local LLM"
}

func NewGPTClient(apiBase, token string) *GPTClient {
	return &GPTClient{
		Client:  &http.Client{},
		ApiBase: apiBase,
		Token:   token,
	}
}
func (c *GPTClient) ListModels()
