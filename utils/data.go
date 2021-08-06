package utils

type Dependencies struct {
	CategoryGroups []struct {
		Categories []struct {
			Name        string   `json:"name"`
			Description string   `json:"description"`
			Projects    []string `json:"projects"`
			Permalink   string   `json:"permalink"`
		} `json:"categories"`
		Description interface{} `json:"description"`
		Name        string      `json:"name"`
		Permalink   string      `json:"permalink"`
	} `json:"category_groups"`
}

type Revealer interface {
	Parse() map[string][]string
	PrettyPrint(categories map[string][]string) // Receives a map[category][dependencies]
}

func Contains(element string, elements []string) bool {
	for _, e := range elements {
		if e == element {
			return true
		}
	}

	return false
}