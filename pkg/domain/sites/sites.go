package sites

type Site struct {
	SiteId         string `json:"siteId"`
	SiteName       string `json:"siteName"`
	SiteUrl        string `json:"siteUrl"`
	SiteAdminEmail string `json:"siteAdminEmail"`
}