package ruby

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/axcdnt/revealit/revealer"
)


var (
	//go:embed ruby.json
	f embed.FS
	sanitizerRegex = regexp.MustCompile(`,|'|"|\s`)
)

type Categories struct {
	CategoryGroups []struct {
		Categories []struct {
			Name        string   `json:"name"`
			Projects    []string `json:"projects"`
		} `json:"categories"`
		Name        string      `json:"name"`
	} `json:"category_groups"`
}

// RubyRunner is the main struct used by PrettyPrint()
type RubyRunner struct {
	Path string
	dependencyFile string
	categories map[string][]string
	total int
}

func New(path string) *RubyRunner {
	return &RubyRunner{
		Path:       fmt.Sprintf("%s/%s", path, "Gemfile"),
		dependencyFile: "Gemfile",
		categories: map[string][]string{},
	}
}

func (r *RubyRunner) Parse() {
	f, err := os.Open(r.Path)
	if err != nil {
		log.Fatal("dependency file could not be found: ", r.dependencyFile)
	}
	defer f.Close()

	allCategories := parseCategories()
	gemCategories := map[string][]string{}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		if strings.HasPrefix(line, "gem") {
			gem := gem(line)
			found := false
			r.total++
			for _, group := range allCategories.CategoryGroups {
				for _, category := range group.Categories {
					if isCategorized(gem, category.Projects) {
						found = true
						gemCategories[category.Name] = append(gemCategories[category.Name], gem)
					}
				}
			}

			if !found {
				gemCategories["Uncategorized"] = append(gemCategories["Uncategorized"], gem)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	r.categories = gemCategories
}

func isCategorized(gem string, projects []string) bool {
	return revealer.Contains(gem, projects)
}

func (r *RubyRunner) PrettyPrint() {
	fmt.Printf("Total: %d\n\n", r.total)
	for category, gems := range r.categories {
		fmt.Printf("%s(%d) \n\t%v\n", category, len(gems), gems)
	}
}

// Private stuff

func gem(line string) string {
	splitted := strings.Split(strings.Trim(line, " "), " ")[1]
	return sanitizerRegex.ReplaceAllString(splitted, "$1")
}

func parseCategories() Categories {
	data, err := f.ReadFile("ruby.json")
	var c Categories
	json.Unmarshal(data, &c)

	if err != nil {
		log.Fatal("an error occurred while reading the dependencies file", err)
	}

	return c
}
