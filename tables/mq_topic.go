package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMqTopicTable(ctx *context.Context) table.Table {

	mqTopic := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := mqTopic.GetInfo().HideFilterArea()
	info.HideEditButton()
	info.HideDeleteButton()
	info.HideNewButton()

	info.AddField("Id", "id", db.Int).FieldHide()
	info.AddField("TOPIC_NAME", "TOPIC_NAME", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).FieldCopyable()
	info.AddField("TOPIC_TYPE", "TOPIC_TYPE", db.Int).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "1" {
			return "普通消息"
		}
		if model.Value == "2" {
			return "定时/延时消息"
		}
		if model.Value == "3" {
			return "事务消息"
		}
		return "unknown"
	}).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "1", Text: "1:普通消息"},
		{Value: "2", Text: "2:定时/延时消息"},
		{Value: "3", Text: "3:事务消息"},
	})
	info.AddField("TOPIC_STATUS", "TOPIC_STATUS", db.Varchar)
	info.AddField("OWNER_USERID", "OWNER_USERID", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike}).FieldCopyable()
	info.AddField("DESCRIPTION", "DESCRIPTION", db.Text)
	info.AddField("CREATETIME", "CREATETIME", db.Datetime).FieldSortable()

	info.SetTable("mq_topic").SetTitle("MQTT Topic列表").SetDescription("MqTopic")

	formList := mqTopic.GetForm()
	formList.AddField("Id", "id", db.Int, form.Number)
	formList.AddField("TOPIC_NAME", "TOPIC_NAME", db.Varchar, form.Text)
	formList.AddField("TOPIC_TYPE", "TOPIC_TYPE", db.Int, form.Number)
	formList.AddField("TOPIC_STATUS", "TOPIC_STATUS", db.Varchar, form.Text)
	formList.AddField("OWNER_USERID", "OWNER_USERID", db.Varchar, form.Text)
	formList.AddField("DESCRIPTION", "DESCRIPTION", db.Text, form.RichText)
	formList.AddField("CREATETIME", "CREATETIME", db.Datetime, form.Datetime)

	formList.SetTable("mq_topic").SetTitle("MqTopic").SetDescription("MqTopic")

	return mqTopic
}
