package binding

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type PathParameterStruct struct {
	Id   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

func TestBindPathParameterGetFrom(t *testing.T) {
	id := 123
	idStr := "123"
	name := "name1"
	pathData := map[string]string{"id": idStr}
	t.Run("path parameter", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://api.bilibili.com/test/123?name=name1", nil)
		q := new(PathParameterStruct)
		rFunc, _ := Form.PathParameter(req, pathData)
		Form.Bind(req, q)
		assert.Equal(t, q.Id, id)
		assert.Equal(t, q.Name, name)
		assert.Equal(t, req.Form.Get("id"), idStr)
		rFunc()
		assert.Equal(t, req.Form.Get("id"), "")
	})

	t.Run("path parameter have two", func(t *testing.T) {
		pathData := map[string]string{"id": idStr,"name":name}
		req, _ := http.NewRequest("GET", "http://api.bilibili.com/test/123/name1", nil)
		q := new(PathParameterStruct)
		Form.PathParameter(req, pathData)
		Form.Bind(req, q)
		assert.Equal(t, q.Id, id)
		assert.Equal(t, q.Name, name)

	})

	t.Run("path parameter have only", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "http://api.bilibili.com/test/123", nil)
		q := new(PathParameterStruct)
		Form.PathParameter(req, pathData)
		Form.Bind(req, q)
		assert.Equal(t, q.Id, id)
		assert.Equal(t, q.Name, "")
	})



}

func TestBindPathParameterPostFrom(t *testing.T) {
	id := 123
	idStr := "123"
	name := "name1"
	pathData := map[string]string{"id": "123"}
	t.Run("path parameter", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "http://api.bilibili.com/test/123", bytes.NewBufferString("name=name1"))
		req.Header.Add("Content-Type", MIMEPOSTForm)
		q := new(PathParameterStruct)
		rFunc, _ := Form.PathParameter(req, pathData)
		Form.Bind(req, q)
		assert.Equal(t, q.Id, id)
		assert.Equal(t, q.Name, name)
		assert.Equal(t, req.Form.Get("id"), idStr)
		rFunc()
		assert.Equal(t, req.Form.Get("id"), "")
	})


	t.Run("path parameter err", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "http://api.bilibili.com/test/123",nil)
		_,err:= Form.PathParameter(req, pathData)
		assert.Error(t,err)

	})

}


func TestBindPathParameterPostJson(t *testing.T) {
	id := 123
	idStr := "123"
	name := "name1"
	pathData := map[string]interface{}{"id": 123}
	t.Run("path parameter have more", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "http://api.bilibili.com/test/123", bytes.NewBufferString(`{"name":"name1"}`))
		req.Header.Add("Content-Type", MIMEJSON)
		q := new(PathParameterStruct)
		rFunc, _ := JSON.PathParameter(req, pathData)
		err:= JSON.Bind(req, q)
		fmt.Println("err",err)
		assert.Equal(t, q.Id, id)
		assert.Equal(t, q.Name, name)
		assert.Equal(t, req.Form.Get("id"), idStr)
		rFunc()
		assert.Equal(t, req.Form.Get("id"), "")
	})

	//
	//t.Run("path parameter body nil", func(t *testing.T) {
	//	req, _ := http.NewRequest("POST", "http://api.bilibili.com/test/123",nil)
	//	_, err := JSON.PathParameter(req, pathData)
	//
	//	assert.Error(t, err)
	//
	//
	//})
	//
	//t.Run("path parameter body err", func(t *testing.T) {
	//	req, _ := http.NewRequest("POST", "http://api.bilibili.com/test/123",bytes.NewBufferString(`{"name":"name1"`))
	//	_, err := JSON.PathParameter(req, pathData)
	//	assert.Error(t, err)
	//
	//
	//})

}
