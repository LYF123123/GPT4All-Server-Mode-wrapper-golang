package wrapper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type GPTClient struct {
	Client  *http.Client
	ApiBase string //"http://localhost:4891/v1"
	Token   string //"not needed for a local LLM"
}

func NewGPTClient(apiBase, token string) (*GPTClient, error) {
	_, err := url.ParseRequestURI(apiBase)
	if err != nil { // the api url is wrong
		return nil, err
	}
	return &GPTClient{
		Client:  &http.Client{},
		ApiBase: apiBase,
		Token:   token,
	}, nil
}

// Models
// apiBase/models
func (c *GPTClient) ListModels() (Models, error) {
	req, err := http.NewRequest("GET", c.ApiBase+"/models", nil)
	if err != nil {
		return Models{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return Models{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Models{}, err
	}
	models := Models{}
	err = json.Unmarshal(body, &models)
	return models, err
}

// apiBase/models/{modelId}
func (c *GPTClient) RetrieveModel(modelId string) (Model, error) {
	req, err := http.NewRequest("GET", c.ApiBase+"/models/"+modelId, nil)
	if err != nil {
		return Model{}, err
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return Model{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Model{}, err
	}
	model := Model{}
	err = json.Unmarshal(body, &model)
	return model, err
}

// Chat
// !Warning Not test
// apiBase/chat/completions
func (c *GPTClient) CreateChatCompletionRawRequest(cReq ChatReq) (ChatResp, error) {
	reqBody, err := json.Marshal(cReq)
	if err != nil {
		return ChatResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return ChatResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return ChatResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ChatResp{}, err
	}
	cResp := ChatResp{}
	err = json.Unmarshal(respBody, &cResp)
	return cResp, err
}

// Completions
// apiBase/completions
func (c *GPTClient) CreateCompletion(prompt string) (CompletionResp, error) {
	cReq := getDefaultDataCompletionReq()
	if prompt != "" {
		cReq.Prompt = prompt
	}
	reqBody, err := json.Marshal(cReq)
	if err != nil {
		return CompletionResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return CompletionResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return CompletionResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CompletionResp{}, err
	}
	cResp := CompletionResp{}
	err = json.Unmarshal(respBody, &cResp)
	return cResp, err
}

// apiBase/completions
func (c *GPTClient) CreateCompletionRawRequest(cReq CompletionReq) (CompletionResp, error) {
	reqBody, err := json.Marshal(cReq)
	if err != nil {
		return CompletionResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return CompletionResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return CompletionResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return CompletionResp{}, err
	}
	cResp := CompletionResp{}
	err = json.Unmarshal(respBody, &cResp)
	return cResp, err
}

// Edits
// !Warning Not test
// apiBase/edits
func (c *GPTClient) CreateEditRawRequest(eReq EditReq) (EditResp, error) {
	reqBody, err := json.Marshal(eReq)
	if err != nil {
		return EditResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return EditResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return EditResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EditResp{}, err
	}
	eResp := EditResp{}
	err = json.Unmarshal(respBody, &eResp)
	return eResp, err
}
