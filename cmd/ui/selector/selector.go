package selector

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/styles"
)

const NO_SELECTION = "no valid selection made"

type Model struct {
	header   string
	options  []string
	cursor   int
	selected int
	err      error
}

func InitialSelectionModel(header string, options []string) Model {
	return Model{
		header:   header,
		options:  options,
		selected: -1,
		cursor:   0,
		err:      nil,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.options) - 1
			}

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.options) {
				m.cursor = 0
			}

		case " ":
			m.selected = m.cursor

		case "enter":
			_, m.err = m.handleSelection()
			if m.err != nil {
				break
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("\n%s\n\n", styles.HeaderStyle.Render(m.header)))
	for i, choice := range m.options {
		prefix := "( )"
		if i == m.selected {
			prefix = styles.SelectedCheckboxStyle.Render("(•)")
			choice = styles.SelectedTextStyle.Render(choice)
		}
		line := fmt.Sprintf("%s %s", prefix, choice)
		if i == m.cursor {
			s.WriteString(fmt.Sprintf("> %s\n", line))
		} else {
			line = styles.BlurStyle.Render(line)
			s.WriteString(fmt.Sprintf(" %s\n", line))
		}
	}
	helpOptions := "(Press [space] to select, enter to continue. Press q, esc, or ctrl-c to quit)"
	s.WriteString(fmt.Sprintf("\n%s\n", helpOptions))
	return s.String()
}

func RunSelector(header string, options []string) (int, string, error) {
	m0 := InitialSelectionModel(header, options)
	p := tea.NewProgram(m0)
	mFinal, err := p.Run()
	if err != nil {
		return -1, "", err
	}
	sel, ok := mFinal.(Model)
	if !ok {
		return -1, "", fmt.Errorf("unexpected model type %T", mFinal)
	}
	if sel.selected < 0 || sel.selected >= len(options) {
		return -1, "", fmt.Errorf(NO_SELECTION)
	}
	return sel.selected, options[sel.selected], nil
}

func (m *Model) handleSelection() (string, error) {
	for i := range m.options {
		if i == m.selected {
			return m.options[i], nil
		}
	}
	return "", fmt.Errorf("no options selected")
}
