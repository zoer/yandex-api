package direct

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server

	token string
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	token = "foo"
	client = NewClient(token)
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestClient(t *testing.T) {
	a := assert.New(t)

	token := "foo"
	c := NewClient(token)

	a.Equal(token, c.Token)
	a.Equal(http.DefaultClient, c.client)
	a.Equal(userAgent, c.UserAgent)

	BaseURL, err := url.Parse(defaultBaseURL)
	a.NoError(err)
	a.Equal(BaseURL, c.BaseURL)
}

func TestNewRequest(t *testing.T) {
	a := assert.New(t)
	c := NewClient("foo")

	cID := 123
	inBody, outBody := &Campaign{CampaignID: cID}, `{"CampaignID":123}`+"\n"
	req, _ := c.NewRequest(inBody)

	a.Equal(req.URL.String(), defaultBaseURL)

	body, _ := ioutil.ReadAll(req.Body)
	a.Equal(string(body), outBody)

	a.Equal(req.Header.Get("Content-Type"), "application/json")

	a.Equal(req.Header.Get("User-Agent"), c.UserAgent)
}

func TestDo(t *testing.T) {
	a := assert.New(t)

	setup()
	defer teardown()

	type foo struct {
		Value string
	}

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		a.Equal(r.Method, "POST")
		fmt.Fprintf(w, `{"Value":"test"}`)
	})

	req, _ := client.NewRequest(nil)
	body := new(foo)

	_, err := client.Do(req, body)
	a.NoError(err)

	expected := &foo{"test"}
	a.Equal(body, expected)
}

func TestDo_httpError(t *testing.T) {
	a := assert.New(t)

	setup()
	defer teardown()

	mux.HandleFunc(client.BaseURL.RequestURI(), func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Bad Gateway", 502)
	})

	req, _ := client.NewRequest(nil)
	_, err := client.Do(req, nil)
	a.Error(err)
}
