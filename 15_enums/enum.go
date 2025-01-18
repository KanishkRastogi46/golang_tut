package main

import "fmt"

// Define an enum
type Languages string

const (
	Golang Languages = "golang"
	Python Languages = "python"
	Java Languages = "java"
	Javascript Languages = "javascript"
	Typescript Languages = "typescript"
	Cpp Languages = "cpp"
	Csharp Languages = "csharp"
)

func favLang (lang Languages) {
	fmt.Println("My favourite language is", lang)
}

func main() {
	fmt.Println("Enums in Go")
	favLang(Golang)
	favLang(Python)
	favLang(Typescript)
	favLang(Cpp)
	favLang(Csharp)
}