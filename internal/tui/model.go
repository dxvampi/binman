package tui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/dxvampi/binman/internal/store"
)

type mode int

const (
	listMode mode = iota
	aliasMode
	pathMode
)

type model struct {
	binaries   []store.Binary
	cursor     int
	mode       mode
	aliasInput textinput.Model
	pathInput  textinput.Model
	aliasTemp  string
	pathTemp   string
}

func initialModel() model {
	binaries, err := store.Load()
	if err != nil {
		binaries = []store.Binary{}
	}

	aliasInput := textinput.New()
	aliasInput.Placeholder = "alias"

	pathInput := textinput.New()
	pathInput.Placeholder = "/path/to/binary"

	return model{
		binaries:   binaries,
		cursor:     0,
		mode:       listMode,
		aliasInput: aliasInput,
		pathInput:  pathInput,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				store.Save(m.binaries)
				if m.cursor >= len(m.binaries) && m.cursor > 0 {
					m.cursor--
				}
			}

		case "c":
			m.mode = aliasMode
			m.aliasInput.Focus()
			return m, nil
		}
	}
	return m, nil
}

func (m model) updateAlias(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.mode = listMode
			m.aliasInput.SetValue("")
			return m, nil

		case "enter":
			m.aliasTemp = m.aliasInput.Value()
			m.aliasInput.SetValue("")
			m.aliasInput.Blur()
			m.mode = pathMode
			m.pathInput.Focus()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.aliasInput, cmd = m.aliasInput.Update(msg)
	return m, cmd
}

func (m model) updatePath(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.mode = aliasMode
			m.pathInput.SetValue("")
			return m, nil

		case "enter":
			m.pathTemp = m.aliasInput.Value()
			m.pathInput.SetValue("")
			m.pathInput.Blur()
			m.mode = listMode
			return m, nil
		}
	}
	var cmd tea.Cmd
	m.pathInput, cmd = m.pathInput.Update(msg)
	return m, cmd
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {
	case listMode:
		return m.updateList(msg)
	case aliasMode:
		return m.updateAlias(msg)
	case pathMode:
		return m.updatePath(msg)
	}
	return m, nil
}

func (m model) View() string {
	if len(m.binaries) == 0 {
		return "There's no configured binaries.\n\n(press 'q' to exit) (press 'c' to add one)\n"
	}

	if m.mode == aliasMode {
		return m.aliasInput.Prompt
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
