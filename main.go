package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vidyasagar0405/go-varcall/bcftools"
	"github.com/vidyasagar0405/go-varcall/helptab"
	"github.com/vidyasagar0405/go-varcall/home"
	"github.com/vidyasagar0405/go-varcall/samtools"
)

type mainModel struct {
	activeTab        int
	tabs             []string
	homeTabModel     home.Model
	samtoolsTabModel samtools.Model
	bcftoolsTabModel bcftools.Model
	helpTabModel     helptab.Model
	keys             keymaps
	help             help.Model
}

func initialMainModel() mainModel {
	return mainModel{
		activeTab:        0,
		tabs:             []string{"Home", "Samtools", "Bcftools", "Help"},
		homeTabModel:     home.InitialModel(),
		samtoolsTabModel: samtools.InitialModel(),
		bcftoolsTabModel: bcftools.InitialModel(),
		helpTabModel:     helptab.InitialModel(),
		keys:             Keys,
		help:             help.New(),
	}
}

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

		case key.Matches(msg, m.keys.prevTab):
			m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
		}
	}

	var updatedModel tea.Model
	switch m.activeTab {
	case 0:
		updatedModel, cmd = m.homeTabModel.Update(msg)
		m.homeTabModel = updatedModel.(home.Model)
	case 1:
		updatedModel, cmd = m.samtoolsTabModel.Update(msg)
		m.samtoolsTabModel = updatedModel.(samtools.Model)
	case 2:
		updatedModel, cmd = m.bcftoolsTabModel.Update(msg)
		m.bcftoolsTabModel = updatedModel.(bcftools.Model)
	case 3:
		updatedModel, cmd = m.helpTabModel.Update(msg)
		m.helpTabModel = updatedModel.(helptab.Model)
	}
	return m, cmd
}

func (m mainModel) View() string {

	titleStyle := lipgloss.NewStyle().Padding(0, 3).Foreground(lipgloss.Color("205")).Bold(true)
	titleString := titleStyle.Render("VARCALL")

	displayString := ""

	tabContentstyle := lipgloss.NewStyle().Padding(0, 3)
	var tabContent string
	switch m.activeTab {
	case 0:
		tabContent = m.homeTabModel.View()
	case 1:
		tabContent = m.samtoolsTabModel.View()
	case 2:
		tabContent = m.bcftoolsTabModel.View()
	case 3:
		tabContent = m.helpTabModel.View()
	}
	helpView := lipgloss.NewStyle().Padding(0, 3).AlignVertical(lipgloss.Bottom).Render(m.help.View(m.keys))
	displayString += fmt.Sprintf("\n%s\n\n%v\n\n%s\n\n%v",
		titleString,
		m.tabView(),
		tabContentstyle.Render(tabContent),
		helpView)

	return displayString
}

func main() {
	app := tea.NewProgram(initialMainModel(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	_, err := app.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
