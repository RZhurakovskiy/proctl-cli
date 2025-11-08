package main

import (
	"fmt"

	"github.com/RZhurakovskiy/proctl-cli/cpu"
	"github.com/RZhurakovskiy/proctl-cli/ui"
)

func main() {
	ui.ShowBanner()

	action := ui.ShowMainMenu()

	switch action {
	case 1:
		cpu.ProcessMenu()
	case 2:
		fmt.Println("Раздел в разработке")
		return
	case 0:
		fmt.Println("Выход...")
		return
	}

}
