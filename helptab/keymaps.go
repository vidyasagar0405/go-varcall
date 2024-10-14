package helptab

import "github.com/charmbracelet/bubbles/key"

type keymaps struct {
	quit    key.Binding
	nextTab key.Binding
	prevTab key.Binding
}

var Keys = keymaps{
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"), // actual keybindings
		key.WithHelp("q", "quit"),   // corresponding help text
	),
	nextTab: key.NewBinding(
		key.WithKeys("tab"),             // actual keybindings
		key.WithHelp("tab", "next tab"), // corresponding help text
	),
	prevTab: key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "previous tab"),
	),
}
