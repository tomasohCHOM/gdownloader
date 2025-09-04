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
	header    string
	err       error
	errMsg    string
}

func InitialTextModel(header string, errMsg string) model {
	ti := textinput.New()
	ti.Focus()
	return model{
		textInput: ti,
		header:    header,
		err:       nil,
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
		styles.HeaderStyle.Render(m.header),
		styles.SelectedTextStyle.Render(m.textInput.View()),
	)
}
