package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `growth_publisher_channel_config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `line` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '业务线，见后端配置',
  `ca_s` varchar(50) NOT NULL DEFAULT '' COMMENT '一级渠道',
  `ca_s_name` varchar(50) NOT NULL DEFAULT '' COMMENT '一级渠道名称',
  `ca_n` varchar(50) NOT NULL DEFAULT '' COMMENT '二级渠道',
  `ca_n_name` varchar(50) NOT NULL DEFAULT '' COMMENT '二级渠道名称',
  `is_cas_all` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否回传某一级渠道下全部二级渠道',
  `data_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '回传数据类型',
  `publisher_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '媒体类型，见后端配置',
  `account_config_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '媒体回传配置id',
  `is_valid` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否有效1是，0否',
  `creator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建人uid',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `category` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '大渠道id',
  `is_delete` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否删除  1是，0否',
  PRIMARY KEY (`id`),
  KEY `idx_cainfo` (`ca_s`,`ca_n`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8 COMMENT='渠道回传配置表'

JSON Sample
-------------------------------------
{    "data_type": 25,    "account_config_id": 98,    "is_delete": 14,    "line": 75,    "ca_s_name": "AwygADIibYaHesSaihleflsvE",    "ca_n_name": "FULUuLEldiYIkfVLZTxdBlkcm",    "ca_s": "EOZnciFPpnyvRSYIZEATSAHmd",    "ca_n": "BoTrmeOsPAjRNBDSGaujUZDJB",    "category": 38,    "creator_id": 76,    "id": 83,    "publisher_type": 65,    "is_valid": 3,    "is_cas_all": 3,    "created_at": "2160-12-25T21:05:45.352452471+08:00",    "updated_at": "2246-09-06T17:27:42.581677156+08:00"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[10] column is set for unsigned
[11] column is set for unsigned
[14] column is set for unsigned
[15] column is set for unsigned



*/

// GrowthPublisherChannelConfig struct is a row record of the growth_publisher_channel_config table in the guazi_sem database
type GrowthPublisherChannelConfig struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 自增主键
	//[ 1] line                                           uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [0]
	Line uint32 `gorm:"column:line;type:uint;default:0;" json:"line"` // 业务线，见后端配置
	//[ 2] ca_s                                           varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CaS string `gorm:"column:ca_s;type:varchar(50);size:50;" json:"ca_s"` // 一级渠道
	//[ 3] ca_s_name                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CaSName string `gorm:"column:ca_s_name;type:varchar(50);size:50;" json:"ca_s_name"` // 一级渠道名称
	//[ 4] ca_n                                           varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CaN string `gorm:"column:ca_n;type:varchar(50);size:50;" json:"ca_n"` // 二级渠道
	//[ 5] ca_n_name                                      varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	CaNName string `gorm:"column:ca_n_name;type:varchar(50);size:50;" json:"ca_n_name"` // 二级渠道名称
	//[ 6] is_cas_all                                     uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [1]
	IsCasAll uint32 `gorm:"column:is_cas_all;type:uint;default:1;" json:"is_cas_all"` // 是否回传某一级渠道下全部二级渠道
	//[ 7] data_type                                      uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [0]
	DataType uint32 `gorm:"column:data_type;type:uint;default:0;" json:"data_type"` // 回传数据类型
	//[ 8] publisher_type                                 uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [0]
	PublisherType uint32 `gorm:"column:publisher_type;type:uint;default:0;" json:"publisher_type"` // 媒体类型，见后端配置
	//[ 9] account_config_id                              uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	AccountConfigID uint32 `gorm:"column:account_config_id;type:uint;default:0;" json:"account_config_id"` // 媒体回传配置id
	//[10] is_valid                                       uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [1]
	IsValid uint32 `gorm:"column:is_valid;type:uint;default:1;" json:"is_valid"` // 是否有效1是，0否
	//[11] creator_id                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CreatorID uint32 `gorm:"column:creator_id;type:uint;default:0;" json:"creator_id"` // 创建人uid
	//[12] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	//[13] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	//[14] category                                       uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [0]
	Category uint32 `gorm:"column:category;type:uint;default:0;" json:"category"` // 大渠道id
	//[15] is_delete                                      uint             null: false  primary: false  isArray: false  auto: false  col: uint        len: -1      default: [0]
	IsDelete uint32 `gorm:"column:is_delete;type:uint;default:0;" json:"is_delete"` // 是否删除  1是，0否

}

var growth_publisher_channel_configTableInfo = &TableInfo{
	Name: "growth_publisher_channel_config",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `自增主键`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "line",
			Comment:            `业务线，见后端配置`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Line",
			GoFieldType:        "uint32",
			JSONFieldName:      "line",
			ProtobufFieldName:  "line",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ca_s",
			Comment:            `一级渠道`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "CaS",
			GoFieldType:        "string",
			JSONFieldName:      "ca_s",
			ProtobufFieldName:  "ca_s",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ca_s_name",
			Comment:            `一级渠道名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "CaSName",
			GoFieldType:        "string",
			JSONFieldName:      "ca_s_name",
			ProtobufFieldName:  "ca_s_name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ca_n",
			Comment:            `二级渠道`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "CaN",
			GoFieldType:        "string",
			JSONFieldName:      "ca_n",
			ProtobufFieldName:  "ca_n",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "ca_n_name",
			Comment:            `二级渠道名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "CaNName",
			GoFieldType:        "string",
			JSONFieldName:      "ca_n_name",
			ProtobufFieldName:  "ca_n_name",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "is_cas_all",
			Comment:            `是否回传某一级渠道下全部二级渠道`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsCasAll",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_cas_all",
			ProtobufFieldName:  "is_cas_all",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "data_type",
			Comment:            `回传数据类型`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "DataType",
			GoFieldType:        "uint32",
			JSONFieldName:      "data_type",
			ProtobufFieldName:  "data_type",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "publisher_type",
			Comment:            `媒体类型，见后端配置`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "PublisherType",
			GoFieldType:        "uint32",
			JSONFieldName:      "publisher_type",
			ProtobufFieldName:  "publisher_type",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "account_config_id",
			Comment:            `媒体回传配置id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "AccountConfigID",
			GoFieldType:        "uint32",
			JSONFieldName:      "account_config_id",
			ProtobufFieldName:  "account_config_id",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "is_valid",
			Comment:            `是否有效1是，0否`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsValid",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_valid",
			ProtobufFieldName:  "is_valid",
			ProtobufType:       "uint32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "creator_id",
			Comment:            `创建人uid`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CreatorID",
			GoFieldType:        "uint32",
			JSONFieldName:      "creator_id",
			ProtobufFieldName:  "creator_id",
			ProtobufType:       "uint32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "created_at",
			Comment:            `创建时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CreatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "created_at",
			ProtobufFieldName:  "created_at",
			ProtobufType:       "uint64",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "updated_at",
			Comment:            `更新时间`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UpdatedAt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "category",
			Comment:            `大渠道id`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "Category",
			GoFieldType:        "uint32",
			JSONFieldName:      "category",
			ProtobufFieldName:  "category",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "is_delete",
			Comment:            `是否删除  1是，0否`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "IsDelete",
			GoFieldType:        "uint32",
			JSONFieldName:      "is_delete",
			ProtobufFieldName:  "is_delete",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (g *GrowthPublisherChannelConfig) TableName() string {
	return "growth_publisher_channel_config"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (g *GrowthPublisherChannelConfig) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (g *GrowthPublisherChannelConfig) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (g *GrowthPublisherChannelConfig) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (g *GrowthPublisherChannelConfig) TableInfo() *TableInfo {
	return growth_publisher_channel_configTableInfo
}
