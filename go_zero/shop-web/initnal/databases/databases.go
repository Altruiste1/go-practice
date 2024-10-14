package databases

import (
	"github.com/elastic/go-elasticsearch/v8"
	"gorm.io/gorm"
)

var MysqlDefaultDB *gorm.DB

var EsClient *elasticsearch.TypedClient

func InitDatabases() {
	initMysql()
	initElasticSearch()
}
