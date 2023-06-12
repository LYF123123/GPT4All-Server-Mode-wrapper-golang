package wrapper

type Models struct {
	Data   []Datum `json:"data"`
	Object string  `json:"object"`
}

type Datum struct {
	Created     string       `json:"created"`
	ID          string       `json:"id"`
	Object      string       `json:"object"`
	OwnedBy     string       `json:"owned_by"`
	Parent      interface{}  `json:"parent"`
	Permissions []Permission `json:"permissions"`
	Root        string       `json:"root"`
}

type Permission struct {
	//Todo
}
