package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (hello *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
		hello.l.Println("Running Hello Handler")

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(rw, "Unable to read request body", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", b)
}
