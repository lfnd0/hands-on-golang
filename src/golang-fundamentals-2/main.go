package main

import (
	"errors"
	"fmt"
)

func main() {
	// estruturas de controle

	// if, else if e else
	isCondition := "true"
	if 0 > 1 {
		result := fmt.Sprintf("0 > 1: %s", isCondition)
		fmt.Println(result)
	} else if 3 > 1 {
		result := fmt.Sprintf("1 < 3: %s", isCondition)
		fmt.Println(result)
	} else {
		fmt.Println("none")
	}

	// if com inicialização
	if hasReturn, error := ifWithInitialization(); error != nil {
		fmt.Println(hasReturn)
	}

	// switch, fallthrough, break e default
	// hasCase := "typescript"
	hasCase := "golang"
	switch hasCase {
	case "golang", "javascript":
		fmt.Println("golang")
		fmt.Println("javascript")
		fallthrough
	case "c-sharp":
		fmt.Println("c-sharp")
		fallthrough
	case "java":
		fmt.Println("java")
		// break // redundante
	case "python":
		fmt.Println("python")
	default:
		fmt.Println("none")
	}
}

func ifWithInitialization() (string, error) {
	return "golang is cool", errors.New("has error")
}
