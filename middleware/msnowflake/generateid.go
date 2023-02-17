package msnowflake

import "github.com/GUAIK-ORG/go-snowflake/snowflake"

// NewSnowflake(datacenterid, workerid int64) (*Snowflake, error)
// 参数1 (int64): 数据中心ID (可用范围:0-31)
// 参数2 (int64): 机器ID    (可用范围:0-31)
// 返回1 (*Snowflake): Snowflake对象 | nil
// 返回2 (error): 错误码

var GenerateID, _ = snowflake.NewSnowflake(int64(0), int64(0))
