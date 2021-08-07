package main

import (
	"flag"
	"log"
	"os"

	"github.com/axcdnt/revealit/revealer"
	"github.com/axcdnt/revealit/ruby"
)

func main() {
	path, _ := os.Getwd()
	language := flag.String("language", "", "the project's main language")
	flag.Parse()

	runner := newRunner(*language, path)
	runner.Parse()
	runner.PrettyPrint()
}

func newRunner(language, path string) revealer.Revealer {
	if language == "ruby" {
		return ruby.New(path)
	} else {
		log.Fatal("language parser not implemented")
	}

	return nil
}