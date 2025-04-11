package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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

func Gui(configJsonUriStr string, settingsJsonUriStr string, selectTab string) {
	// Create a new Fyne app
	myApp := app.New()
	myWindow := myApp.NewWindow("CliX (Cli eXtender)")

	// Create content for each tab
	tab1Content := TabConfigContent(configJsonUriStr, myWindow)
	tab2Content := TabJsonContent(configJsonUriStr, myWindow)
	tab3Content := TabJsonContent(settingsJsonUriStr, myWindow)

	// Create the tabs and assign their labels and content
	configTab := container.NewTabItem("Config", tab1Content)
	configJsonTab := container.NewTabItem("config.json", tab2Content)
	settingsJsonTab := container.NewTabItem("settings.json", tab3Content)

	// Create an AppTabs container that holds all the tabs
	tabContainer := container.NewAppTabs(configTab, configJsonTab, settingsJsonTab)

	// Set the initial tab
	switch selectTab {
	case StaticConfig:
		tabContainer.Select(configTab)
	case StaticConfigJson:
		tabContainer.Select(configJsonTab)
	case StaticSettingsJson:
		tabContainer.Select(settingsJsonTab)
	}

	// When user updates the config file to keep both tabs content consistent,
	// update the contents of the tab when selected
	tabContainer.OnSelected = func(item *container.TabItem) {
		configTab.Content = TabConfigContent(configJsonUriStr, myWindow)
		configJsonTab.Content = TabJsonContent(configJsonUriStr, myWindow)
		settingsJsonTab.Content = TabJsonContent(settingsJsonUriStr, myWindow)
	}

	// Set the content of the window to be the tab container
	myWindow.SetContent(tabContainer)

	// Resize the window and show the app
	myWindow.Resize(fyne.NewSize(600, 500))
	myWindow.ShowAndRun()
}

func TabConfigContent(uriStr string, myWindow fyne.Window) *fyne.Container {
	aliases := GetListConfigAlias()
	var tabs []*container.TabItem
	for _, alias := range aliases {

		command_struct := GetConfigAliasValue(alias)

		// Marshal struct into a JSON object
		data, err := json.Marshal(command_struct) // `command` is your Command struct
		if err != nil {
			panic(err)
		}

		// Format JSON
		var out bytes.Buffer
		err = json.Indent(&out, []byte(data), "", "  ")
		if err != nil {
			panic(err)
		}
		// Save the formatted JSON to a variable
		formattedJSON := out.String()

		textBox := widget.NewMultiLineEntry()
		textBox.SetText(formattedJSON)
		tabs = append(tabs, container.NewTabItem(alias, textBox))
	}

	innerTabContainer := container.NewAppTabs(tabs...)
	innerTabContainer.SetTabLocation(container.TabLocationLeading)

	// Save file function called by saveBtn
	saveFile := func() {

		cfg := Config{
			Commands: make(map[string]Command),
		}

		// Access the tabs manually
		// Grab the tab's text and content to recreate the config file to save
		for i, tab := range tabs {
			commandKey := tab.Text
			textBox := tab.Content.(*widget.Entry)
			commandJSON := textBox.Text
			fmt.Printf("Tab %d: %s\n", i+1, tab.Text)
			fmt.Printf("Tab %d Content: %s\n", i+1, textBox.Text)

			var command Command
			err := json.Unmarshal([]byte(commandJSON), &command)
			if err != nil {
				log.Printf("Error unmarshaling JSON for command %s: %v", commandKey, err)
				dialog.ShowError(err, myWindow)
				return
			}

			cfg.Commands[commandKey] = command
		}

		jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			log.Fatal(err)
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

		_, err = writer.Write(jsonBytes)
		if err != nil {
			log.Println("Failed to write to file:", err)
			return
		}

		// Show success message
		dialog.ShowInformation("Success", "File saved successfully", myWindow)
		log.Println("File saved successfully.")
	}

	saveBtn := widget.NewButton("Save", saveFile)

	innerTabContainer.Refresh()

	return container.NewStack(container.NewBorder(nil, saveBtn, nil, nil, innerTabContainer)) // Container that takes up all available space
}

func TabJsonContent(uriStr string, myWindow fyne.Window) *fyne.Container {
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

	return container.NewStack(container.NewBorder(nil, saveBtn, nil, nil, textArea)) // Container that takes up all available space
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

func recreateConfig(alias string, commandJSON string) (Config, error) {
	var command Command
	err := json.Unmarshal([]byte(commandJSON), &command)
	if err != nil {
		return Config{}, err
	}

	commandsMap := map[string]Command{
		alias: command,
	}

	config := Config{
		Commands: commandsMap,
	}

	return config, nil
}

func OpenConfigJsonInTabs(uriStr string) {

	myApp := app.New()
	myWindow := myApp.NewWindow("Clix Config File")

	aliases := GetListConfigAlias()

	if len(aliases) == 0 {

		config_file_path := strings.Replace(uriStr, "file://", "", -1)

		multilineString := `No aliases found. 
Please check your config file located at ` + config_file_path

		display := widget.NewLabel(multilineString)
		myWindow.SetContent(container.NewBorder(nil, nil, nil, nil, display)) // top, bottom, left, right, center

		myWindow.Resize(fyne.NewSize(500, 400))
		myWindow.ShowAndRun()
	}

	var tabs []*container.TabItem
	for _, alias := range aliases {

		command_struct := GetConfigAliasValue(alias)

		// Marshal struct into a JSON object
		data, err := json.Marshal(command_struct) // `command` is your Command struct
		if err != nil {
			panic(err)
		}

		// Format JSON
		var out bytes.Buffer
		err = json.Indent(&out, []byte(data), "", "  ")
		if err != nil {
			panic(err)
		}
		// Save the formatted JSON to a variable
		formattedJSON := out.String()

		textBox := widget.NewMultiLineEntry()
		textBox.SetText(formattedJSON)
		tabs = append(tabs, container.NewTabItem(alias, textBox))
	}

	tabContainer := container.NewAppTabs(tabs...)
	tabContainer.SetTabLocation(container.TabLocationLeading)

	// Save file function called by saveBtn
	saveFile := func() {

		cfg := Config{
			Commands: make(map[string]Command),
		}

		// Access the tabs manually
		// Grab the tab's text and content to recreate the config file to save
		for i, tab := range tabs {
			commandKey := tab.Text
			textBox := tab.Content.(*widget.Entry)
			commandJSON := textBox.Text
			fmt.Printf("Tab %d: %s\n", i+1, tab.Text)
			fmt.Printf("Tab %d Content: %s\n", i+1, textBox.Text)

			var command Command
			err := json.Unmarshal([]byte(commandJSON), &command)
			if err != nil {
				log.Printf("Error unmarshaling JSON for command %s: %v", commandKey, err)
				dialog.ShowError(err, myWindow)
				return
			}

			cfg.Commands[commandKey] = command
		}

		jsonBytes, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			log.Fatal(err)
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

		_, err = writer.Write(jsonBytes)
		if err != nil {
			log.Println("Failed to write to file:", err)
			return
		}

		// Show success message
		dialog.ShowInformation("Success", "File saved successfully", myWindow)
		log.Println("File saved successfully.")
	}

	saveBtn := widget.NewButton("Save", saveFile)

	myWindow.SetContent(container.NewBorder(nil, saveBtn, nil, nil, tabContainer)) // top, bottom, left, right, center

	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.ShowAndRun()
}
