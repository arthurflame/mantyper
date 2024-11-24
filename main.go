package main

import (
	"fmt"
	"os"

	"github.com/arthurflame/mantyper/pkgs/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	// f, err := tea.LogToFile("debug.log", "debug")
	// if err != nil {
	// 	fmt.Printf("error - %v\n", err)
	// 	os.Exit(1)
	// }
	// defer f.Close()

	model := tui.Model{
		Choices:  []string{"Type", "Configure"},
		Selected: make(map[int]struct{}),
	}

	p := tea.NewProgram(model)
	m, err := p.Run()
	if err != nil {
		fmt.Printf("error - %v\n", err)
		os.Exit(1)
	}
	if m, ok := m.(tui.Model); ok && m.Choice != "" {
		fmt.Println("you have selected", m.Choice)
		switch m.Choice {
		case "Type":
			fmt.Println("run type view?")
		case "Configure":
			fmt.Println("run configure view?")
		}
	}
}
