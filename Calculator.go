package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func showCalc() {
	var historyarr []string
	output := ""
	historystr := ""
	ishistory := false
	input := widget.NewLabel(output)
	historybtn := widget.NewButton("History", func() {
		if ishistory {
			historystr = ""
		} else {
			for i := len(historyarr) - 1; i >= 0; i-- {
				historystr = historystr + historyarr[i]
				historystr += "\n"
				ishistory = true
			}
		}
		input.SetText(historystr)
	})
	Clearbtn := widget.NewButton("Clear", func() {
		output = ""
		input.SetText(output)
	})
	Openbtn := widget.NewButton("Open", func() {
		output += "("
		input.SetText(output)
	})
	Closebtn := widget.NewButton("Close", func() {
		output += ")"
		input.SetText(output)
	})
	backbtn := widget.NewButton("back", func() {
		output = output[:len(output)-1]
		input.SetText(output)
	})
	Onebtn := widget.NewButton("1", func() {
		output += "1"
		input.SetText(output)
	})
	Twobtn := widget.NewButton("2", func() {
		output += "2"
		input.SetText(output)
	})
	Threebtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	Fourbtn := widget.NewButton("4", func() {
		output += "4"
		input.SetText(output)
	})
	Fivebtn := widget.NewButton("5", func() {
		output += "5"
		input.SetText(output)
	})
	Sixbtn := widget.NewButton("6", func() {
		output += "6"
		input.SetText(output)
	})
	Sevenbtn := widget.NewButton("7", func() {
		output += "7"
		input.SetText(output)
	})
	Eightbtn := widget.NewButton("8", func() {
		output += "8"
		input.SetText(output)
	})
	Ninebtn := widget.NewButton("9", func() {
		output += "9"
		input.SetText(output)
	})
	Zerobtn := widget.NewButton("0", func() {
		output += "0"
		input.SetText(output)
	})
	Dividebtn := widget.NewButton("/", func() {
		output += "/"
		input.SetText(output)
	})
	Multbtn := widget.NewButton("*", func() {
		output += "*"
		input.SetText(output)
	})
	Plusbtn := widget.NewButton("+", func() {
		output += "+"
		input.SetText(output)
	})
	Minusbtn := widget.NewButton("-", func() {
		output += "-"
		input.SetText(output)
	})
	Dotbtn := widget.NewButton(".", func() {
		output += "."
		input.SetText(output)
	})
	Equalbtn := widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strtoappend := output + " = " + ans
				historyarr = append(historyarr, strtoappend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		input.SetText(output)
	})

	Equalbtn.Importance = widget.HighImportance

	calcContainer := container.NewVBox(
		container.NewVBox(
			input,
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(2,
					historybtn,
					backbtn,
				),
				container.NewGridWithColumns(4,
					Clearbtn,
					Closebtn,
					Openbtn,
					Dividebtn,
				),
				container.NewGridWithColumns(4,
					Ninebtn,
					Eightbtn,
					Sevenbtn,
					Multbtn,
				),
				container.NewGridWithColumns(4,
					Sixbtn,
					Fivebtn,
					Fourbtn,
					Minusbtn,
				),
				container.NewGridWithColumns(4,
					Threebtn,
					Twobtn,
					Onebtn,
					Plusbtn,
				),
				container.NewGridWithColumns(2,
					container.NewGridWithColumns(2,
						Zerobtn,
						Dotbtn,
					),
					Equalbtn,
				))),
	)
	w := myapp.NewWindow("Calc")
	w.Resize(fyne.NewSize(500, 280))

	w.SetContent(
		container.NewBorder(Deskbtn, nil, nil, nil, calcContainer),
	)

	w.Show()
}
