package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mistakenelf/teacup/markdown"
)

// Help tab model
type HelpModel struct {
	markdown markdown.Model
}

func InitialHelpModel() HelpModel {
	markdownModel := markdown.New(true, false, lipgloss.AdaptiveColor{Light: "#F25D93", Dark: "#F27D93"})
	markdownModel.FileName = "README.md"

	return HelpModel{
		markdown: markdownModel,
	}
}

func (m HelpModel) Init() tea.Cmd {
	return nil
}

func (m HelpModel) Update(msg tea.Msg) (HelpModel, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		cmds = append(cmds, m.markdown.SetSize(msg.Width, msg.Height))

		return m, tea.Batch(cmds...)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	m.markdown, cmd = m.markdown.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (hm HelpModel) View() string {
	return hm.markdown.View()
}
