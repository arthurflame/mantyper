package main

import (
	"fmt"
	"os"

	"github.com/arthurflame/mantyper/pkgs/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(tui.InitialModel())
	_, err := p.Run()
	if err != nil {
		fmt.Printf("error - %v\n", err)
		os.Exit(1)
	}
}
