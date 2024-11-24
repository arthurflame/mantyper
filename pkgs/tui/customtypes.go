package tui

type Model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
	Choice   string
}
