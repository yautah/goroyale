package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	// "io/ioutil"
	// "net/http"
	"net/url"
	"time"
)

const baseURL = "https://api.clashroyale.com/v1"
const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiIsImtpZCI6IjI4YTMxOGY3LTAwMDAtYTFlYi03ZmExLTJjNzQzM2M2Y2NhNSJ9.eyJpc3MiOiJzdXBlcmNlbGwiLCJhdWQiOiJzdXBlcmNlbGw6Z2FtZWFwaSIsImp0aSI6IjQxYWRlMzdiLTYxNTgtNDcyNS04ZjdmLWZkMjI1MjQzNzFhNyIsImlhdCI6MTUzNDkyMjA0Nywic3ViIjoiZGV2ZWxvcGVyLzYwNjE0NTcwLTMwNDgtNWZkNS04NzdlLTY3ZDljNGVmNjA2NyIsInNjb3BlcyI6WyJyb3lhbGUiXSwibGltaXRzIjpbeyJ0aWVyIjoiZGV2ZWxvcGVyL3NpbHZlciIsInR5cGUiOiJ0aHJvdHRsaW5nIn0seyJjaWRycyI6WyIxMTguMTQ0LjI0NC4xNDEiLCI0Ny4yNTQuNDQuMjM4Il0sInR5cGUiOiJjbGllbnQifV19.FbW5PdsX-ori6K4ZCVM9Wma8F5q2VzjBJllHbJOWCLqAEQxyXAanFmu_jCCGLWZTQTmZks7PB5a1a8XRnGcIaw"

// Client allows you to easily interact with RoyaleAPI.
type Client struct {
	Token string

	client *resty.Client
	// using empty struct because it has a byte size of 0
	// i don't care what's in the channel, just that something is
}

// New creates a new RoyaleAPI client.
func New(token string, timeout time.Duration) (c *Client, err error) {
	c = &Client{
		client: resty.New(),
	}
	if token == "" {
		err = errors.New("client requires token for authorization with the API")
		return
	}
	c.Token = token

	if timeout != 0 {
		c.client.SetTimeout(timeout * time.Second)
	}
	c.client.SetTimeout(30 * time.Second)
	// c.client.SetDebug(true)

	return
}

func (c *Client) get(path string, params url.Values) (bytes []byte, err error) {

	path = baseURL + path

	c.client.SetHeader("authorization", fmt.Sprintf("Bearer %s", c.Token))
	c.client.R().SetQueryString(params.Encode())

	// req.URL.RawQuery = params.Encode()

	resp, err := c.client.R().Get(path)

	if err != nil {
		fmt.Println("err2", err)
		return
	}

	bytes = resp.Body()

	if resp.StatusCode() != 200 {
		var apiErr APIError
		json.Unmarshal(bytes, &apiErr)
		apiErr.StatusCode = resp.StatusCode()
		return []byte{}, apiErr
	}

	return
}

// func (c *Client) get1(path string, params url.Values) (bytes []byte, err error) {

// path = baseURL + path
// req, err := http.NewRequest("GET", path, nil)
// if err != nil {
// fmt.Println("err1", err)
// return
// }
// req.Header.Add("authorization", fmt.Sprintf("Bearer %s", c.Token))
// req.URL.RawQuery = params.Encode()

// resp, err := c.client.Do(req)
// if err != nil {
// fmt.Println("err2", err)
// return
// }
// defer resp.Body.Close()

// bytes, err = ioutil.ReadAll(resp.Body)

// if resp.StatusCode != 200 {
// var apiErr APIError
// json.Unmarshal(bytes, &apiErr)
// apiErr.StatusCode = resp.StatusCode
// return []byte{}, apiErr
// }

// return
// }
