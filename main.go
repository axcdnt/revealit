package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/axcdnt/revealit/ruby"
	"github.com/axcdnt/revealit/utils"
)

func main() {
	wd, _ := os.Getwd()
	language := flag.String("language", "", "the project's main language")
	flag.Parse()

	runner := newRunner(*language, wd)
	result := runner.Parse()
	runner.PrettyPrint(result)
}

func newRunner(language, wd string) utils.Revealer {
	if language == "ruby" {
		return &ruby.RubyRunner{Path: fmt.Sprintf("%s/%s", wd, "Gemfile")}
	} else {
		log.Fatal("language parser not implemented")
	}

	return nil
}