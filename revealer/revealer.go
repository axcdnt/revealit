package revealer

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

