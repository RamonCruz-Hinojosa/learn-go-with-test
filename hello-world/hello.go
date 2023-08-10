package hello_world

const englishHelloPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "world"
	}

	return englishHelloPrefix + name
}
