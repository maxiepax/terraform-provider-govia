package govia

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	username   string
	password   string
	url        string
	httpClient *http.Client
}

func newClient(u, p, url string) *Client {

	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	c := &http.Client{Timeout: 10 * time.Second, Transport: transport}

	return &Client{
		httpClient: c,
		username:   u,
		password:   p,
		url:        url,
	}
}

func (c *Client) get(path string, target interface{}) (r *http.Response, err error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/%s", c.url, path), nil)
	if err != nil {
		return
	}

	req.SetBasicAuth(c.username, c.password)

	r, err = c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer r.Body.Close()

	switch r.StatusCode {
	case 200, 404:
		break
	default:
		return r, fmt.Errorf("govia returned error: %s", r.Status)
	}

	err = json.NewDecoder(r.Body).Decode(target)
	return
}

func (c *Client) post(path string, item interface{}, ret interface{}) error {

	json_grp, _ := json.Marshal(item)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/v1/%s", c.url, path), bytes.NewReader(json_grp))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	/*

		log.Println("spew")
		log.Println(spew.Sdump(bodyString))
	*/

	if r.StatusCode != 200 {
		bodyBytes, _ := io.ReadAll(r.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("govia returned error: %s : %s", r.Status, bodyString)
	}

	err = json.NewDecoder(r.Body).Decode(ret)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) patch(path string, item interface{}) error {
	json_grp, _ := json.Marshal(item)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/v1/%s", c.url, path), bytes.NewReader(json_grp))
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	if r.StatusCode != 200 {
		return fmt.Errorf("govia returned error: %s", r.Status)
	}

	return nil
}

func (c *Client) delete(path string) error {

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/%s", c.url, path), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Content-Type", "application/json")

	r, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer r.Body.Close()

	switch r.StatusCode {
	case 200:
		return nil
	case 204:
		return nil
	default:
		return fmt.Errorf("govia returned error: %s", r.Status)
	}
}
