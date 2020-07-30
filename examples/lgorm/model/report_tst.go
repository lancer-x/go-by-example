package model

import (
	"database/sql"
	"gorm.io/gorm"
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


CREATE TABLE `report_test` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `type` varchar(20) NOT NULL DEFAULT '' COMMENT '媒体类型',
  `callback` varchar(100) NOT NULL DEFAULT '' COMMENT '头条',
  `source` varchar(50) NOT NULL DEFAULT '' COMMENT '头条',
  `conv_time` varchar(30) NOT NULL DEFAULT '' COMMENT '头条',
  `event_type` int unsigned NOT NULL DEFAULT '0' COMMENT '头条',
  `header` varchar(200) NOT NULL DEFAULT '' COMMENT '神马',
  `body` varchar(200) NOT NULL DEFAULT '' COMMENT '神马',
  `account_id` varchar(200) NOT NULL DEFAULT '' COMMENT '广点通',
  `user_action_set_id` varchar(200) NOT NULL DEFAULT '' COMMENT '广点通',
  `actions` varchar(200) NOT NULL DEFAULT '' COMMENT '广点通',
  `event_time` varchar(200) NOT NULL DEFAULT '' COMMENT '快手',
  `eventTime` varchar(200) NOT NULL DEFAULT '' COMMENT '快手',
  `jsEventType` int unsigned NOT NULL DEFAULT '0' COMMENT '快手线索',
  `eventType` int unsigned NOT NULL DEFAULT '0' COMMENT '快手线索',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间,即折损时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `token` varchar(100) NOT NULL DEFAULT '',
  `conversionTypes` varchar(200) NOT NULL DEFAULT '',
  `purchase_amount` varchar(10) NOT NULL DEFAULT '' COMMENT '快手金额',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='growth回传测试记录表'

JSON Sample
-------------------------------------
{    "body": "oceJYfZPrpqgBXZLamFAYxMty",    "user_action_set_id": "XxSVLjeNsooODKUgYgHLqBDqv",    "conversion_types": "OcmrKLTDuNikZPbdtKEEdVDDg",    "source": "jvAifFpZNdrTBaHPsedtkIeFZ",    "token": "ZAhNLhNPwgSQBxbXFxcDcrLeU",    "callback": "LEFOnupSAbfWXPuPIoNxCvUQc",    "header": "btocQeTxlEBrtQyHZKuFnifyW",    "account_id": "WeooFLuywGquXtRcaTcvXryyM",    "js_event_type": 1,    "updated_at": "2224-06-04T21:58:01.223002553+08:00",    "created_at": "2116-08-04T07:29:40.062646512+08:00",    "purchase_amount": "fZlaQNLFOwDTLTfhvvyffkHZC",    "id": 21,    "type": "RBUAIrOIEuNjDUyvUyWXcyCAq",    "conv_time": "klUjSsoCXrdGjolmpuLdtjHIo",    "actions": "eAFViMppYFhfxWIUZHQHPBXMt"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 5] column is set for unsigned
[13] column is set for unsigned
[14] column is set for unsigned



*/

// ReportTest struct is a row record of the report_test table in the go_by_example database
type ReportTest struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" json:"id"` // 自增id
	//[ 1] type                                           varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	Type string `gorm:"column:type;type:varchar;size:20;" json:"type"` // 媒体类型
	//[ 2] callback                                       varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Callback string `gorm:"column:callback;type:varchar;size:100;" json:"callback"` // 头条
	//[ 3] source                                         varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	Source string `gorm:"column:source;type:varchar;size:50;" json:"source"` // 头条
	//[ 4] conv_time                                      varchar(30)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 30      default: []
	ConvTime string `gorm:"column:conv_time;type:varchar;size:30;" json:"conv_time"` // 头条
	//[ 5] event_type                                     uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EventType uint32 `gorm:"column:event_type;type:uint;default:0;" json:"event_type"` // 头条
	//[ 6] header                                         varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Header string `gorm:"column:header;type:varchar;size:200;" json:"header"` // 神马
	//[ 7] body                                           varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Body string `gorm:"column:body;type:varchar;size:200;" json:"body"` // 神马
	//[ 8] account_id                                     varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	AccountID string `gorm:"column:account_id;type:varchar;size:200;" json:"account_id"` // 广点通
	//[ 9] user_action_set_id                             varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	UserActionSetID string `gorm:"column:user_action_set_id;type:varchar;size:200;" json:"user_action_set_id"` // 广点通
	//[10] actions                                        varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	Actions string `gorm:"column:actions;type:varchar;size:200;" json:"actions"` // 广点通
	//[11] event_time                                     varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	EventTime string `gorm:"column:event_time;type:varchar;size:200;" json:"event_time"` // 快手
	//[12] eventTime                                      varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	EventTimealt1 string `gorm:"column:eventTime;type:varchar;size:200;" json:"event_time"` // 快手
	//[13] jsEventType                                    uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	JsEventType uint32 `gorm:"column:jsEventType;type:uint;default:0;" json:"js_event_type"` // 快手线索
	//[14] eventType                                      uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	EventTypealt1 uint32 `gorm:"column:eventType;type:uint;default:0;" json:"event_type"` // 快手线索
	//[15] created_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间,即折损时间
	//[16] updated_at                                     timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
	//[17] token                                          varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	Token string `gorm:"column:token;type:varchar;size:100;" json:"token"`
	//[18] conversionTypes                                varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	ConversionTypes string `gorm:"column:conversionTypes;type:varchar;size:200;" json:"conversion_types"`
	//[19] purchase_amount                                varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	PurchaseAmount string `gorm:"column:purchase_amount;type:varchar;size:10;" json:"purchase_amount"` // 快手金额

}

var report_testTableInfo = &TableInfo{
	Name: "report_test",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            `自增id`,
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
			Name:               "type",
			Comment:            `媒体类型`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "Type",
			GoFieldType:        "string",
			JSONFieldName:      "type",
			ProtobufFieldName:  "type",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "callback",
			Comment:            `头条`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Callback",
			GoFieldType:        "string",
			JSONFieldName:      "callback",
			ProtobufFieldName:  "callback",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "source",
			Comment:            `头条`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "Source",
			GoFieldType:        "string",
			JSONFieldName:      "source",
			ProtobufFieldName:  "source",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "conv_time",
			Comment:            `头条`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       30,
			GoFieldName:        "ConvTime",
			GoFieldType:        "string",
			JSONFieldName:      "conv_time",
			ProtobufFieldName:  "conv_time",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "event_type",
			Comment:            `头条`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EventType",
			GoFieldType:        "uint32",
			JSONFieldName:      "event_type",
			ProtobufFieldName:  "event_type",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "header",
			Comment:            `神马`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Header",
			GoFieldType:        "string",
			JSONFieldName:      "header",
			ProtobufFieldName:  "header",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "body",
			Comment:            `神马`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Body",
			GoFieldType:        "string",
			JSONFieldName:      "body",
			ProtobufFieldName:  "body",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "account_id",
			Comment:            `广点通`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "AccountID",
			GoFieldType:        "string",
			JSONFieldName:      "account_id",
			ProtobufFieldName:  "account_id",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "user_action_set_id",
			Comment:            `广点通`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "UserActionSetID",
			GoFieldType:        "string",
			JSONFieldName:      "user_action_set_id",
			ProtobufFieldName:  "user_action_set_id",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "actions",
			Comment:            `广点通`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "Actions",
			GoFieldType:        "string",
			JSONFieldName:      "actions",
			ProtobufFieldName:  "actions",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "event_time",
			Comment:            `快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "EventTime",
			GoFieldType:        "string",
			JSONFieldName:      "event_time",
			ProtobufFieldName:  "event_time",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "eventTime",
			Comment:            `快手`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "EventTimealt1",
			GoFieldType:        "string",
			JSONFieldName:      "event_timealt1",
			ProtobufFieldName:  "event_timealt1",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "jsEventType",
			Comment:            `快手线索`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "JsEventType",
			GoFieldType:        "uint32",
			JSONFieldName:      "js_event_type",
			ProtobufFieldName:  "js_event_type",
			ProtobufType:       "uint32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "eventType",
			Comment:            `快手线索`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "EventTypealt1",
			GoFieldType:        "uint32",
			JSONFieldName:      "event_typealt1",
			ProtobufFieldName:  "event_typealt1",
			ProtobufType:       "uint32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "created_at",
			Comment:            `创建时间,即折损时间`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "token",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "Token",
			GoFieldType:        "string",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "conversionTypes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "ConversionTypes",
			GoFieldType:        "string",
			JSONFieldName:      "conversion_types",
			ProtobufFieldName:  "conversion_types",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "purchase_amount",
			Comment:            `快手金额`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "PurchaseAmount",
			GoFieldType:        "string",
			JSONFieldName:      "purchase_amount",
			ProtobufFieldName:  "purchase_amount",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *ReportTest) TableName() string {
	return "report_test"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *ReportTest) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *ReportTest) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *ReportTest) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *ReportTest) TableInfo() *TableInfo {
	return report_testTableInfo
}
