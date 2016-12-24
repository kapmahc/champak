package web

// KEY http request context key type
type KEY string

// H hash object
type H map[string]interface{}

const (
	// LOCALE locale key
	LOCALE = KEY("locale")
	// DATA data key
	DATA = KEY("data")
	// TO key of to
	TO = "to"
	// APPLICATION application layout
	APPLICATION = "application"
	// DASHBOARD dashboard layout
	DASHBOARD = "dashboard"
)
