package tabs

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type HomeModel struct {
	err           error
	projectNameIn textinput.Model
	outputFileIn  textinput.Model
}

func InitialHomeModel() HomeModel {
	projectNameIn := textinput.New()
	projectNameIn.Placeholder = "Enter project name"
	projectNameIn.Focus()

	inputFileIn := textinput.New()
	inputFileIn.Placeholder = "Enter output file name"

	return HomeModel{
		projectNameIn: projectNameIn,
		outputFileIn:  inputFileIn,
		err:           nil,
	}
}

type SessionState struct {
	WorkingDir     string
	OutputFileName string
	OutputDirName  string
	FullOutputPath string
}

func (m HomeModel) UpdateSessionState(state SessionState) SessionState {
	state.WorkingDir = m.projectNameIn.Value()
	state.OutputFileName = m.outputFileIn.Value()
	state.FullOutputPath = fmt.Sprintf("%s/data/%s/%s", state.WorkingDir, state.OutputDirName, state.OutputFileName)
	return state
}

func (m HomeModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m HomeModel) Update(msg tea.Msg) (HomeModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.projectNameIn.Focused() {
				m.projectNameIn.Blur()
				m.outputFileIn.Focus()
			} else if m.outputFileIn.Focused() {
				m.outputFileIn.Blur()
			} else {
				m.projectNameIn.Focus()
			}
		}
	}

	m.projectNameIn, cmd = m.projectNameIn.Update(msg)
	m.outputFileIn, cmd = m.outputFileIn.Update(msg)

	return m, cmd
}

func (m HomeModel) View() string {
	return fmt.Sprintf(
		"Project Name: %s\nOutput File: %s\n\n",
		m.projectNameIn.View(),
		m.outputFileIn.View(),
	)
}
