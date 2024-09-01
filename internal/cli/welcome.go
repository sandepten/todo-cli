package cli

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Welcome to the todo CLI! (in ASCII art)
const Logo = `
Yb        dP 888888 88      dP""b8  dP"Yb  8b    d8 888888   888888  dP"Yb    888888 88  88 888888   888888  dP"Yb  8888b.   dP"Yb     dP""b8 88     88
 Yb  db  dP  88__   88     dP   '" dP   Yb 88b  d88 88__       88   dP   Yb     88   88  88 88__       88   dP   Yb  8I  Yb dP   Yb   dP   '" 88     88
  YbdPYbdP   88""   88  .o Yb      Yb   dP 88YbdP88 88""       88   Yb   dP     88   888888 88""       88   Yb   dP  8I  dY Yb   dP   Yb      88  .o 88
   YP  YP    888888 88ood8  YboodP  YbodP  88 YY 88 888888     88    YbodP      88   88  88 888888     88    YbodP  8888Y"   YbodP     YboodP 88ood8 88
`

// print the welcome logo when the CLI is started
func Welcome() string {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Bold(true).Render(Logo)
}

// welcomeModel is the initial model for the Bubble Tea program
type welcomeModel struct {
	textInput textinput.Model
}

// Init is the initial command we run when the program starts
func (m welcomeModel) Init() tea.Cmd {
	return tea.Batch(
		tea.EnterAltScreen,
		textinput.Blink,
	)
}

// Update is called when messages are received
func (m welcomeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, tea.Quit
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View renders the welcome logo and the text input prompt
func (m welcomeModel) View() string {
	return fmt.Sprintf(
		"\n%s\n\nHey, you are a new user? Could you tell me your name?\n\n%s\n",
		Welcome(),
		m.textInput.View(),
	)
}

// NewWelcomeModel creates a new welcome model
func NewWelcomeModel() welcomeModel {
	ti := textinput.New()
	ti.Placeholder = "Enter your name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return welcomeModel{
		textInput: ti,
	}
}
