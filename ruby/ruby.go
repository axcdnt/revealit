package ruby

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/axcdnt/revealit/utils"
)

//go:embed ruby.json
var f embed.FS

type RubyRunner struct {
	Path string
}

func (r *RubyRunner) Parse() map[string][]string {
	f, err := os.Open(r.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dependencies := parseGemfile()
	result := map[string][]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "gem ") {
			// Refactor here
			splitted := strings.Split(line, ", ")
			gem := splitted[0]
			gem = strings.ReplaceAll(gem, "\"", "")
			gem = strings.ReplaceAll(gem, "'", "")
			gem = strings.TrimPrefix(gem, "gem ")

			for _, group := range  dependencies.CategoryGroups {
				for _, category := range group.Categories {
					if utils.Contains(gem, category.Projects) {
						result[category.Name] = append(result[category.Name], gem)
						continue
					} else {
						result["Uncategorized"] = append(result[category.Name], gem)
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func (r *RubyRunner) PrettyPrint(categories map[string][]string) {
	for category, gems := range categories {
		fmt.Printf("%s(%d) \n\t%v\n", category, len(gems), gems)
	}
}

func parseGemfile() utils.Dependencies {
	data, err := f.ReadFile("ruby.json")
	var d utils.Dependencies
	json.Unmarshal(data, &d)

	if err != nil {
		log.Fatal("an error occurred while reading the dependencies file", err)
	}

	return d
}
