package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	quit    key.Binding
	nextTab key.Binding
	prevTab key.Binding
	Help    key.Binding
	enter   key.Binding
}

var Keys = KeyMap{
	quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"), // actual keybindings
		key.WithHelp("q", "quit"),   // corresponding help text
	),
	nextTab: key.NewBinding(
		key.WithKeys("ctrl+n"),
		key.WithHelp("ctrl+n", "next tab"),
	),
	prevTab: key.NewBinding(
		key.WithKeys("ctrl+p"),
		key.WithHelp("ctrl+p", "previous tab"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "enter input"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.nextTab, k.prevTab}, // first column
		{k.Help, k.quit},       // second column
	}
}
