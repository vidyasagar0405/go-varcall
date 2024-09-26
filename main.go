package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gituhb.com/vidyasagar0450/go-varcall/tabs"
)

type errMsg error

type model struct {
	samtoolsTabModel *tabs.samtoolsModel
	bcftoolsTabModel *tabs.bcftoolsModel
	help             help.Model
	err              error
	helpTabModel     *tabs.helpModel
	state            *tabs.SessionState
	Keys             KeyMap
	tabs             []string
	homeTabModel     *tabs.homeModel
	textInput        textinput.Model
	spinner          spinner.Model
	activeTab        int
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Points
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{
		Keys:    Keys,
		help:    help.New(),
		state:   *tabs.SessionState{},
		spinner: s,

		tabs:             []string{"Home", "Samtools", "Bcftools", "Help"},
		activeTab:        0,
		homeTabModel:     *tabs.initialHomeModel(),
		samtoolsTabModel: *tabs.initialSamtoolsModel(),
		bcftoolsTabModel: *tabs.initialBcftoolsModel(),
		helpTabModel:     *tabs.initialHelpModel(),
	}
}

func (m model) tabView() string {
	var tabViews []string
	for i, tab := range m.tabs {
		style := lipgloss.NewStyle().Padding(0, 1, 0, 3)
		if i == m.activeTab {
			style = style.Bold(true).Foreground(lipgloss.Color("205"))
		}
		tabViews = append(tabViews, style.Render(tab))
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, tabViews...)
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.quit):
			return m, tea.Quit
		case key.Matches(msg, m.Keys.nextTab):
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
		case key.Matches(msg, m.Keys.prevTab):
			m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
		case key.Matches(msg, m.Keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		}
	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	}

	var cmd tea.Cmd
	switch m.activeTab {
	case 0:
		m.homeTabModel, cmd = m.homeTabModel.Update(msg)
		m.state = m.homeTabModel.updateSessionState(m.state)
	case 1:
		m.samtoolsTabModel, cmd = m.samtoolsTabModel.Update(msg)
	case 2:
		m.bcftoolsTabModel, cmd = m.bcftoolsTabModel.Update(msg)
	case 3:
		m.helpTabModel, cmd = m.helpTabModel.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	titleStyle := lipgloss.NewStyle().
		MarginLeft(2).
		Padding(1).
		Bold(true).
		Border(lipgloss.HiddenBorder(), false).
		Foreground(lipgloss.Color("#F25D93")).
		SetString("varcall")

	tabContent := ""
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
	helpView := m.help.View(m.Keys)
	return fmt.Sprintf("%v\n\n%s\n\n%s\n\n%v\n\n%v", titleStyle, m.tabView(), tabContent, m.spinner.View(), helpView)
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
