package home

import (
	"github.com/charmbracelet/bubbles/key"
)

func (k keymaps) ShortHelp() []key.Binding {
	return []key.Binding{k.quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keymaps) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.nextTab, k.prevTab}, // first column
		{k.quit},       // second column
	}
}
