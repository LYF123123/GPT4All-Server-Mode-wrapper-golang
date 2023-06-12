package wrapper

import "testing"

func TestListModels(t *testing.T) {
	c, err := NewGPTClient("http://localhost:4891/v1", "")
	if err != nil {
		t.Fatal(err)
	}
	models, err := c.ListModels()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(models)
}

func TestRetrieveModel(t *testing.T) {
	c, err := NewGPTClient("http://localhost:4891/v1", "")
	if err != nil {
		t.Fatal(err)
	}
	model, err := c.RetrieveModel("mpt-7b-chat")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(model)
}

func TestCreateCompletion(t *testing.T) {
	c, err := NewGPTClient("http://localhost:4891/v1", "")
	if err != nil {
		t.Fatal(err)
	}
	cResp,err:=c.CreateCompletion("")
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(cResp)

}

func TestCreateCompletionRawRequest(t *testing.T) {
	c, err := NewGPTClient("http://localhost:4891/v1", "")
	if err != nil {
		t.Fatal(err)
	}
	cReq:=getDefaultDataCompletionReq()
	cReq.Prompt="Please tell me the things about Rust Programming language."
	cResp,err:=c.CreateCompletionRawRequest(cReq)
	if err!=nil{
		t.Fatal(err)
	}
	t.Log(cResp)
}