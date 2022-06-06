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
	if age == 1 ||
		age == 2 ||
		age == 3 ||
		age == 4 {
		thing = " год. "
	} else if age == 21 ||
		age == 22 ||
		age == 23 ||
		age == 24 ||
		age == 31 ||
		age == 32 ||
		age == 33 ||
		age == 34 ||
		age == 41 ||
		age == 42 ||
		age == 43 ||
		age == 44 ||
		age == 51 ||
		age == 52 ||
		age == 53 ||
		age == 54 ||
		age == 61 ||
		age == 62 ||
		age == 63 ||
		age == 64 ||
		age == 71 ||
		age == 72 ||
		age == 73 ||
		age == 74 ||
		age == 81 ||
		age == 82 ||
		age == 83 ||
		age == 84 ||
		age == 91 ||
		age == 92 ||
		age == 93 ||
		age == 94 {
		thing = " года. "
	} else if age > 100 {
		thing = " очень много лет или вы ошиблись"
	} else {
		thing = " лет. "
	}
	fmt.Println("Имя: " + name)
	fmt.Println("Возраст: " + strconv.Itoa(age))
	fmt.Println("Город: " + city)
	fmt.Println("Хобби: " + hobby)
	fmt.Println("Ваше имя: " + name + ". Вам " + strconv.Itoa(age) + thing + ". Ваш родной город: " + city + ". Вы любите заниматся: " + hobby)
}
