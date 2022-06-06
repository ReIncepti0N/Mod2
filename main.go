package main

import (
	"fmt"
	"strconv"
)

func main() {
	var name, city, hobby, thing string
	var age int
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	fmt.Println("Введите возраст:")
	fmt.Scan(&age)
	fmt.Println("Введите город проживания:")
	fmt.Scan(&city)
	fmt.Println("Введите увлечения или хобби:")
	fmt.Scan(&hobby)
	switch age {
	case 1, 2, 3, 4, 21, 31, 41, 51, 61, 71, 81, 91:
		thing = " год. "
	case 22, 23, 24, 32, 33, 34, 42, 43, 44, 52, 53, 54, 62, 63, 64, 72, 73, 74, 82, 83, 84, 92, 93, 94:
		thing = " года. "
	default:
		thing = " лет. "
	}

	if age > 100 {
		thing = " очень много лет или вы ошиблись"
	}

	fmt.Printf("%s\n", "Имя: "+name)
	fmt.Println("Возраст: " + strconv.Itoa(age))
	fmt.Println("Город: " + city)
	fmt.Println("Хобби: " + hobby)
	fmt.Println("Ваше имя: " + name + ". Вам " + strconv.Itoa(age) + thing + "Ваш родной город: " + city + ". Вы любите заниматся: " + hobby)
}
