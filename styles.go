package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m mainModel) tabView() string {
	var tabViews []string
	for i, tab := range m.tabs {
		style := lipgloss.NewStyle().Padding(0, 4).Border(lipgloss.NormalBorder(), false, true, false, false)
		// separator := lipgloss.NewStyle().Foreground(lipgloss.Color("#caa5f6"))
		if i == m.activeTab {
			style = style.Bold(true).Foreground(lipgloss.Color("205"))
		}
		tabViews = append(tabViews, style.Render(tab))
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, tabViews...)
}

type DefStyles struct {
	defaultWithPadd lipgloss.Style
}

var (
	defaultWithPadd     = lipgloss.NewStyle().Padding(0, 3)

	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)
