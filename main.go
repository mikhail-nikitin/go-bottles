package main

import "fmt"

func pluralizeBottles(num int) (result string) {
	if num == 0 {
		result = "Нет"
	} else {
		result = fmt.Sprintf("%d", num)
	}

	result += " бутыл"

	if num % 10 == 0 || num % 10 >= 5 || (num % 100 > 10 && num % 100 < 20) {
		result += "ок"
	} else if num % 10 == 1 {
		result += "ка"
	} else {
		result += "ки"
	}
	return
}

func main() {
	maxBottles := 1000
	for numBottles := maxBottles; numBottles > 0; numBottles-- {
		numBottlesString := pluralizeBottles(numBottles)
		fmt.Println(numBottlesString, "пива на стене")
		fmt.Println(numBottlesString, "пива!")
		fmt.Println("Возьми одну, пусти по кругу")

		fmt.Println(pluralizeBottles(numBottles - 1), "пива на стене!")
		fmt.Println()
	}

	fmt.Println("Нет бутылок пива на стене!")
	fmt.Println("Нет бутылок пива!", )
	fmt.Println("Пойди в магазин и купи ещё,")
	fmt.Println(pluralizeBottles(maxBottles), "пива на стене!")
}
