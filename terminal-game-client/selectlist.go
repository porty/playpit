package main

import (
	"errors"
	"strings"

	ui "github.com/gizak/termui"
)

type SelectList struct {
	ui.Block
	Items         []string
	Overflow      string
	ItemFgColor   ui.Attribute
	ItemBgColor   ui.Attribute
	selectedIndex int
}

func NewSelectList() *SelectList {
	l := &SelectList{Block: *ui.NewBlock()}
	l.Overflow = "hidden"
	l.ItemFgColor = ui.ThemeAttr("list.item.fg")
	l.ItemBgColor = ui.ThemeAttr("list.item.bg")
	return l
}

// Buffer implements Bufferer interface.
func (l *SelectList) Buffer() ui.Buffer {
	buf := l.Block.Buffer()

	switch l.Overflow {
	case "wrap":
		str := strings.Join(l.Items[0:l.selectedIndex], "\n") +
			"[" + l.Items[l.selectedIndex] + "](fg-bold)\n" +
			strings.Join(l.Items[l.selectedIndex+1:], "\n")
		cs := ui.DefaultTxBuilder.Build(str, l.ItemFgColor, l.ItemBgColor)
		i, j, k := 0, 0, 0
		innerArea := l.InnerBounds()
		for i < innerArea.Dy() && k < len(cs) {
			w := cs[k].Width()
			if cs[k].Ch == '\n' || j+w > innerArea.Dx() {
				i++
				j = 0
				if cs[k].Ch == '\n' {
					k++
				}
				continue
			}
			buf.Set(innerArea.Min.X+j, innerArea.Min.Y+i, cs[k])

			k++
			j++
		}

	case "hidden":
		trimItems := l.Items
		innerArea := l.InnerBounds()
		if len(trimItems) > innerArea.Dy() {
			trimItems = trimItems[:innerArea.Dy()]
		}
		for i, v := range trimItems {
			if i == l.selectedIndex {
				v = "[" + v + "](fg-bold)"
			}
			cs := ui.DTrimTxCls(ui.DefaultTxBuilder.Build(v, l.ItemFgColor, l.ItemBgColor), innerArea.Dx())
			j := 0
			for _, vv := range cs {
				w := vv.Width()
				buf.Set(innerArea.Min.X+j, innerArea.Min.Y+i, vv)
				j += w
			}
		}
	}
	return buf
}

func (l *SelectList) SelectItem(i int) error {
	if i >= len(l.Items) {
		return errors.New("Index out of range")
	}
	l.selectedIndex = i
	return nil
}

func (l *SelectList) SelectedItem() int {
	return l.selectedIndex
}
