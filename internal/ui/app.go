package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func changeTextBox(label *widget.Label, text string) {
	label.SetText(text)
}

func viewContainer(container *fyne.Container) {
	fmt.Println(container.Size())
}

func menuConfig() {

	hello := widget.NewLabel("Hello fyne")
	buttonVaultFolderPath := widget.NewButton("FIND VAULT FOLDER", func() { changeTextBox(hello, "VAULT SELECTED") })
	containerHello := container.NewVBox(hello, buttonVaultFolderPath)

	viewContainer(containerHello)
}
