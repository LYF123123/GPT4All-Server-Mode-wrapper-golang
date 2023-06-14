package wrapper

// Models
type Models struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

type Model struct {
	Created     string       `json:"created"`
	ID          string       `json:"id"`
	Object      string       `json:"object"`
	OwnedBy     string       `json:"owned_by"`
	Parent      interface{}  `json:"parent"`
	Permissions []Permission `json:"permissions"`
	Root        string       `json:"root"`
}

type Permission struct {
	AllowCreateEngine  bool        `json:"allow_create_engine"`
	AllowFineTuning    bool        `json:"allow_fine_tuning"`
	AllowLogprobs      bool        `json:"allow_logprobs"`
	AllowSampling      bool        `json:"allow_sampling"`
	AllowSearchIndices bool        `json:"allow_search_indices"`
	AllowView          bool        `json:"allow_view"`
	Created            string      `json:"created"`
	Group              interface{} `json:"group"`
	ID                 string      `json:"id"`
	IsBlocking         bool        `json:"is_blocking"`
	Object             string      `json:"object"`
	Organization       string      `json:"organization"`
}

// Chat

type ChatReq struct {
	Model            string                      `json:"model"`
	Messages         []ChatReqMessage            `json:"messages"`
	Functions        []ChatReqFunction           `json:"functions,omitempty"`
	FunctionCall     interface{}                 `json:"function_call,omitempty"`
	Temperature      float64                     `json:"temperature,omitempty"`
	TopP             float64                     `json:"top_p,omitempty"`
	Stream           bool                        `json:"stream,omitempty"`
	MaxTokens        int                         `json:"max_tokens,omitempty"`
	PresencePenalty  float64                     `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64                     `json:"frequency_penalty,omitempty"`
	LogitBias        map[interface{}]interface{} `json:"logit_bias,omitempty"`
	User             string                      `json:"user,omitempty"`
}

type ChatReqMessage struct {
	Role         string      `json:"role"`
	Content      string      `json:"content,omitempty"`
	Name         string      `json:"name,omitempty"`
	FunctionCall interface{} `json:"function_call,omitempty"`
}

type ChatReqFunction struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Parameters  interface{} `json:"parameters,omitempty"`
}

type ChatResp struct {
	ID      string           `json:"id"`
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Choices []ChatRespChoice `json:"choices"`
	Usage   ChatRespUsage    `json:"usage"`
}

type ChatRespChoice struct {
	Index        int64           `json:"index"`
	Message      ChatRespMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
}

type ChatRespMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRespUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// Completions

type CompletionReq struct {
	Model            string                      `json:"model"`
	Prompt           string                      `json:"prompt,omitempty"`
	Suffix           string                      `json:"suffix,omitempty"`
	MaxTokens        int                         `json:"max_tokens,omitempty"`
	Temperature      float64                     `json:"temperature,omitempty"`
	TopP             float64                     `json:"top_p,omitempty"`
	N                int                         `json:"n,omitempty"`
	Stream           bool                        `json:"stream,omitempty"`
	Logprobs         int                         `json:"logprobs,omitempty"`
	Echo             bool                        `json:"echo,omitempty"`
	Stop             []string                    `json:"stop,omitempty"`
	PresencePenalty  float64                     `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64                     `json:"frequency_penalty,omitempty"`
	BestOf           int                         `json:"best_of,omitempty"`
	LogitBias        map[interface{}]interface{} `json:"logit_bias,omitempty"`
	User             string                      `json:"user,omitempty"`
}

func getDefaultDataCompletionReq() CompletionReq {
	c := CompletionReq{
		Model:       "mpt-7b-chat",
		Prompt:      "Please tell me the things about Golang Programming language.", //This is the default prompt
		MaxTokens:   500,
		Temperature: 0.28,
		TopP:        0.95,
		N:           1,
		Echo:        true,
		Stream:      false,
	}
	return c
}

type CompletionResp struct {
	CompletionChoices []CompletionChoice `json:"choices"`
	Created           int64              `json:"created"`
	ID                string             `json:"id"`
	Model             string             `json:"model"`
	Object            string             `json:"object"`
	Usage             CompletionUsage    `json:"usage"`
}

type CompletionChoice struct {
	FinishReason string        `json:"finish_reason"`
	Index        int64         `json:"index"`
	Logprobs     interface{}   `json:"logprobs"`
	References   []interface{} `json:"references"`
	Text         string        `json:"text"`
}

type CompletionUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// Edits

type EditReq struct {
	Model        string  `json:"model"`
	Input        string  `json:"input,omitempty"`
	Instructions string  `json:"instructions"`
	N            int     `json:"n,omitempty"`
	Temperature  float64 `json:"temperature,omitempty"`
	TopP         float64 `json:"top_p,omitempty"`
}

type EditResp struct {
	Object  string           `json:"object"`
	Created int64            `json:"created"`
	Choices []EditRespChoice `json:"choices"`
	Usage   EditRespUsage    `json:"usage"`
}

type EditRespChoice struct {
	Text  string `json:"text"`
	Index int64  `json:"index"`
}

type EditRespUsage struct {
	PromptTokens     int64 `json:"prompt_tokens"`
	CompletionTokens int64 `json:"completion_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// Image

type ImageReq struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImageResp struct {
	Created int64            `json:"created"`
	Data    []ImageRespDatum `json:"data"`
}

type ImageRespDatum struct {
	URL string `json:"url"`
}

type ImageEditReq struct {
	Image          string `form:"image"`
	Mask           string `form:"mask,omitempty"`
	Prompt         string `form:"prompt"`
	N              int    `form:"n"`
	Size           string `form:"size,omitempty"`
	ResponseFormat string `form:"response_format,omitempty"`
	User           string `form:"user,omitempty"`
}

type ImageEditResp struct {
	Created int64                `json:"created"`
	Data    []ImageEditRespDatum `json:"data"`
}

type ImageEditRespDatum struct {
	URL string `json:"url"`
}

type ImageVariationReq struct {
	Image          string `form:"image"`
	N              int    `form:"n"`
	Size           string `form:"size,omitempty"`
	ResponseFormat string `form:"response_format,omitempty"`
	User           string `form:"user,omitempty"`
}
type ImageVariationResp struct {
	Created int64                     `json:"created"`
	Data    []ImageVariationRespDatum `json:"data"`
}

type ImageVariationRespDatum struct {
	URL string `json:"url"`
}

// Embeddings

type EmbeddingReq struct {
	Model string   `json:"model"`
	Input []string `json:"input"`
	User  string   `json:"user,omitempty"`
}

type EmbeddingResp struct {
	Object string               `json:"object"`
	Data   []EmbeddingRespDatum `json:"data"`
	Model  string               `json:"model"`
	Usage  EmbeddingRespUsage   `json:"usage"`
}

type EmbeddingRespDatum struct {
	Object    string    `json:"object"`
	Embedding []float64 `json:"embedding"`
	Index     int64     `json:"index"`
}

type EmbeddingRespUsage struct {
	PromptTokens int64 `json:"prompt_tokens"`
	TotalTokens  int64 `json:"total_tokens"`
}
