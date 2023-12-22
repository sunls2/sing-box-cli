package entity

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	prev *Previous
	view *View

	main func() string
}

func NewApp(options []list.Item, main func() string) App {
	l := list.New(options, OptionDelegate{}, 20, 10)
	l.SetShowTitle(false)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	return App{
		view: NewView(l),
		main: main,
	}
}

func (a App) Init() tea.Cmd {
	return nil
}

func (a App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return a, tea.Quit
		case "enter", "right", "l":
			if a.view.model != nil {
				break
			}
			selected := a.view.list.SelectedItem().(Option)
			if selected.Next == nil && selected.Model == nil {
				return a, nil
			}
			a.prev = &Previous{prev: a.prev, items: a.view.list.Items(), index: a.view.list.Index()}
			if selected.Model != nil {
				a.view.model = selected.Model
				break
			}

			a.view.UpdateList(selected.Next, 0)
			return a, nil
		case "esc", "left", "h":
			if a.prev == nil {
				break
			}
			a.view.UpdateList(a.prev.items, a.prev.index)
			a.prev = a.prev.prev
			return a, nil
		}
	}
	return a, a.view.Update(msg)
}
func (a App) View() string {
	return a.main() + a.view.View()
}

var _ tea.Model = (*App)(nil)
