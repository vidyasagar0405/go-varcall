package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type mainModel struct {
	activeTab        int
	tabs             []string
	homeTabModel     homeModel
	samtoolsTabModel samtoolsModel
	bcftoolsTabModel bcftoolsModel
	help             help.Model
	keys             keymaps
}

func initialMainModel() mainModel {
	return mainModel{
		activeTab:        0,
		tabs:             []string{"Home", "Samtools", "Bcftools", "Help"},
		homeTabModel:     homeModel{},
		samtoolsTabModel: samtoolsModel{},
		bcftoolsTabModel: bcftoolsModel{},
		keys:             Keys,
		help:             help.New(),
	}
}

type homeModel struct{}
type samtoolsModel struct{}
type bcftoolsModel struct{}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	cmd = nil

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.quit):
			cmd = tea.Quit

		case key.Matches(msg, m.keys.nextTab):
            m.activeTab = (m.activeTab + 1) % len(m.tabs)

		case key.Matches(msg, m.keys.nextTab):
            m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
		}
	}

	return m, cmd
}

func (m mainModel) View() string {

	displayString := "\n"

	displayString += m.tabView()

	return displayString
}

func main() {
	app := tea.NewProgram(initialMainModel(), tea.WithAltScreen())
	_, err := app.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
