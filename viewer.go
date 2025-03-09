package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type ViewerState struct {
	form          *huh.Form
	selValue      string
	posibleValues []string
}

func (m Model) ViewerSwitch() (Model, tea.Cmd) {
	m.page = viewerPage

	m.state.viewer.posibleValues = m.client.FetchStreams(m.state.selector.selValue)

	m.state.viewer.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Value(&m.state.viewer.selValue).
				Options(huh.NewOptions(m.state.viewer.posibleValues...)...),
		),
	)

	return m, m.state.viewer.form.Init()
}

func (m Model) ViewerUpdate(msg tea.Msg) (Model, tea.Cmd) {
	form, cmd := m.state.viewer.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.state.viewer.form = f
	}

	if m.state.viewer.form.State == huh.StateCompleted {
		return m.LoggerSwitch()
	}

	return m, cmd
}

func (m Model) ViewerView() string {
	return m.state.viewer.form.View()
}
