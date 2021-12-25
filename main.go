package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myapp fyne.App = app.New()
var mywindow fyne.Window = myapp.NewWindow("Virtual Desktop")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var Deskbtn fyne.Widget
var img fyne.CanvasObject
var panelContent *fyne.Container

func main() {
	myapp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("/mnt/data/Photos/iMAGES/gaurav.jpg")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		showWeatherApp(mywindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator", theme.ContentAddIcon(), func() {
		showCalc()
	})

	btn3 = widget.NewButtonWithIcon("Gallery", theme.StorageIcon(), func() {
		showGallery(mywindow)
	})

	btn4 = widget.NewButtonWithIcon("TextEditor", theme.DocumentIcon(), func() {
		showTextEditor(mywindow)
	})

	Deskbtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		mywindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(5, Deskbtn, btn1, btn2, btn3, btn4)))

	mywindow.Resize(fyne.NewSize(1280, 720))
	mywindow.CenterOnScreen()
	mywindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	mywindow.ShowAndRun()

}
