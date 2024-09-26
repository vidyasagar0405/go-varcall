package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/markdown"
)

// Help tab model
type helpModel struct {
	markdown markdown.Model
	content  string
}

func (hm helpModel) New() helpModel {
	markdownModel := markdown.New(true, false, lipgloss.AdaptiveColor{Light: "#0000ff", Dark: "#000099"})
	markdownModel.FileName = "../README.md"

	return helpModel{
		markdown: markdownModel,
	}
}

func initialHelpModel() helpModel {
	return helpModel{
		content: `
Help Content:
- Press 'tab' to switch between inputs
- Press 'q' or 'ctrl+c' to quit
`,
	}
}

func (m helpModel) Init() tea.Cmd {
	m.New()
	return nil
}

func (m helpModel) Update(msg tea.Msg) (helpModel, tea.Cmd) {
	return m, nil
}

func (hm helpModel) View() string {
	return hm.markdown.View()
}
