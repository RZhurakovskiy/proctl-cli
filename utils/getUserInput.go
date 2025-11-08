package utils

import "fmt"

func GetUserInput() int {
	var choice int
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Ошибка! Введите номер пункта (цифру).")

		// оставлю с большой буквы, архитектурно на данный момент используется в других функциях другого пакета !!!проправаить потом!
		ClearScanBuffer()
		return -1
	}
	return choice
}
