package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetAppUrlinforTable(ctx *context.Context) table.Table {

	// appUrlinfor := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	appUrlinfor := table.NewDefaultTable(table.DefaultConfigWithDriver("sqlite"))

	info := appUrlinfor.GetInfo().HideFilterArea()

	info.SetFilterFormLayout(form.LayoutFlow)
	info.AddField("Id", "id", db.Int).
		FieldHide().
		FieldSortable()
	info.AddField("链接名称", "url_name", db.Varchar).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("链接值", "url_path", db.Varchar).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	info.AddField("链接描述", "url_desc", db.Text).
		FieldFilterable().
		FieldSortable().
		FieldEditAble()
	// info.AddField("是否启用", "url_status", db.Tinyint).
	// FieldFilterable().
	// FieldSortable().
	// FieldEditAble()

	info.AddField("是否启用", "url_status", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "是"
		}
		if model.Value == "1" {
			return "否"
		}
		return "未知"
	}).FieldFilterable().FieldSortable().FieldEditAble()
	// info.AddField("分组名称", "url_group_id", db.Int).
	// FieldFilterable().
	// FieldSortable().
	// FieldEditAble()

	// 连接 "profiles" 表，并显示 "profiles" 表中的 "name" 字段
	info.AddField("分组名称", "group_name", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		// fmt.Println(modstel.Value)
		// if model.Value == "" {
		// return ""
		// }
		// 假设这里已经设置了相关的连表查询逻辑
		// 这里是一个伪代码，表示如何获取关联表的数据
		// 实际实现时，您需要根据实际情况编写代码来查询关联表并返回显示的结果
		// 例如，您可能需要执行 SQL 查询或者使用 ORM 方法来获取关联数据
		// profileName := "Profile Name From Profiles Table"
		return model.Value
		// return profileName
	}).FieldJoin(types.Join{
		Table:     "app_urlgroup", // 连接的表名
		Field:     "url_group_id",
		JoinField: "id",
		BaseTable: "app_urlinfor",
	})

	info.AddField("创建时间", "timestamp", db.Datetime)

	info.SetTable("app_urlinfor").SetTitle("链接详情").SetDescription("链接详情")

	formList := appUrlinfor.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate().
		FieldHide()
	formList.AddField("链接名称", "url_name", db.Varchar, form.Text).FieldMust()
	formList.AddField("链接值", "url_path", db.Varchar, form.Text).FieldMust()
	formList.AddField("链接描述", "url_desc", db.Text, form.RichText)
	// formList.AddField("是否启用", "url_status", db.Tinyint, form.Number).FieldDefault("1")
	formList.AddField("是否启用", "url_status", db.Tinyint, form.Radio).
		// radio的选项，Text代表显示内容，Value代表对应值
		FieldOptions(types.FieldOptions{
			{Text: "是", Value: "0"},
			{Text: "否", Value: "1"},
		}).
		FieldDefault("0") // 设置默认的值

	formList.AddField("分组名称", "url_group_id", db.Int, form.SelectSingle).
		FieldOptionsFromTable("app_urlgroup", "group_name", "id").
		FieldDisplay(func(model types.FieldModel) interface{} {
			return model.Value
		}).
		FieldMust()

	// formList.AddField("分组名称", "url_group_id", db.Int, form.Number).FieldMust()
	formList.AddField("创建时间", "timestamp", db.Datetime, form.Datetime)

	formList.SetTable("app_urlinfor").SetTitle("链接详情").SetDescription("链接详情")

	return appUrlinfor
}
