package main

import (
	"fmt"
	"log-scroller/fetcher"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type page int

const (
	selectorPage page = iota
	viewerPage
	loggerPage
)

type Model struct {
	page     page
	pageCant page
	//
	ready    bool
	viewport viewport.Model
	client   *fetcher.AwsFetcher
	state    *State
}

type State struct {
	selector SelectorState
	viewer   ViewerState
	logger   LoggerState
}

func NewModel() Model {
	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		profile = "default"
	}

	m := Model{
		page:   selectorPage,
		state:  &State{},
		client: fetcher.NewAwsFetcher(profile),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return m.SelectorInit()
	// var cmd tea.Cmd
	// m, cmd = m.LoggerSwitch()

	// return cmd
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height
		}
	}

	switch m.page {
	case selectorPage:
		m, cmd = m.SelectorUpdate(msg)
	case viewerPage:
		m, cmd = m.ViewerUpdate(msg)
	case loggerPage:
		m, cmd = m.LoggerUpdate(msg)
	}

	cmds = append(cmds, cmd)

	m.viewport.SetContent(m.getContent())
	m.viewport, cmd = m.viewport.Update(msg)

	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	return m.viewport.View()
}

func (m Model) getContent() string {
	page := "unknown"

	switch m.page {
	case selectorPage:
		page = m.SelectorView()
	case viewerPage:
		page = m.ViewerView()
	case loggerPage:
		page = m.LoggerView()
	}

	return page
}

func main() {
	p := tea.NewProgram(NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
