package main

import (
	"fmt"
	"math/rand"
)

// Компилятор Go не требует обязательного объявления типа данных.

// Константа
const secondsPerDay = 86400

func main() {
	// Объявляем три переменные (динамическая типизация)
	distance := 62100000
	planet := ""
	trip := ""

	// Новая строка: заголовок таблицы
	fmt.Printf("%-10v %-5v %-12v %5v\n", "Planet", "Days", "Trip type", "Price")
	// Новая строка: разделитель
	fmt.Println("======================================")

	/**
	В цикле создаем 10 строк
	*/
	for count := 0; count < 10; count++ {
		// Массив с названием планет
		// Инициализация массива через композитные литералы
		planets := [5]string{"Церера", "Плутон", "Юпитер", "Марс", "Эрида"}

		// Планета из массива по случайному индексу
		planet = planets[rand.Intn(3)]

		// Случайно выбираем скорость
		speed := rand.Intn(15) + 16 // 16-30 km/s

		// Рассчитываем продолжительность
		duration := distance / speed / secondsPerDay // days

		// Рассчитываем прайс
		price := 20.0 + speed // millions

		// Случайно выбираем тип поездки
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}

		// Вывод строк таблицы в цикле
		fmt.Printf("%-10v %-5v %-12v $%-3v\n", planet, duration, trip, price)
		fmt.Println("--------------------------------------")
	}

	// Делаем паузу
	var input string
	fmt.Scanf("%v", &input)
}
