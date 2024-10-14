package samtools

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
type Model struct {
    keys keymaps
}

func InitialModel() Model {
    return Model{
        keys: Keys,
    }
}

func (m Model) Init() tea.Cmd {
    return nil
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

    displayString += fmt.Sprintf("%s\n\n%s\n\n%s\n\n", headerString.Render("view"), headerString.Render("sort"), headerString.Render("index"))

    return displayString
}
