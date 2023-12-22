package entity

import (
	"fmt"
	"io"

	"sing-box-cli/internal/constant"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Option struct {
	Name  string
	Next  []list.Item
	Model tea.Model
}

func (opt Option) FilterValue() string {
	return opt.Name
}

var _ list.Item = (*Option)(nil)

type OptionDelegate struct{}

func (d OptionDelegate) Height() int                             { return 1 }
func (d OptionDelegate) Spacing() int                            { return 0 }
func (d OptionDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d OptionDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	str := fmt.Sprintf("%d. %s", index+1, item.(Option).Name)
	if index == m.Index() {
		str = "> " + str
	} else {
		str = constant.Indent + str
	}
	fmt.Fprint(w, str)
}
