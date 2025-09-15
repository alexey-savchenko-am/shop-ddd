package common

type QueryDB interface {
	Select(dest interface{}, query string, args ...any) error
	Get(dest interface{}, query string, args ...any) error
}
