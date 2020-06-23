package yuque

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestGetAllDocs(t *testing.T) {
	m, e := GetAllDocs()
	if e != nil {
		panic(e)
	}

	b, err := jsoniter.Marshal(m)
	fmt.Printf("%v  %v\n", string(b), err)
}

func TestGetDocDetail(t *testing.T) {
	m, e := GetDocDetail("3280355")
	if e != nil {
		panic(e)
	}

	fmt.Printf("%#v \n", m)
}
