package ui

import (
	"fmt"

	"github.com/RZhurakovskiy/proctl-cli/utils"
)

func ShowMainMenu() int {
	var action int

	fmt.Println("Главное меню:")
	menu := []MenuItem{
		{1, "Мониторинг и управление по загрузке CPU"},
		{2, "Мониторинг и управление по использованию памяти"},
		{0, "Выйти"},
	}

	for _, item := range menu {
		fmt.Printf(" [%d] %s\n", item.ID, item.Text)
	}
	fmt.Println("---------------------------------")
	action = utils.GetUserInput()
	return action
}
