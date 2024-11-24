package tui

import (
	"strings"
)

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString("What should we do?\n\n")

	for i := range m.Choices {
		if m.Cursor == i {
			s.WriteString("(*) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.Choices[i])
		s.WriteString("\n")
	}

	s.WriteString("\nPress q to quit.\n")
	return s.String()
}
