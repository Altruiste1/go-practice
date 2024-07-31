package global

import (
	"github.com/elastic/go-elasticsearch/v8"
	"go-practice/go_zero/shop/internal/config"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	EsClient *elasticsearch.TypedClient
	Cfg      *config.Config
)

func InitGlobal() {
	initElasticSearch()
	initMysql()
}
