package main

import "fmt"

func main() {
	var name, city, hobby string
	var age int
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	fmt.Println("Введите возраст:")
	fmt.Scan(&age)
	fmt.Println("Введите город проживания:")
	fmt.Scan(&city)
	fmt.Println("Введите увлечения или хобби:")
	fmt.Scan(&hobby)
}
