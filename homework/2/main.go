package main

import "fmt"

const appName = "GoFit"

func main() {
  name := "Иван"
	age  := 30
	height := 1.75
	subscribed := true 

	fmt.Println("Добро пожаловать в", appName+"!")
	fmt.Println("Профиль пользователя:")
	fmt.Println("Имя:", name)
	fmt.Println("Возраст:", age, "лет")
	fmt.Println("Рост:", height, "м")
	fmt.Println("Подписан на рассылку:", subscribed)
}