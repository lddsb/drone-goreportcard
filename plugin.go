package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	// Plugin `plugin struct`
	Plugin struct {
		URL string
	}
)

var baseURL = "https://goreportcard.com/"

// Exec `exec plugin`
func (p *Plugin) Exec() error {
	urlLen := len(p.URL)
	if 4 > urlLen {
		return errors.New("invalid repository url")
	}

	suffix := p.URL[urlLen-4:]
	if ".git" == suffix {
		p.URL = p.URL[:urlLen-4]
	}
	resp, err := http.Get(baseURL + "checks?repo=" + p.URL)
	if nil != err {
		return err
	}

	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		bs, err := ioutil.ReadAll(resp.Body)
		if nil == err {
			return errors.New("API response error: " + string(bs))
		}

		return err
	}

	log.Println("Everything is OK")
	return err
}
