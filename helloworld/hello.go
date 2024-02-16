package helloworld

import "fmt"

const (
	spanishLanguage    = "Spanish"
	frenchLanguage     = "French"
	italianLanguage    = "Italian"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	italianHelloPrefix = "Ciao, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	case italianLanguage:
		prefix = italianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}
