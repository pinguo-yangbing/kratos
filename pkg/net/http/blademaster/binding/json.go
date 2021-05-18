package binding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (jsonBinding) Bind(req *http.Request, obj interface{}) error {

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(obj); err != nil {
		return errors.WithStack(err)
	}
	return validate(obj)
}

func (jsonBinding) PathParameter(req *http.Request, pathData map[string]interface{}) (func(), error) {
	contentLen := req.ContentLength
	buf := bytes.NewBuffer(make([]byte, 0, contentLen))
	if req.Body == nil {
		return func() {}, errors.WithStack(errors.New("missing json body"))
	}

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		return func() {}, errors.WithStack(err)
	}

	oldBody := buf.Bytes()

	lastIndex := bytes.LastIndexByte(oldBody, '}')
	if lastIndex < 1 {
		return func() {}, errors.WithStack(errors.New("error json body"))
	}
	fmt.Println("lastIndex",lastIndex)

	oldBody[lastIndex] = ','



	var nBuffer = bytes.NewBuffer([]byte(""))
	nBuffer.Write(oldBody)
	for field, value := range pathData {
		nBuffer.WriteByte('"')
		nBuffer.Write([]byte(field))
		nBuffer.Write([]byte(`":"`))
		vv := value.(int)
		nBuffer.Write([]byte(string(vv)))
		nBuffer.WriteByte('"')
	}
	nBuffer.WriteByte('}')

	fmt.Println("body",nBuffer.String())
	req.Body = ioutil.NopCloser(nBuffer)
	//req.Body.Read(nBuffer.Bytes())
	//fmt.Println("body",string(req.GetBody()))
	return func() {
		//fmt.Println("")
		//body := nBuffer.Bytes()[0 : lastIndex+1]
		//fmt.Println("bidy",string(body))
		//body[lastIndex] = '}'
		//req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}, nil
}