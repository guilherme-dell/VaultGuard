package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowSelectDirScreen(dirPathlabel *widget.Label, window *fyne.Window) {
	dialog := dialog.NewFolderOpen(func(list fyne.ListableURI, err error) {
		if err != nil || list == nil {
			return
		}

		vaultDir := list.String()

		if len(vaultDir) > 0 {
			vaultDirPath := strings.TrimPrefix(vaultDir, "file://")
			dirPathlabel.SetText(vaultDirPath)
			fmt.Println(vaultDirPath)
		}

	}, *window)

	dialog.Show()
}
