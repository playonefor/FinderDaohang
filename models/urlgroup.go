// models/urlgroup.go
package models

type UrlGroup struct {
	ID        uint `gorm:"primaryKey"`
	GroupName string
	UrlInfos  []UrlInfo `gorm:"foreignKey:UrlGroupID"`
}

// TableName overrides the table name
func (UrlGroup) TableName() string {
	return "app_urlgroup"
}
