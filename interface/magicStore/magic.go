package main

import "fmt"

type magicStore struct {
	value interface{}
	name  string
}

func (ms *magicStore) SetValue(v interface{}) {
	ms.value = v
}

func (ms *magicStore) GetValue() interface{} {
	return ms.value
}

func NewMagicStore(nm string) *magicStore {
	return &magicStore{name: nm}
}

func main() {
	IntStore := NewMagicStore("Integer Store")
	IntStore.SetValue(4.2)
	if v, ok := IntStore.GetValue().(float64); ok {
		v += 100
		fmt.Println(v)
	} else {
		fmt.Println("not correct type assertion")
	}

	StringStore := NewMagicStore("String Store")
	StringStore.SetValue("my string: ")
	if v, ok := StringStore.GetValue().(string); ok {
		v += "yes, this is mine"
		fmt.Println(v)
	} else {
		fmt.Println("not correct type assertion")
	}

}
