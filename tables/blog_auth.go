package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBlogAuthTable(ctx *context.Context) table.Table {

	blogAuth := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))
	blogAuth = table.NewDefaultTable(table.Config{
		Driver:     db.DriverMysql,
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: "blogdb",
		PrimaryKey: table.PrimaryKey{
			Type: db.Int,
			Name: table.DefaultPrimaryKeyName,
		},
	})

	info := blogAuth.GetInfo().HideFilterArea()
	//info := blogAuth.GetInfo().SetFilterFormLayout(form.LayoutThreeCol).ExportValue()

	info.AddField("Id", "id", db.Int).FieldFilterable().FieldSortable()
	info.AddField("用户名", "username", db.Varchar).FieldCopyable().FieldFilterable()
	info.AddField("Password", "password", db.Varchar)

	info.SetTable("blog_auth").SetTitle("博客用户").SetDescription("博客用户")

	formList := blogAuth.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("Username", "username", db.Varchar, form.Text)
	formList.AddField("密码", "password", db.Varchar, form.Password)

	formList.SetTable("blog_auth").SetTitle("博客用户").SetDescription("博客用户")

	return blogAuth
}
