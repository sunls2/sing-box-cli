package entity

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type View struct {
	list  list.Model
	model tea.Model
}

func NewView(list list.Model) *View {
	return &View{list: list}
}

func (v *View) Update(msg tea.Msg) (cmd tea.Cmd) {
	if v.model != nil {
		v.model, cmd = v.model.Update(msg)
		return
	}
	v.list, cmd = v.list.Update(msg)
	return
}

func (v *View) UpdateList(items []list.Item, index int) {
	v.list.SetItems(items)
	v.list.Select(index)
	v.model = nil
}

func (v *View) View() string {
	if v.model != nil {
		return v.model.View()
	}
	return v.list.View()
}
