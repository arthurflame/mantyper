package tui

func InitialModel() model {
	return model{
		choices:  []string{"Type", "Configure"},
		selected: make(map[int]struct{}),
	}
}
