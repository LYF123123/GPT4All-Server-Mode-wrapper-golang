package wrapper

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
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
	req, err := http.NewRequest("POST", c.ApiBase+"/chat/completions", bytes.NewBuffer(reqBody))
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
	req, err := http.NewRequest("POST", c.ApiBase+"/edits", bytes.NewBuffer(reqBody))
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

// Image
// !Warning Not test
// apiBase/images/generations
func (c *GPTClient) CreateImageRawRequest(iReq ImageReq) (ImageResp, error) {
	reqBody, err := json.Marshal(iReq)
	if err != nil {
		return ImageResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/images/generations", bytes.NewBuffer(reqBody))
	if err != nil {
		return ImageResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return ImageResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ImageResp{}, err
	}
	iResp := ImageResp{}
	err = json.Unmarshal(respBody, &iResp)
	return iResp, err
}

// !Warning Not test
// apiBase/images/edits
func (c *GPTClient) CreateImageEditRawRequest(iReq ImageEditReq) (ImageEditResp, error) {
	reqBody := url.Values{}
	reqBody.Set("image", iReq.Image)
	reqBody.Set("mask", iReq.Mask)
	reqBody.Set("prompt", iReq.Prompt)
	reqBody.Set("n", strconv.Itoa(iReq.N)) //! maybe not work
	reqBody.Set("size", iReq.Size)
	reqBody.Set("response_format", iReq.ResponseFormat)
	reqBody.Set("user", iReq.User)

	req, err := http.NewRequest("POST", c.ApiBase+"/images/edits", bytes.NewBuffer([]byte(reqBody.Encode())))
	if err != nil {
		return ImageEditResp{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return ImageEditResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ImageEditResp{}, err
	}
	iResp := ImageEditResp{}
	err = json.Unmarshal(respBody, &iResp)
	return iResp, err
}

// !Warning Not test
// apiBase/images/variations
func (c *GPTClient) CreateImageVariationRawRequest(iReq ImageVariationReq) (ImageVariationResp, error) {
	reqBody := url.Values{}
	reqBody.Set("image", iReq.Image)
	reqBody.Set("n", strconv.Itoa(iReq.N)) //! maybe not work
	reqBody.Set("size", iReq.Size)
	reqBody.Set("response_format", iReq.ResponseFormat)
	reqBody.Set("user", iReq.User)

	req, err := http.NewRequest("POST", c.ApiBase+"/images/variations", bytes.NewBuffer([]byte(reqBody.Encode())))
	if err != nil {
		return ImageVariationResp{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return ImageVariationResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ImageVariationResp{}, err
	}
	iResp := ImageVariationResp{}
	err = json.Unmarshal(respBody, &iResp)
	return iResp, err
}

// Embeddings
// !Warning Not test
// apiBase/embeddings
func (c *GPTClient) CreateEmbeddingsRawRequest(eReq EmbeddingReq) (EmbeddingResp, error) {
	reqBody, err := json.Marshal(eReq)
	if err != nil {
		return EmbeddingResp{}, err
	}
	req, err := http.NewRequest("POST", c.ApiBase+"/embeddings", bytes.NewBuffer(reqBody))
	if err != nil {
		return EmbeddingResp{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)
	resp, err := c.Client.Do(req)
	if err != nil {
		return EmbeddingResp{}, err
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return EmbeddingResp{}, err
	}
	eResp := EmbeddingResp{}
	err = json.Unmarshal(respBody, &eResp)
	return eResp, err
}
