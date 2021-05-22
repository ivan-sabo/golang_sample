package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Homepage struct {
	l *log.Logger
}

func NewHomepage(l *log.Logger) *Homepage {
	return &Homepage{l}
}

func (h *Homepage) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Test")
}
