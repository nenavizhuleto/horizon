package analysis

type Severity int

const (
	_ Severity = iota
	INFO
	WARN
	PANIC
)

type Report struct {
	UserData any
	Severity Severity
}
