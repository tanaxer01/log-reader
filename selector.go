package main

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type SelectorState struct {
	form     *huh.Form
	inpValue string
	selValue string
}

func (m Model) SelectorInit() tea.Cmd {
	m.state.selector.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Nombre:").
				Value(&m.state.selector.inpValue),
			huh.NewSelect[string]().
				Value(&m.state.selector.selValue).
				OptionsFunc(func() []huh.Option[string] {
					s := states[m.state.selector.inpValue]
					time.Sleep(1000 * time.Millisecond)
					return huh.NewOptions(s...)
				}, &m.state.selector.inpValue),
		),
	)

	return m.state.selector.form.Init()
}

var states = map[string][]string{
	"Canada": {
		"Alberta",
		"British Columbia",
		"Manitoba",
		"New Brunswick",
		"Newfoundland and Labrador",
		"North West Territories",
		"Nova Scotia",
		"Nunavut",
		"Ontario",
		"Prince Edward Island",
		"Quebec",
		"Saskatchewan",
		"Yukon",
	},
	"Mexico": {
		"Aguascalientes",
		"Baja California",
		"Baja California Sur",
		"Campeche",
		"Chiapas",
		"Chihuahua",
		"Coahuila",
		"Colima",
		"Durango",
		"Guanajuato",
		"Guerrero",
		"Hidalgo",
		"Jalisco",
		"México",
		"Mexico City",
		"Michoacán",
		"Morelos",
		"Nayarit",
		"Nuevo León",
		"Oaxaca",
		"Puebla",
		"Querétaro",
		"Quintana Roo",
		"San Luis Potosí",
		"Sinaloa",
		"Sonora",
		"Tabasco",
		"Tamaulipas",
		"Tlaxcala",
		"Veracruz",
		"Ignacio de la Llave",
		"Yucatán",
		"Zacatecas",
	},
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
