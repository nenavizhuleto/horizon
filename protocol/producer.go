package protocol

type Producer struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	GroupID   string       `json:"group_id"`
	GroupName string       `json:"group_name"`
	Modules   []ModuleName `json:"modules"`
}
