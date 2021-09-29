package main

import (
	b64 "encoding/base64"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Base64 Converter")

	base64 := binding.NewString()
	txt := binding.NewString()

	inputBase64 := widget.NewEntryWithData(base64)
	outputText := widget.NewEntryWithData(txt)

	inputBase64.OnChanged = func(text string) {
		textBase4, _ := b64.URLEncoding.DecodeString(text)
		_ = txt.Set(string(textBase4))
	}

	outputText.OnChanged = func(text string) {
		_ = base64.Set(b64.StdEncoding.EncodeToString([]byte(text)))
	}

	w.SetContent(container.NewGridWithColumns(4,
		widget.NewLabel("Base64: "),
		inputBase64,
		widget.NewLabel("Text: "),
		outputText,
	))

	w.Resize(fyne.NewSize(800, 500))
	w.ShowAndRun()
}
