package main

import ui "github.com/gizak/termui"

func main() {
	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	modalShown := false

	gameList := getList()
	// ui.Body.AddRows(
	// 	ui.NewRow(
	// 		ui.NewCol(6, 0, getList()),
	// 		ui.NewCol(6, 0),
	// 		ui.NewRow(
	// 			ui.NewCol(3, 0),
	// 			ui.NewCol(3, 0),
	// 			ui.NewCol(6, 0))))
	//ui.Body.AddRows(ui.NewCol(6, 0, getList()), ui.NewCol(6, 0, getList()))
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, gameList),
			ui.NewCol(6, 0, getDescription()),
		),
	)

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		nextIndex := gameList.SelectedItem() - 1
		if nextIndex < 0 {
			nextIndex = len(gameList.Items) - 1
		}
		gameList.SelectItem(nextIndex)
		ui.Render(ui.Body)
	})
	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		nextIndex := gameList.SelectedItem() + 1
		if nextIndex >= len(gameList.Items) {
			nextIndex = 0
		}
		gameList.SelectItem(nextIndex)
		ui.Render(ui.Body)
	})
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		if modalShown {
			ui.Clear()
			ui.Render(ui.Body)
			modalShown = false
		} else {
			modal := getModal("u sure?")
			ui.Render(ui.Body, modal)
			modalShown = true
		}
	})

	ui.Loop()
}

func getList() *SelectList {
	strs := []string{
		"Civilisation",
		"Descent",
		"Doom",
		"Doom 2",
		"Dune 2",
		"Quake",
		"Quarantine",
	}

	//ls := ui.NewList()
	ls := NewSelectList()
	ls.Items = strs
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "Games (installed)"
	ls.Height = 15
	ls.SelectItem(0)
	//ls.Width = 25
	ls.Y = 0
	return ls
}

func getDescription() ui.GridBufferer {
	par2 := ui.NewPar("Fucking [rad game](fg-bold) where Ghandi tries to rip you a new one")
	par2.Height = 10
	//par2.Width = 37
	//par2.Y = 4
	par2.BorderLabel = "Game Description"
	par2.BorderFg = ui.ColorYellow
	return par2
}

func getModal(text string) *ui.Par {
	p := ui.NewPar("Confirm")
	p.Height = 3
	p.Width = 20
	p.Y = 20
	p.Border = true
	p.Text = text
	return p
}
