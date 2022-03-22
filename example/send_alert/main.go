package main

import (
	"fmt"

	"github.com/ncostamagna/alertzy-sdk/alertzy"
)

func main() {

	cTrans := alertzy.NewClient("https://alertzy.app", "zanov1i1fulmr56")
	var err error

	// simple
	err = cTrans.Send("My simple title", "My message", alertzy.Normal, "test", "", "", nil)
	fmt.Print(err)

	// with image & link
	err = cTrans.Send("My title with image & link", "My message", alertzy.Normal, "test", "https://go.dev/doc/gopher/doc.png", "http://google.com", nil)
	fmt.Print(err)

	// with buttons
	err = cTrans.Send("My title with buttons", "My message", alertzy.Normal, "test", "", "", []alertzy.Buttons{
		{Text: "SDK Repo", Link: "http://google.com", Color: alertzy.Dark},
		{Text: "Google", Link: "http://google.com", Color: alertzy.Danger},
	})
	fmt.Print(err)
}
