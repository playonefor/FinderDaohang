// models/urlinfo.go
package models

type UrlInfo struct {
	ID         uint `gorm:"primaryKey"`
	UrlName    string
	UrlPath    string
	UrlDesc    string
	UrlStatus  bool
	UrlGroupID uint
}

// TableName overrides the table name
func (UrlInfo) TableName() string {
	return "app_urlinfor"
}
