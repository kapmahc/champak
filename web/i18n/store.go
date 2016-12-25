package i18n

// Store store
type Store interface {
	Set(lang, code, message string, override bool) error
	All(lang string) (map[string]string, error)
	Get(lang, code string) (string, error)
	Del(lang, code string) error
}
