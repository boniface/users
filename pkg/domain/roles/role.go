package roles

type Role struct {
	SiteId      string `json:"siteId"`
	Id          string `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}

type RolesPool struct {
	Id          string `json:"id"`
	RoleName    string `json:"roleName"`
	Description string `json:"description"`
}
