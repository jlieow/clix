package util

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func SampleGUI() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}

func OpenFileWithPicker() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Fyne Text Editor")

	textArea := widget.NewMultiLineEntry()
	textArea.SetPlaceHolder("Open a file to display contents...")

	openFile := func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}
			if reader == nil {
				return
			}

			data, err := ioutil.ReadAll(reader)
			if err != nil {
				dialog.ShowError(err, myWindow)
				return
			}

			textArea.SetText(string(data))
		}, myWindow)
	}

	openBtn := widget.NewButton("Open File", openFile)

	myWindow.SetContent(container.NewBorder(openBtn, nil, nil, nil, textArea))
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()
}

func OpenConfigJsonInGui(uriStr string) error {
	myApp := app.New()
	myWindow := myApp.NewWindow("CliX Config File")

	textArea := widget.NewMultiLineEntry()

	uri, err := storage.ParseURI(uriStr)
	if err != nil {
		log.Fatal(err)
	}

	reader, err := storage.OpenFileFromURI(uri)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

	textArea.SetText(string(data))

	// Save file function called by saveBtn
	saveFile := func() {
		// Validate JSON before saving
		var js interface{}
		if err := json.Unmarshal([]byte(textArea.Text), &js); err != nil {
			dialog.ShowError(err, myWindow)
			return
		}

		uri, err := storage.ParseURI(uriStr)
		if err != nil {
			log.Println("Failed to parse URI:", err)
			return
		}

		writer, err := storage.Writer(uri)
		if err != nil {
			log.Println("Failed to open file for writing:", err)
			return
		}
		defer writer.Close()

		_, err = writer.Write([]byte(textArea.Text))
		if err != nil {
			log.Println("Failed to write to file:", err)
			return
		}

		// Show success message
		dialog.ShowInformation("Success", "File saved successfully", myWindow)
		log.Println("File saved successfully.")
	}

	saveBtn := widget.NewButton("Save", saveFile)

	myWindow.SetContent(container.NewBorder(nil, saveBtn, nil, nil, textArea)) // top, bottom, left, right, center
	myWindow.Resize(fyne.NewSize(600, 400))
	myWindow.ShowAndRun()

	return nil
}
