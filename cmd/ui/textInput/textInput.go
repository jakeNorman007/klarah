package textInput

import (
	"fmt"
	"errors"
	"regexp"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/JakeNorman007/klarah/cmd/program"
)

var (
    titleStyle = lipgloss.NewStyle()
    errorStyle = lipgloss.NewStyle()
)

type errMsg error

type Output struct {
    Output string
}

func (o *Output) update(value string) {
    o.Output = value
}

type model struct {
    textInput textinput.Model
    err       error
    output    *Output
    header    string
    exit      *bool
}

func sanitizeInput(input string) error {
    match, err := regexp.Match("^[a-zA-Z0-9_-]+$", []byte(input))
    if !match {
        return fmt.Errorf("String does not match regex pattern, err: %v", err)
    }

    return nil
}

func InitTextInputModel(output *Output, header string, program *program.Project) model {
    txtInp := textinput.New()
    txtInp.Focus()
    txtInp.CharLimit = 156
    txtInp.Width = 20
    txtInp.Validate = sanitizeInput

    return model {
        textInput: txtInp,
        err:       nil,
        output:    output,
        header:    titleStyle.Render(header),
        exit:      &program.Exit,
    }
}

func CreateErrorInputModel(err error) model {
    txtInp := textinput.New()
    txtInp.Focus()
    txtInp.CharLimit = 156
    txtInp.Width = 20
    exit := true

    return model {
        textInput: txtInp,
        err:       errors.New(errorStyle.Render(err.Error())),
        output:    nil,
        header:    "", 
        exit:      &exit,
    }

}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter:
            if len(m.textInput.Value()) > 1 {
                m.output.update(m.textInput.Value())
                
                return m, tea.Quit
            }
        case tea.KeyCtrlC, tea.KeyEsc:
            *m.exit = true

            return m, tea.Quit
        }
    case errMsg:
        m.err = msg
        *m.exit = true

        return m, nil
    }

    m.textInput, cmd = m.textInput.Update(msg)

    return m, cmd
}

func (m model) View() string {
    return fmt.Sprintf("%s%s\n\n", m.header, m.textInput.View())
}

func (m model) Err() string {
    return m.err.Error()
}
