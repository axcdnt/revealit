package revealer

// Revealer is the main interface if you want to implement
// this idea for different programming languages.
//
// The Parse() method must satisfy a common struct that contains
// a 'categories' property that is a `map[string][]string{}`.
type Revealer interface {
	Parse()
	PrettyPrint()
}

func Contains(element string, elements []string) bool {
	for _, e := range elements {
		if element == e {
			return true
		}
	}

	return false
}

