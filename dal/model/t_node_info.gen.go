// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameTNodeInfo = "t_node_info"

// TNodeInfo mapped from table <t_node_info>
type TNodeInfo struct {
	ID              int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UpdateTimestamp time.Time `gorm:"column:update_timestamp;not null;default:CURRENT_TIMESTAMP" json:"update_timestamp"`
	InsertTimestamp time.Time `gorm:"column:insert_timestamp;not null;default:CURRENT_TIMESTAMP" json:"insert_timestamp"`
	ChainID         int64     `gorm:"column:chain_id;not null" json:"chain_id"`
	RPCURL          string    `gorm:"column:rpc_url;not null" json:"rpc_url"`
	Type            int32     `gorm:"column:type;not null" json:"type"`
	Usability       int32     `gorm:"column:usability;default:100" json:"usability"`
}

// TableName TNodeInfo's table name
func (*TNodeInfo) TableName() string {
	return TableNameTNodeInfo
}
