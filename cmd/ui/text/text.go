package text

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/styles"
)

type errMsg error

type model struct {
	textInput textinput.Model
	prompt    string
	done      bool
	errMsg    string
	err       error
}

func InitialTextModel(prompt string) model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156 // arbitrary, adjust as needed
	ti.Width = 40

	return model{
		prompt:    prompt,
		textInput: ti,
		done:      false,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			// m.userOptions.ExitState = true
			return m, tea.Quit

		case tea.KeyEnter:
			m.textInput.Blur()
			// m.userOptions.Username = m.textInput.Value()
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n\n%s",
		styles.DimStyle.Render(m.errMsg),
		styles.HeaderStyle.Render(m.prompt),
		styles.SelectedTextStyle.Render(m.textInput.View()),
	)
}

func RunTextInput(prompt string) (string, error) {
	p := tea.Program(InitialTextModel(prompt))
	m, err := p.Run()
	if err != nil {
		return "", err
	}
	final := m.(Model)
	return final.textInput.Value(), nil
}
