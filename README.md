[![Sourcegraph](https://sourcegraph.com/github.com/daqiancode/jsoniter/-/badge.svg)](https://sourcegraph.com/github.com/daqiancode/jsoniter?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/daqiancode/jsoniter)
[![Build Status](https://travis-ci.org/json-iterator/go.svg?branch=master)](https://travis-ci.org/json-iterator/go)
[![codecov](https://codecov.io/gh/json-iterator/go/branch/master/graph/badge.svg)](https://codecov.io/gh/json-iterator/go)
[![rcard](https://goreportcard.com/badge/github.com/daqiancode/jsoniter)](https://goreportcard.com/report/github.com/daqiancode/jsoniter)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/json-iterator/go/master/LICENSE)
[![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/json-iterator/Lobby)

A high-performance 100% compatible drop-in replacement of "encoding/json"
## New features:
1. Decapitalize for each field of struct. 
2. Snakecase for each field of struct. 

## Added new features: Decapitalize & Snakecase
```go
type Struct struct {
	Field    string
	FieldAge int
	MyAddr   string `json:"Addr"`
	IPAddr   string 
}
func Test_Capitalize(t *testing.T) {
	s := Struct{Field: "field", FieldAge: 10, MyAddr: "earth"}
	json := jsoniter.Decapitalized
	bs, err := json.Marshal(s)
	require.Nil(t, err)
    fmt.Println(bs ,err)
	var s1 Struct
	jsoniter.Unmarshal(bs, &s1)
	require.Equal(t, s1.Field, s.Field)
	require.Equal(t, s1.MyAddr, s.MyAddr)
	fmt.Println(string(bs))
    //output: {"field":"field","fieldAge":10,"Addr":"earth" , "ipAddr":""}
}
func Test_Snake(t *testing.T) {
	s := Struct{Field: "field", FieldAge: 10, MyAddr: "earth"}
	json := jsoniter.SnakeCase
	bs, err := json.Marshal(s)
	require.Nil(t, err)
	require.Equal(t, true, true)
	var s1 Struct
	json.Unmarshal(bs, &s1)
	require.Equal(t, s1.Field, s.Field)
	require.Equal(t, s1.MyAddr, s.MyAddr)
	fmt.Println(string(bs))
    //output: {"field":"field","field_age":10,"Addr":"earth"}
}
```

# Benchmark

![benchmark](http://jsoniter.com/benchmarks/go-benchmark.png)

Source code: https://github.com/daqiancode/jsoniter-benchmark/blob/master/src/github.com/daqiancode/jsoniter-benchmark/benchmark_medium_payload_test.go

Raw Result (easyjson requires static code generation)

|                 | ns/op       | allocation bytes | allocation times |
| --------------- | ----------- | ---------------- | ---------------- |
| std decode      | 35510 ns/op | 1960 B/op        | 99 allocs/op     |
| easyjson decode | 8499 ns/op  | 160 B/op         | 4 allocs/op      |
| jsoniter decode | 5623 ns/op  | 160 B/op         | 3 allocs/op      |
| std encode      | 2213 ns/op  | 712 B/op         | 5 allocs/op      |
| easyjson encode | 883 ns/op   | 576 B/op         | 3 allocs/op      |
| jsoniter encode | 837 ns/op   | 384 B/op         | 4 allocs/op      |

Always benchmark with your own workload.
The result depends heavily on the data input.

# Usage

100% compatibility with standard lib

Replace

```go
import "encoding/json"
json.Marshal(&data)
```

with

```go
import "github.com/daqiancode/jsoniter"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Marshal(&data)
```

Replace

```go
import "encoding/json"
json.Unmarshal(input, &data)
```

with

```go
import "github.com/daqiancode/jsoniter"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
json.Unmarshal(input, &data)
```


[More documentation](http://jsoniter.com/migrate-from-go-std.html)

# How to get

```
go get github.com/daqiancode/jsoniter
```

# Contribution Welcomed !

Contributors

- [thockin](https://github.com/thockin)
- [mattn](https://github.com/mattn)
- [cch123](https://github.com/cch123)
- [Oleg Shaldybin](https://github.com/olegshaldybin)
- [Jason Toffaletti](https://github.com/toffaletti)

Report issue or pull request, or email taowen@gmail.com, or [![Gitter chat](https://badges.gitter.im/gitterHQ/gitter.png)](https://gitter.im/json-iterator/Lobby)
