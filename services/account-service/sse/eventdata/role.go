package eventdata

type Role struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permissionIds"`
}
