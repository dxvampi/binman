package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/dxvampi/binman/internal/store"
)

type model struct {
	binaries []store.Binary
	cursor   int
}

func initialModel() model {
	binaries, err := store.Load()
	if err != nil {
		binaries = []store.Binary{}
	}

	return model{
		binaries: binaries,
		cursor:   0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.binaries)-1 {
				m.cursor++
			}
		case "d":
			if len(m.binaries) > 0 {
				m.binaries = append(m.binaries[:m.cursor], m.binaries[m.cursor+1:]...)

				err := store.Save(m.binaries)
				if err != nil {
				}

				if m.cursor >= len(m.binaries) && m.cursor > 0 {
					m.cursor--
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	if len(m.binaries) == 0 {
		return "There's no configured binaries.\n\n(press 'q' to exit)\n"
	}

	s := "Your binaries:\n\n"

	for i, bin := range m.binaries {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s -> %s\n", cursor, bin.Alias, bin.Path)
	}

	s += "\n(press 'q' to exit) (press 'd' to remove an alias)\n"
	return s
}

func Run() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
