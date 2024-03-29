package protocol

type Analysis struct {
	Event     Event     `json:"event"`
	Detection Detection `json:"detection"`
	Report    Report    `json:"report"`
}
