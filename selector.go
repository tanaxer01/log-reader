package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type SelectorState struct {
	form          *huh.Form
	inpValue      string
	selValue      string
	posibleValues []string
}

func (m Model) SelectorInit() tea.Cmd {
	m.state.selector.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("LogGroup name:").
				Value(&m.state.selector.inpValue),
			huh.NewSelect[string]().
				Value(&m.state.selector.selValue).
				OptionsFunc(func() []huh.Option[string] {
					if len(m.state.selector.inpValue) > 5 {
						m.state.selector.posibleValues = m.client.FetchGroups(m.state.selector.inpValue)
					}

					return huh.NewOptions(m.state.selector.posibleValues...)
				}, &m.state.selector.inpValue),
		),
	)

	return m.state.selector.form.Init()
}

func (m Model) SelectorUpdate(msg tea.Msg) (Model, tea.Cmd) {
	form, cmd := m.state.selector.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.state.selector.form = f
	}

	if m.state.selector.form.State == huh.StateCompleted {
		return m.ViewerSwitch()
	}

	return m, cmd
}

func (m Model) SelectorView() string {
	return m.state.selector.form.View()
}

/*
func (m Model) SelectorSwitch() (Model, tea.Cmd) {
	m.page = selectorPage

	m.state.selector.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Selector"),
		),
	)

	return m, m.state.selector.form.Init()
}



*/
