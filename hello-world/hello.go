package hello_world

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	frogHelloPrefix    = "ribbit, "

	french  = "French"
	spanish = "Spanish"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case "frog":
		prefix = frogHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
