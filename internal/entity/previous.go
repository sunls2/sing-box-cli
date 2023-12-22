package entity

import "github.com/charmbracelet/bubbles/list"

type Previous struct {
	items []list.Item
	index int

	prev *Previous
}
