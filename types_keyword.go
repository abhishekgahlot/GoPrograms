package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type webPage struct {
	url  string
	body []byte
	err  error
}

func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
		return
	}
	defer resp.Body.Close()
	w.body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.err = err
	}
}

func (w *webPage) isOk() bool {
	return w.err == nil
}

func main() {
	w := &webPage{url: "http://www.google.com/"}
	// or new (webPage) or &webPage
	// or you could do w.url = something
	w.get()
	if w.isOk() {
		fmt.Printf("%s, %s, %d", w.url, w.err, len(w.body))
	} else {
		fmt.Printf("Messed up!")
	}
}
