package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	editType "github.com/GoAdminGroup/go-admin/template/types/table"
)

func GetServicedistTable(ctx *context.Context) table.Table {

	servicedist := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := servicedist.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).FieldHide()
	info.AddField("应用", "appname", db.Varchar).FieldWidth(80).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "MQTT", Text: "MQTT"},
		{Value: "EKAFKA", Text: "EKAFKA"},
	}).FieldEditAble(editType.Select).FieldEditOptions(types.FieldOptions{
		{Value: "MQTT", Text: "MQTT"},
		{Value: "EKAFKA", Text: "EKAFKA"},
	})
	info.AddField("资源池", "poolalias", db.Varchar).FieldWidth(80).FieldFilterable(types.FilterType{FormType: form.SelectSingle}).FieldFilterOptions(types.FieldOptions{
		{Value: "wuxi", Text: "苏州（无锡AZ）  "},
		{Value: "suzhou", Text: "苏州（苏州AZ）  "},
		{Value: "dongguan", Text: "广州（东莞AZ）  "},
		{Value: "zhengzhou2", Text: "郑州（港区AZ）  "},
		{Value: "zhengzhou1", Text: "郑州（高新区AZ）"},
		{Value: "yaan", Text: "成都（雅安AZ）  "},
		{Value: "beijing", Text: "北京（北京AZ）  "},
		{Value: "zhuzhou", Text: "长沙（株洲AZ）  "},
		{Value: "jinan", Text: "济南（济南AZ）  "},
		{Value: "chongqing", Text: "重庆（重庆AZ）  "},
		{Value: "shanghai", Text: "上海（上海AZ）  "},
		{Value: "ningbo", Text: "杭州（宁波AZ）  "},
		{Value: "xian", Text: "西安（西安AZ）  "},
		{Value: "guangzhou", Text: "广州节点        "},
		{Value: "beijing", Text: "北京节点        "},
		{Value: "changsha", Text: "长沙节点        "},
		{Value: "heyin-beijing", Text: "合营北京节点    "},
		{Value: "heyin-jiangsu", Text: "合营江苏节点    "},
		{Value: "heyin-fujian", Text: "合营福建节点    "},
		{Value: "heyin-liaonin", Text: "合营辽宁节点    "},
		{Value: "heyin-zhejiang", Text: "合营浙江节点    "},
	}).FieldEditAble(editType.Select).FieldEditOptions(types.FieldOptions{
		{Value: "wuxi", Text: "苏州（无锡AZ）  "},
		{Value: "suzhou", Text: "苏州（苏州AZ）  "},
		{Value: "dongguan", Text: "广州（东莞AZ）  "},
		{Value: "zhengzhou2", Text: "郑州（港区AZ）  "},
		{Value: "zhengzhou1", Text: "郑州（高新区AZ）"},
		{Value: "yaan", Text: "成都（雅安AZ）  "},
		{Value: "beijing", Text: "北京（北京AZ）  "},
		{Value: "zhuzhou", Text: "长沙（株洲AZ）  "},
		{Value: "jinan", Text: "济南（济南AZ）  "},
		{Value: "chongqing", Text: "重庆（重庆AZ）  "},
		{Value: "shanghai", Text: "上海（上海AZ）  "},
		{Value: "ningbo", Text: "杭州（宁波AZ）  "},
		{Value: "xian", Text: "西安（西安AZ）  "},
		{Value: "guangzhou", Text: "广州节点        "},
		{Value: "beijing", Text: "北京节点        "},
		{Value: "changsha", Text: "长沙节点        "},
		{Value: "heyin-beijing", Text: "合营北京节点    "},
		{Value: "heyin-jiangsu", Text: "合营江苏节点    "},
		{Value: "heyin-fujian", Text: "合营福建节点    "},
		{Value: "heyin-liaonin", Text: "合营辽宁节点    "},
		{Value: "heyin-zhejiang", Text: "合营浙江节点    "},
	})
	info.AddField("主机名", "hostname", db.Varchar).FieldWidth(250).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("管理IPV4", "ipv4_admin", db.Varchar).FieldWidth(100).FieldEditAble(editType.Text).
		FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})
	info.AddField("业务IPV4", "ipv4_service", db.Varchar).FieldHide()
	info.AddField("VIP", "VIP", db.Varchar).FieldEditAble(editType.Text)
	info.AddField("IPV6", "IPV6", db.Varchar).FieldHide()
	info.AddField("服务列表", "servicelist", db.Varchar).FieldEditAble(editType.Text).FieldFilterable(types.FilterType{Operator: types.FilterOperatorLike})

    info.AddSelectBox("应用名", types.FieldOptions{
		{Value: "MQTT", Text: "MQTT"},
		{Value: "EKAFKA", Text: "EKAFKA"},
	}, action.FieldFilter("appname"))

	info.SetTable("servicedist").SetTitle("服务部署列表").SetDescription("Servicedist")

	formList := servicedist.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default)
	formList.AddField("应用", "appname", db.Varchar, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Text: "MQTT", Value: "MQTT"},
		{Text: "EKAFKA", Value: "EKAFKA"},
	}).FieldDefault("MQTT")
	formList.AddField("资源池", "poolalias", db.Varchar, form.SelectSingle).FieldOptions(types.FieldOptions{
		{Value: "wuxi", Text: "苏州（无锡AZ）  "},
		{Value: "suzhou", Text: "苏州（苏州AZ）  "},
		{Value: "dongguan", Text: "广州（东莞AZ）  "},
		{Value: "zhengzhou1", Text: "郑州（港区AZ）  "},
		{Value: "zhengzhou2", Text: "郑州（高新区AZ）"},
		{Value: "yaan", Text: "成都（雅安AZ）  "},
		{Value: "beijing", Text: "北京（北京AZ）  "},
		{Value: "zhuzhou", Text: "长沙（株洲AZ）  "},
		{Value: "jinan", Text: "济南（济南AZ）  "},
		{Value: "chongqing", Text: "重庆（重庆AZ）  "},
		{Value: "shanghai", Text: "上海（上海AZ）  "},
		{Value: "ningbo", Text: "杭州（宁波AZ）  "},
		{Value: "xian", Text: "西安（西安AZ）  "},
		{Value: "guangzhou", Text: "广州节点        "},
		{Value: "beijing", Text: "北京节点        "},
		{Value: "changsha", Text: "长沙节点        "},
		{Value: "heyin-beijing", Text: "合营北京节点    "},
		{Value: "heyin-jiangsu", Text: "合营江苏节点    "},
		{Value: "heyin-fujian", Text: "合营福建节点    "},
		{Value: "heyin-liaonin", Text: "合营辽宁节点    "},
		{Value: "heyin-zhejiang", Text: "合营浙江节点    "},
	}).FieldDefault("suzhou")
	formList.AddField("主机名", "hostname", db.Varchar, form.Text)
	formList.AddField("管理IPV4", "IPV4_ADMIN", db.Varchar, form.Text)
	formList.AddField("业务IPV4", "IPV4_SERVICE", db.Varchar, form.Text)
	formList.AddField("VIP", "VIP", db.Varchar, form.Text)
	formList.AddField("IPV6", "IPV6", db.Varchar, form.Text)
	formList.AddField("服务列表", "servicelist", db.Varchar, form.SelectBox).FieldOptions(types.FieldOptions{
		{Text: "Nginx", Value: "Nginx"},
		{Text: "Tomcat", Value: "Tomcat"},
		{Text: "HAProxy", Value: "HAProxy"},
		{Text: "Keepalived", Value: "Keepalived"},
		{Text: "BCRDB", Value: "BCRDB"},
		{Text: "Redis-Sentinel哨兵", Value: "Redis-Sentinel"},
		{Text: "Redis-Master", Value: "Redis-Master"},
		{Text: "Redis-Slave", Value: "Redis-Slave"},
		{Text: "Zookeeper", Value: "Zookeeper"},
		{Text: "ActiveMQ", Value: "ActiveMQ"},
		{Text: "RabbitMQ", Value: "RabbitMQ"},
		{Text: "KafkaBroker", Value: "KafkaBroker"},
		{Text: "NTP", Value: "NTP"},
		{Text: "Kerberos", Value: "Kerberos"},
		{Text: "LDAP", Value: "LDAP"},
		{Text: "RocketMQ-Broker-A-Master", Value: "RocketMQ-Broker-A-Master"},
		{Text: "RocketMQ-Broker-A-Slave", Value: "RocketMQ-Broker-A-Slave"},
		{Text: "RocketMQ-Broker-B-Master", Value: "RocketMQ-Broker-B-Master"},
		{Text: "RocketMQ-Broker-B-Slave", Value: "RocketMQ-Broker-B-Slave"},
		{Text: "RocketMQ-Broker-C-Master", Value: "RocketMQ-Broker-C-Master"},
		{Text: "RocketMQ-Broker-C-Slave", Value: "RocketMQ-Broker-C-Slave"},
		{Text: "MQTT-Bridge", Value: "MQTT-Bridge"},
		{Text: "RocketMQ-NameServer路由服务", Value: "RocketMQ-NameServer"},
		{Text: "IAM鉴权服务", Value: "IAM"},
		{Text: "Measure计量采集模块", Value: "Measure"},
		{Text: "Tuxedo-mq-console控制台", Value: "Tuxedo-mq-console"},
	}).
		// 这里返回一个[]string，对应的值是本列的fruit字段的值，即编辑表单时会显示的对应值
		FieldDisplay(func(model types.FieldModel) interface{} {
			return []string{"pear"}
		})
	/*formList.AddField("servicelist", "servicelist", db.Varchar, form.SelectBox).
	// 多选的选项，text代表显示内容，value代表对应值
	FieldOptions(types.FieldOptions{
		{
			Text:  "apple",
			Value: "apple",
		}, {
			Text:  "banana",
			Value: "banana",
		}, {
			Text:  "watermelon",
			Value: "watermelon",
		}, {
			Text:  "pear",
			Value: "pear",
		},
	}).
	// 这里返回一个[]string，对应的值是本列的fruit字段的值，即编辑表单时会显示的对应值
	FieldDisplay(func(model types.FieldModel) interface{} {
		return []string{"pear"}
	})*/

	formList.SetTable("servicedist").SetTitle("Servicedist").SetDescription("Servicedist")

	return servicedist
}
