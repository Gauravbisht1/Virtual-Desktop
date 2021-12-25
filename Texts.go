package main

import (
	"io/ioutil"
	"strconv"

	//"fyne.io/fyne/dialog"
	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func showTextEditor(w fyne.Window) {
	count := 1
	// a := app.New()
	// w := a.NewWindow("Text Editor")
	// x := fyne.Size{800, 500}
	// w.Resize(x)

	content := container.NewVBox(
		container.NewHBox(
			widget.NewLabel("Text Editor"),
		),
	)
	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))
	// w := myapp.NewWindow("TextEditor")
	// w.Resize(fyne.NewSize(500, 280))

	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text...")

	input.Resize(fyne.NewSize(400, 400))

	savebtn := widget.NewButton("Save File", func() {
		savedialog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textdata := []byte(input.Text)
				uc.Write(textdata)
			}, w)
		savedialog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		savedialog.Show()
	})

	openbtn := widget.NewButton("Open File", func() {
		opendialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				readdata, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", readdata)
				viewdata := widget.NewMultiLineEntry()
				viewdata.SetText(string(output.StaticContent))
				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))
				w.SetContent(container.NewScroll(viewdata))

				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
		opendialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		opendialog.Show()
	})

	textcontainer := container.NewVBox(
		content,
		input,
		container.NewHBox(
			savebtn,
			openbtn,
		),
	)
	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, textcontainer))

	// w.SetContent(
	// 	container.NewVBox(
	// 		content,
	// 		input,
	// 		container.NewHBox(
	// 			savebtn,
	// 			openbtn,
	// 		),
	// 	),
	// )
	w.Show()
}
