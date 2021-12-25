package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"

	// "fyne.io/fyne/v2/theme"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func showGallery(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Gallary")
	// w.Resize(fyne.Size{1200, 800})
	var pics []string

	files, err := ioutil.ReadDir("/mnt/data/Photos/iMAGES")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() == false {
			s := strings.Split(f.Name(), ".")[1]
			if s == "png" || s == "jpg" {
				pics = append(pics, "/mnt/data/Photos/iMAGES"+"/"+f.Name())
			}
		}
	}
	tabs := container.NewAppTabs(
		container.NewTabItem(strings.Split(pics[0], "iMAGES/")[1], canvas.NewImageFromFile(pics[0])),
	)
	for i := 1; i < len(pics); i++ {
		p := strings.Split(pics[i], "iMAGES/")[1]
		tabs.Append(container.NewTabItem(p, canvas.NewImageFromFile(pics[i])))
	}
	// tabs.SetTabLocation(container.TabLocationLeading)
	// s:=tabs.Items
	imgContainer := container.NewVBox(
		container.NewVBox(
			tabs,
		),
	)

	// w.Resize(fyne.NewSize(500, 280))
	w.SetContent(container.NewBorder(Deskbtn, nil, nil, nil, imgContainer))
	// w.Show()
	// w.SetContent(tabs)
	w.Show()
}
