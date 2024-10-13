package main

import "github.com/charmbracelet/lipgloss"


func (m mainModel) tabView() string {
	var tabViews []string
	for i, tab := range m.tabs {
		style := lipgloss.NewStyle().Padding(0, 4)
        separator := lipgloss.NewStyle().Foreground(lipgloss.Color("#caa5f6"))
		if i == m.activeTab {
			style = style.Bold(true).Foreground(lipgloss.Color("205"))
		}
		tabViews = append(tabViews, style.Render(tab), separator.Render("|"))
	}
	return lipgloss.JoinHorizontal(lipgloss.Left, tabViews...)
}


type DefStyles struct{
    defaultWithPadd lipgloss.Style
}
