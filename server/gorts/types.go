package gorts

import "reflect"

type StructField struct {
	JSONName string
	Type     string
}

type StructInfo struct {
	Name   string
	Fields []StructField
}

type RPCMethod struct {
	Name      string
	ArgsType  reflect.Type
	ReplyType reflect.Type
}

type RPCClass struct {
	Name    string
	Methods []RPCMethod
}
