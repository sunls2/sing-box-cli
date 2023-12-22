package model

import (
	"sing-box-cli/config"
	"sing-box-cli/internal/constant"

	tea "github.com/charmbracelet/bubbletea"
)

type C1d1 struct {
	tip string
}

func (c C1d1) Init() tea.Cmd {
	return nil
}

func (c C1d1) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if ok {
		c.tip = "press " + key.String()
	}
	return c, nil
}

func (c C1d1) View() string {
	return constant.Indent + "C1D1 VIEW !!!\n" + constant.Indent + config.Cfg.BaseDir + "\n" + constant.Indent + c.tip
}

var _ tea.Model = (*C1d1)(nil)
