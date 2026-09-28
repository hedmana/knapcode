package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	lastKey string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
		m.lastKey = msg.String()
	}
	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView(fmt.Sprintf("Bubbletea works! Last key: %q\n\nPress q to quit.", m.lastKey))
	v.AltScreen = true
	return v
}

func main() {
	if _, err := tea.NewProgram(model{}).Run(); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
