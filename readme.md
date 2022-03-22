# Project
Alertzy SDK in golang<br />
https://alertzy.app/

# Send Alerts
```go
cTrans := alertzy.NewClient("https://alertzy.app", "[my account key]")


// simple
err := cTrans.Send("My simple title", "My message", alertzy.Normal, "test", "", "", nil)

// with image & link
err := cTrans.Send("My title with image & link", "My message", alertzy.Normal, "test", "https://go.dev/doc/gopher/doc.png", "https://github.com/ncostamagna/alertzy-sdk", nil)

// with buttons
err := cTrans.Send("My title with buttons", "My message", alertzy.Normal, "test", "", "", []alertzy.Buttons{
    {Text: "SDK Repo", Link: "https://github.com/ncostamagna/alertzy-sdk", Color: alertzy.Dark},
    {Text: "Google", Link: "http://google.com", Color: alertzy.Danger},
})
```

# Example
Examples in **example/send_alert/main.go**