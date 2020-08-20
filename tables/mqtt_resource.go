package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMqttResourceTable(ctx *context.Context) table.Table {

	mqttResource := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := mqttResource.GetInfo().HideFilterArea()
	info.HideEditButton()
	info.HideDeleteButton()
	info.HideNewButton()

	info.AddField("Id", "id", db.Int).FieldHide()
	info.AddField("MQTTRESOURCE_ID", "MQTTRESOURCE_ID", db.Varchar).FieldCopyable()
	info.AddField("REGION_ID", "REGION_ID", db.Varchar)
	info.AddField("CUSTOMER_ID", "CUSTOMER_ID", db.Varchar).FieldHide()
	info.AddField("USER_ID", "USER_ID", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).FieldCopyable()
	info.AddField("IS_DELETE", "IS_DELETE", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "0" {
			return "on service"
		}
		if model.Value == "1" {
			return "deleted"
		}
		return "unknown"
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "0", Text: "0:on service"},
		{Value: "1", Text: "1:deleted"},
	})
	info.AddField("CREATE_TIME", "CREATE_TIME", db.Datetime).FieldSortable()
	info.AddField("MODIFY_TIME", "MODIFY_TIME", db.Datetime).FieldSortable()

	info.SetTable("mqtt_resource").SetTitle("MQTT实例列表").SetDescription("MqttResource")

	formList := mqttResource.GetForm()
	formList.AddField("Id", "id", db.Int, form.Number)
	formList.AddField("MQTTRESOURCE_ID", "MQTTRESOURCE_ID", db.Varchar, form.Text)
	formList.AddField("REGION_ID", "REGION_ID", db.Varchar, form.Text)
	formList.AddField("CUSTOMER_ID", "CUSTOMER_ID", db.Varchar, form.Text)
	formList.AddField("USER_ID", "USER_ID", db.Varchar, form.Text)
	formList.AddField("IS_DELETE", "IS_DELETE", db.Tinyint, form.Number)
	formList.AddField("CREATE_TIME", "CREATE_TIME", db.Datetime, form.Datetime)
	formList.AddField("MODIFY_TIME", "MODIFY_TIME", db.Datetime, form.Datetime)

	formList.SetTable("mqtt_resource").SetTitle("MqttResource").SetDescription("MqttResource")

	return mqttResource
}
