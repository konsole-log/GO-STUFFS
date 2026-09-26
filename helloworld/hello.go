// Package helloworld provides function for saying hello world in three different languages(spanish, nepali, english)
package helloworld

const (
	nepali             = "Nepali"
	spanish            = "Spanish"
	englishHelloPrefix = "Hello "
	nepaliHelloPrefix  = "Namaste "
	spanishHelloPrefix = "Hola "
)

// Hello returns a greeting
func Hello(name, greeting string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(greeting) + name
}

func greetingPrefix(greeting string) (prefix string) {
	switch greeting {
	case spanish:
		prefix = spanishHelloPrefix
	case nepali:
		prefix = nepaliHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
