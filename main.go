package main

const (
	englishHelloPrefix   = "Hello, "
	spanishHelloPrefix   = "Hola, "
	frenchHelloPrefix    = "Bonjour, "
	ukrainianHelloPrefix = "Вітаю, "
	spanish              = "Spanish"
	french               = "French"
	ukrainian            = "Ukrainian"
)

// 'Hello' takes name and language parameters with type string and returns greetings
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case ukrainian:
		prefix = ukrainianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
