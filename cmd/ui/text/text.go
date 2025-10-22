package text

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/tomasohCHOM/gdownloader/cmd/ui/styles"
)

type model struct {
	textInput textinput.Model
	prompt    string
	exited    bool
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
			m.exited = true
			return m, tea.Quit

		case tea.KeyEnter:
			if len(m.textInput.Value()) > 1 {
				m.textInput.Blur()
				return m, tea.Quit
			}
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"\n%s\n\n%s\n",
		styles.PromptStyle.Render(m.prompt),
		styles.NormalTextStyle.Render(m.textInput.View()),
	)
}

func initialTextModel(prompt string) model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40

	return model{
		prompt:    prompt,
		textInput: ti,
		exited:    false,
	}
}

func RunTextInput(prompt string) (string, bool, error) {
	p := tea.NewProgram(initialTextModel(prompt))
	m, err := p.Run()
	if err != nil {
		return "", false, err
	}
	final := m.(model)
	return final.textInput.Value(), final.exited, nil
}
