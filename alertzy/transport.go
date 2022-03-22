package alertzy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	c "github.com/ncostamagna/streetflow/client"
)

type DataResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"`
}

type Priority int8

const (
	Normal Priority = iota
	High
	GRPC
)

type Color string

const (
	Primary Color = "primary"
	Success Color = "success"
	Warning Color = "warning"
	Danger  Color = "danger"
	Info    Color = "info"
	Light   Color = "light"
	Dark    Color = "dark"
)

type Buttons struct {
	Text  string `json:"text"`
	Link  string `json:"link"`
	Color Color  `json:"color"`
}

// Transport object
type Transport interface {
	Send(title, message string, priority Priority, group, image, link string, buttons []Buttons) error
}

type clientHTTP struct {
	client     c.RequestBuilder
	accountKey string
}

// NewClient - create a new transport client, for example HTTP, Socket, gRPC, etc..
func NewClient(baseURL, accountKey string) Transport {

	header := http.Header{}

	return &clientHTTP{
		client: c.RequestBuilder{
			Headers:        header,
			BaseURL:        baseURL,
			ConnectTimeout: 5000 * time.Millisecond,
			LogTime:        true,
		},
		accountKey: accountKey,
	}

}

func (c *clientHTTP) Send(title, message string, priority Priority, group, image, link string, buttons []Buttons) error {

	if message == "" {
		return errors.New("message is required")
	}

	if title == "" {
		return errors.New("title is required")
	}

	u := url.URL{}
	u.Path += "/send"
	q := u.Query()
	q.Set("accountKey", c.accountKey)
	q.Set("title", title)
	q.Set("message", message)

	q.Set("priority", fmt.Sprint(priority))

	if group != "" {
		q.Set("group", group)
	}

	if image != "" {
		q.Set("image", image)
	}

	if link != "" {
		q.Set("link", link)
	}

	if len(buttons) != 0 {
		v, _ := json.Marshal(buttons)
		q.Set("buttons", string(v))
	}

	u.RawQuery = q.Encode()

	reps := c.client.Post(u.String(), nil)

	if reps.Err != nil {
		return reps.Err
	}

	if reps.StatusCode > 299 {
		return fmt.Errorf("code: %d, message: %s", reps.StatusCode, reps)
	}

	var response DataResponse
	if err := json.Unmarshal(reps.Bytes(), &response); err != nil {
		return err
	}

	if response.Error != "" {
		return errors.New(response.Error)
	}

	return nil
}
