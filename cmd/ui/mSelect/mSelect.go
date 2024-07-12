package mSelect

import (
    "fmt"
	"github.com/charmbracelet/lipgloss"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/JakeNorman007/klarah/cmd/step"
    "github.com/JakeNorman007/klarah/cmd/program"
)

var (
    focusStyle = lipgloss.NewStyle()
    titleStyle = lipgloss.NewStyle()
    selectedItemStyle = lipgloss.NewStyle()
    selectedItemDescriptionStyle = lipgloss.NewStyle()
    descriptionStyle = lipgloss.NewStyle()
)

type Selection struct {
    Choices map[string]bool
}

func (s *Selection) Update(optionName string, value bool) {
    s.Choices[optionName] = value
}

type model struct {
    cursor      int
    options     []step.Item
    selected    map[int]struct{}
    choice      *Selection
    header      string
    exit        *bool
}

func (m model) Init() tea.Cmd {
    return nil
}

func InitModelSelect(options []step.Item, selection *Selection, header string, program *program.Project) model {
    return model {
        options:    options,
        selected:   make(map[int]struct{}),
        choice:     selection,
        header:     titleStyle.Render(header),
        exit:       &program.Exit,
    }
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            *m.exit = true

            return m, tea.Quit

        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.options) - 1 {
                m.cursor++
            }
        case "enter", " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        case "y":
            for selectedKey := range m.selected {
                m.choice.Update(m.options[selectedKey].Flag, true)
                m.cursor = selectedKey
            }

            return m, tea.Quit
        }
    }
    
    return m, nil
}

func (m model) View() string {
    s := m.header + "\n\n"

    for i, option := range m.options {
        cursor := " "
        if m.cursor == i {
            cursor = focusStyle.Render(">")
            option.Title = selectedItemStyle.Render(option.Title)
            option.Description = selectedItemStyle.Render(option.Description)
        }

        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = focusStyle.Render("*")
        }

        title := focusStyle.Render(option.Title)
        description := focusStyle.Render(option.Description)

        s += fmt.Sprintf("%s (%s) %s\n%s\n\n", cursor, checked, title, description)
    }

    s += fmt.Sprintf("Press %s to confirm your choice.\n\n", focusStyle.Render("y"))

    return s
}
