package main

const defaultName = "world"

const (
	en = "English"
	es = "Spanish"
	fr = "French"
)

var prefixRegister = map[string]string{
	en: "Hello, ",
	es: "Hola, ",
	fr: "Bonjour, ",
}

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultName
	}

	prefix, ok := prefixRegister[lang]
	if !ok {
		prefix = prefixRegister[en]
	}

	return prefix + name
}
