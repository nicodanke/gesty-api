package eventdata

type Action struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Enabled      bool   `json:"enabled"`
	CanBeDeleted bool   `json:"can_be_deleted"`
}
