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

	m.state.logger.items = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc non ligula nec turpis porttitor maximus id in arcu. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Sed quam metus, suscipit eget neque vel, tincidunt faucibus lacus. Donec bibendum dignissim nulla, vel finibus mauris lacinia id. Proin sagittis diam vel libero tincidunt elementum. Vivamus odio quam, mattis id velit a, congue rhoncus quam. In hac habitasse platea dictumst. Suspendisse euismod, eros sit amet luctus sollicitudin, quam orci vestibulum dui, ac fermentum leo nibh ac lectus. Cras egestas id ante a accumsan. Proin fermentum accumsan purus in sagittis. Nam euismod felis dui, nec malesuada magna iaculis id. Donec ipsum purus, sollicitudin id eleifend elementum, sagittis sed sapien. Vivamus hendrerit id dui consequat mattis.",

		"Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos. Integer laoreet, mauris in egestas faucibus, mauris lacus finibus neque, sed semper velit tortor ut ipsum. Praesent tincidunt tellus vitae cursus gravida. Sed sed mi vitae erat pharetra porta. Nulla viverra tellus augue, et blandit lacus auctor non. In pretium ipsum gravida tristique accumsan. Aenean sit amet eleifend arcu. Suspendisse ac risus aliquet, cursus dolor nec, fringilla erat. Sed sollicitudin nulla nunc, ac fermentum massa dignissim eget. Cras leo lacus, condimentum vel varius eu, posuere et ligula. Pellentesque justo nisl, mollis a sem ut, venenatis euismod leo. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Praesent pulvinar ultrices imperdiet. Donec a ipsum vel velit sodales commodo.",

		"Integer vehicula est sed ullamcorper gravida. Mauris rhoncus, purus ac tempor ultricies, justo magna ultrices leo, sed finibus velit justo at nisl. Cras imperdiet purus est, at faucibus ligula rutrum lacinia. Sed vel felis porttitor, molestie est sit amet, venenatis massa. Etiam id vestibulum velit. Duis commodo dolor et blandit tristique. Sed in lacinia magna. Aliquam id nunc nunc. Mauris maximus venenatis turpis, vel semper leo imperdiet at. Donec sodales at magna non finibus. Proin eleifend lacinia luctus. Nulla venenatis non ligula eu molestie. Curabitur nec dictum odio, sed mollis libero. Etiam dignissim sollicitudin orci nec egestas. Donec lobortis tempus blandit. Fusce congue arcu ligula, pellentesque placerat velit egestas id.",
	}

	m.state.logger.folded = []bool{true, true, true}

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

		s += "\n\n"
	}

	return s
}
