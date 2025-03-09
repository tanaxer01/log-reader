package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type LoggerState struct {
	items  []string
	folded []bool
	cursor int
}

func (m Model) LoggerSwitch() (Model, tea.Cmd) {
	m.page = loggerPage

	m.state.logger.items = m.client.FetchEvents(m.state.selector.selValue, m.state.viewer.selValue)
	m.state.logger.folded = make([]bool, len(m.state.logger.items))

	return m, nil
}

func (m Model) LoggerUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			m.state.logger.cursor = (m.state.logger.cursor + 1) % len(m.state.logger.items)
		case "k":
			if m.state.logger.cursor > 0 {
				m.state.logger.cursor--
			} else {
				m.state.logger.cursor = len(m.state.logger.items) - 1
			}
		case "enter":
			m.state.logger.folded[m.state.logger.cursor] = !m.state.logger.folded[m.state.logger.cursor]
		}
	}
	return m, nil
}

func (m Model) LoggerView() string {
	if len(m.state.logger.items) == 0 {
		return "AAA"
	}

	s := ""
	loggerState := m.state.logger

	for i, para := range loggerState.items {
		if loggerState.folded[i] {
			cursor := ""
			if m.state.logger.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s", cursor, para[:80])
		} else {
			for j := 0; j < len(para); j += 80 {
				s += para[j:min(j+80, len(para))] + "\n"
			}
		}

		s += "\n"
	}

	return s
}
