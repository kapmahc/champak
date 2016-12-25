package mux

var (
	router Router
)

// Use use
func Use(r Router) {
	router = r
}
