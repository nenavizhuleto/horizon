package protocol

type Severity int

const (
	_ Severity = iota
	INFO
	WARN
	PANIC
)

type Report struct {
	UserData any      `json:"user_data"`
	Severity Severity `json:"severity"`
}
