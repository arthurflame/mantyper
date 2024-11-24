package tui

func InitialModel() Model {
	return Model{
		Choices:  []string{"Type", "Configure"},
		Selected: make(map[int]struct{}),
	}
}
