package main

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
		key.WithKeys("ctrl+n"),             // actual keybindings
		key.WithHelp("ctrl+n", "next tab"), // corresponding help text
	),
	prevTab: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "previous tab"),
	),
}
