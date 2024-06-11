package main

import (
	"vaultGuard/internal/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()

	w := a.NewWindow("VaultGuard")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 800))

	dirPathLabel := widget.NewLabel("")

	selectDirButton := widget.NewButton("SELECT A VAULT TO BACKUP", func() { ui.ShowSelectDirScreen(dirPathLabel, &w) })

	menu := container.NewVBox(selectDirButton, dirPathLabel)

	w.SetContent(menu)

	w.ShowAndRun()

}
