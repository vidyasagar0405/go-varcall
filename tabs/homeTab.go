package tab

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type homeModel struct {
	projectNameInput textinput.Model
	outputFileInput  textinput.Model
}

func initialHomeModel() homeModel {
	projectNameInput := textinput.New()
	projectNameInput.Placeholder = "Enter project name"
	projectNameInput.Focus()

	outputFileInput := textinput.New()
	outputFileInput.Placeholder = "Enter output file name"

	return homeModel{
		projectNameInput: projectNameInput,
		outputFileInput:  outputFileInput,
	}
}

func (m homeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m homeModel) Update(msg tea.Msg) (homeModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.projectNameInput.Focused() {
				m.projectNameInput.Blur()
				m.outputFileInput.Focus()
			} else if m.outputFileInput.Focused() {
				m.outputFileInput.Blur()
			} else {
				m.projectNameInput.Focus()
			}
		}
	}

	m.projectNameInput, cmd = m.projectNameInput.Update(msg)
	m.outputFileInput, cmd = m.outputFileInput.Update(msg)

	return m, cmd
}

func (m homeModel) View() string {
	return fmt.Sprintf(
		"Project Name: %s\nOutput File: %s\n\n",
		m.projectNameInput.View(),
		m.outputFileInput.View(),
	)
}
