package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAppUrlgroupTable(ctx *context.Context) table.Table {

	// appUrlgroup := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	appUrlgroup := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := appUrlgroup.GetInfo().HideFilterArea()

	info.SetFilterFormLayout(form.LayoutFlow)
	info.AddField("Id", "id", db.Int)
	info.AddField("组名", "group_name", db.Varchar).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("创建时间", "timestamp", db.Datetime).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()

	info.SetTable("app_urlgroup").SetTitle("导航分组").SetDescription("导航分组")

	formList := appUrlgroup.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("组名", "group_name", db.Varchar, form.Text)
	formList.AddField("创建时间", "timestamp", db.Datetime, form.Datetime)

	formList.SetTable("app_urlgroup").SetTitle("导航分组").SetDescription("导航分组")

	return appUrlgroup
}
