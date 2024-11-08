package main

import "fmt"

type array []any

func (a array) Map(mapFunc func(interface{}) interface{}) array {
	out := array{}
	for _, v := range a {
		out = append(out, mapFunc(v))
	}
	return out
}

func (a array) Filter(filterFunc func(any) bool) array {
	out := make([]any, len(a))
	for _, v := range a {
		if filterFunc(v) {
			out = append(out, filterFunc(v))
		}
	}
	return out
}

type player struct {
	name string
}

func main() {
	players := []player{
		{"Nicholas"},
	}

	fmt.Printf("players %v \n", players)

	var anyArray array
	for _, p := range players {
		anyArray = append(anyArray, p)
	}

	fmt.Printf("anyArray %v \n", anyArray)

	outArray := anyArray.Map(func(v any) any {
		return v.(player).name
	})

	fmt.Printf("outArray %v \n", outArray)

}
