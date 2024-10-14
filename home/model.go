package home

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	projectName textinput.Model
	keys        keymaps
}

func InitialModel() Model {

	projectName := textinput.New()
	projectName.Placeholder = "Project name"

	return Model{
		projectName: projectName,
		keys:        Keys,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.quit):
			cmd = tea.Quit
		}
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {

	headerString := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6b53de"))

	displayString := "\n"

	displayString += fmt.Sprintf("%s%s\n\n%s\n\n%s\n\n%s\n",
		headerString.Render("Project name: "),
		m.projectName.View(),
		headerString.Render("FastQC"),
		headerString.Render("Multiqc"),
		headerString.Render("BWA"),
	)

	return displayString
}
