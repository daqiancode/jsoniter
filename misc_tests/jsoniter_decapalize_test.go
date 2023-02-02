package misc_tests

import (
	"fmt"
	"testing"

	"github.com/daqiancode/jsoniter"
	"github.com/stretchr/testify/require"
)

type Struct struct {
	Field    string
	FieldAge int
	MyAddr   string `json:"Addr"`
	HTTPAddr string
}

func Test_Decapitalize(t *testing.T) {
	s := Struct{Field: "field", FieldAge: 10, MyAddr: "earth"}
	json := jsoniter.Decapitalized
	bs, err := json.Marshal(s)
	require.Nil(t, err)
	var s1 Struct
	jsoniter.Unmarshal(bs, &s1)
	require.Equal(t, s1.Field, s.Field)
	require.Equal(t, s1.MyAddr, s.MyAddr)
}

func Test_Snake(t *testing.T) {
	s := Struct{Field: "field", FieldAge: 10, MyAddr: "earth"}
	json := jsoniter.SnakeCase
	bs, err := json.Marshal(s)
	require.Nil(t, err)
	var s1 Struct
	json.Unmarshal(bs, &s1)
	require.Equal(t, s1.Field, s.Field)
	require.Equal(t, s1.MyAddr, s.MyAddr)
	fmt.Println(string(bs))
}

// func Test_s(t *testing.T) {
// 	fmt.Println(snake("CamelCase"))
// 	fmt.Println(snake("camelCamelCase"))
// 	fmt.Println(snake("CCamelCCamelCase"))
// 	fmt.Println(snake("CamelCamelCase"))
// 	fmt.Println(snake("A"))
// 	fmt.Println(snake("a"))
// 	fmt.Println(snake("HTTPId"))
// 	fmt.Println(snake("HTTP2Id"))
// 	// snake("camelCamelCase")
// 	// snake("CCamelCCamelCase")
// 	// snake("CamelCamelCase")
// 	// snake("aB")
// 	// snake("A")
// 	// snake("a")
// 	// snake("HTTPId")
// }
