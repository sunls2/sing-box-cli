package app

import (
	"fmt"
	"os"

	"sing-box-cli/config"
	"sing-box-cli/internal/constant"
	"sing-box-cli/internal/entity"
	"sing-box-cli/internal/model"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var options = []list.Item{
	entity.Option{
		Name: "启动",
	},
	entity.Option{
		Next: []list.Item{
			entity.Option{
				Name: "C1",
				Next: []list.Item{
					entity.Option{Name: "c1d1", Model: model.C1d1{}},
					entity.Option{Name: "c1d2"},
				},
			},
			entity.Option{
				Name: "C2",
				Next: []list.Item{
					entity.Option{Name: "c2d1"},
					entity.Option{Name: "c2d2"},
					entity.Option{Name: "c2d3"},
				},
			},
			entity.Option{
				Name: "C3",
				Next: []list.Item{
					entity.Option{Name: "c3d1"},
					entity.Option{Name: "c3d2"},
					entity.Option{Name: "c3d3"},
					entity.Option{Name: "c3d4"},
				},
			},
		},
		Name: "配置",
	},
}

func appView() string {
	return fmt.Sprintf("\n%sSING-BOX: 已开启\n%sCONFIG: %+v\n\n", constant.Indent, constant.Indent, config.Cfg)
}

func Run() {
	p := tea.NewProgram(entity.NewApp(options, appView), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
