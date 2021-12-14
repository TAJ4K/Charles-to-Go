package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/atotto/clipboard"
)

func main(){
	ui.Main(uiInit)
}

func uiInit(){
	mainwin := ui.NewWindow("Charles to Go header converter", 350, 20, true)
	mainwin.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	vbContainer := ui.NewVerticalBox()
	vbContainer.SetPadded(false)

	button := ui.NewButton("Convert!")

	button.OnClicked(func(*ui.Button) {
		process()
	})

	vbContainer.Append(button, false)

	mainwin.SetChild(vbContainer)

	mainwin.Show()
}

func process(){
	convertee, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	beginning := regexp.MustCompile(`(?m)^`)
	middle := regexp.MustCompile(`(?m)	`)
	end := regexp.MustCompile(`(?m)\r`)

	s1 := beginning.ReplaceAllString(convertee, `req.Header.Set("`)
	s2 := middle.ReplaceAllString(s1, `", "`)
	s3 := end.ReplaceAllString(s2, `")`)
	s4 := s3 + `")`
	s5 := strings.ReplaceAll(s4, `, ""`, ", `\"")
	s6 := strings.ReplaceAll(s5, `"")`, "\"`)")

	err = clipboard.WriteAll(s6)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}