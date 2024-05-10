package main

import "fmt"

func main() {
	// variáveis
	var message string = "hello world" // menos utilizada
	fmt.Println(message)
	fmt.Printf("%T\n", message)

	text := "hello golang" // mais utilizada
	fmt.Println(text)

	age := 29
	fmt.Printf("%T\n", age)

	decimal := 14.1
	fmt.Printf("%T\n", decimal)

	dictionary := map[string]string{
		"message": "hands on golang",
	}
	fmt.Println(dictionary)
	fmt.Printf("%T\n", dictionary)

	// variáveis públicas e privadas
	myPrivateFunction() // inicial minúscula

	MyPublicFunction() // inicial maiúscula

	// interfaces
	any := map[string]interface{}{ // define um tipo "any"
		"name": "logan",
		"age":  16,
	}
	fmt.Println(any)

	// arrays e slices
	var fruits [3]string = [3]string{"apple", "orange", "banana"} // arrays têm tamanho definido
	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var names []string = []string{"logan", "spencer", "george"} // slices têm tamanho dinâmico
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))
	names = append(names, "gunner")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))

	// structs
	newUser := User{
		Name:     "robbie",
		Password: "1001001",
	}
	fmt.Println(newUser)
}

func myPrivateFunction() {
	fmt.Println("my private function")
}

func MyPublicFunction() {
	fmt.Println("my public function")
}

// define um objeto
type User struct {
	Name     string
	Password string
}
