package controllers

import (
	"github.com/teris-io/shortid"
	"log"
)

func IdGeneration() string {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		log.Fatal(err)
	}
	id, err := sid.Generate()
	if err != nil {
		log.Fatal(err)
	}

	return id
}
