package selector

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/styles"
)

type model struct {
	options  []string
	cursor   int
	selected int
	prompt   string
	exited   bool
	err      error
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			m.exited = true
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

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("\n%s\n\n", styles.PromptStyle.Render(m.prompt)))
	for i, choice := range m.options {
		prefix := "( )"
		if i == m.selected {
			prefix = styles.SelectedCheckboxStyle.Render("(•)")
			choice = styles.NormalTextStyle.Render(choice)
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
	s.WriteString(fmt.Sprintf("\n%s", helpOptions))
	return s.String()
}

func initialSelectionModel(prompt string, options []string) model {
	return model{
		options:  options,
		cursor:   0,
		selected: -1,
		prompt:   prompt,
		err:      nil,
	}
}

func (m *model) handleSelection() (string, error) {
	for i := range m.options {
		if i == m.selected {
			return m.options[i], nil
		}
	}
	return "", fmt.Errorf("no options selected")
}

func RunSelector(prompt string, options []string) (string, bool, error) {
	p := tea.NewProgram(initialSelectionModel(prompt, options))
	m, err := p.Run()
	if err != nil {
		return "", false, err
	}
	sel := m.(model)
	if sel.selected < 0 || sel.selected >= len(options) {
		return "", true, nil
	}
	return options[sel.selected], false, nil
}
