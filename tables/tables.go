package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "mq_topic" => http://localhost:9033/admin/info/mq_topic
// "mqtt_resource" => http://localhost:9033/admin/info/mqtt_resource
// "mqtt_stats" => http://localhost:9033/admin/info/mqtt_stats
//
// "v_mqtt_resource" => http://localhost:9033/admin/info/v_mqtt_resource
//
// "servicedist" => http://localhost:9033/admin/info/servicedist
//
// example end
//
var Generators = map[string]table.Generator{
	"mq_topic":      GetMqTopicTable,
	"mqtt_stats":    GetMqttStatsTable,
	"mqtt_resource": GetMqttResourceTable,

	"blog_auth": GetBlogAuthTable,

	"servicedist": GetServicedistTable,

	// generators end
}
