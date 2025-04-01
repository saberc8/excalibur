package global

import (
	"aquila/config"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AquilaConfig config.Configuration
	AquilaViper  *viper.Viper
	AquilaLog    *zap.Logger
	AquilaDb     *gorm.DB // DB 数据库连接
	AquilaRedis  *redis.Client
	AquilaMinio  *minio.Client // 新增 Minio 客户端
)

type GModel struct {
	ID        int64          `gorm:"column:id;type:int;primaryKey;autoIncrement:true;comment:主键id" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间" json:"createdAt"` // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间" json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:软删除时间" json:"-"`
}
