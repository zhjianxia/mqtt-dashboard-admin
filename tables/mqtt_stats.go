package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMqttStatsTable(ctx *context.Context) table.Table {

	mqttStats := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := mqttStats.GetInfo().HideFilterArea()
	info.HideEditButton()
	info.HideDeleteButton()
	info.HideNewButton()

	info.AddField("Id", "id", db.Bigint).FieldHide()
	info.AddField("父Topic", "PRETOPIC", db.Varchar).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("子Topic", "SUBTOPIC", db.Varchar)
	info.AddField("MQTTTYPE", "MQTTTYPE", db.Varchar)
	info.AddField("QOS", "QOS", db.Int)
	info.AddField("COUNT", "COUNT", db.Bigint)
	info.AddField("TPS", "TPS", db.Double)
	info.AddField("DATE", "DATE", db.Date).FieldSortable()
	info.AddField("TIMESTAMP", "TIMESTAMP", db.Bigint)
	//non exist column
	info.AddField("QOSDESC", "qosdesc", db.Varchar).FieldJoin(types.Join{
		Table:     "qosdesc", // 连表的表名
		Field:     "qos",     // 要连表的字段
		JoinField: "qosid",   // 连表的表的字段
	})

	info.SetTable("mqtt_stats").SetTitle("MQTT用量统计").SetDescription("MqttStats")

	formList := mqttStats.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Number)
	formList.AddField("PRETOPIC", "PRETOPIC", db.Varchar, form.Text)
	formList.AddField("SUBTOPIC", "SUBTOPIC", db.Varchar, form.Text)
	formList.AddField("MQTTTYPE", "MQTTTYPE", db.Varchar, form.Text)
	formList.AddField("QOS", "QOS", db.Int, form.Number)
	formList.AddField("COUNT", "COUNT", db.Bigint, form.Number)
	formList.AddField("TPS", "TPS", db.Double, form.Text)
	formList.AddField("DATE", "DATE", db.Date, form.Datetime)
	formList.AddField("TIMESTAMP", "TIMESTAMP", db.Bigint, form.Number)

	formList.SetTable("mqtt_stats").SetTitle("MqttStats").SetDescription("MqttStats")

	return mqttStats
}
