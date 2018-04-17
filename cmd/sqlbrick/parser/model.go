// Copyright (c) 2018-present Anbillon Team (anbillonteam@gmail.com).
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package parser

var defaultDDLTag = map[string]string{
	"CREATE":   "CreateTable",
	"ALTER":    "AlterTable",
	"DROP":     "DropTable",
	"TRUNCATE": "TruncateTable",
	"RENAME":   "RenameTable",
}

var fieldTypes = map[string]string{
	"bool":        "bool",
	"boolean":     "bool",
	"tinyint":     "int8",
	"smallint":    "int16",
	"integer":     "int32",
	"int":         "int",
	"bigint":      "int64",
	"smallserial": "int16",
	"serial":      "int32",
	"bigserial":   "int64",
	"float":       "float",
	"double":      "float64",
	"real":        "float64",
	"numeric":     "float64",
	"decimal":     "float64",
	"text":        "string",
	"varchar":     "string",
	"char":        "string",
	"bit":         "string",
	"interval":    "string",
	"money":       "string",
	"cidr":        "string",
	"inet":        "string",
	"macaddr":     "string",
	"uuid":        "string",
	"json":        "string",
	"xml":         "string",
	"year":        "string",
	"date":        "time.Time",
	"datetime":    "time.Time",
	"timestamp":   "time.Time",
	"time":        "time.Time",
	"binary":      "[]byte",
	"varbinary":   "[]byte",
	"tinyblob":    "[]byte",
	"blob":        "[]byte",
	"mediumblob":  "[]byte",
	"longblob":    "[]byte",
	"bytea":       "[]byte",
}

var nullableFieldTypes = map[string]string{
	"bool":        "sql.NullBool",
	"boolean":     "sql.NullBool",
	"tinyint":     "sql.NullInt64",
	"smallint":    "sql.NullInt64",
	"integer":     "sql.NullInt64",
	"int":         "sql.NullInt64",
	"bigint":      "sql.NullInt64",
	"smallserial": "sql.NullInt64",
	"serial":      "sql.NullInt64",
	"bigserial":   "sql.NullInt64",
	"float":       "sql.NullFloat64",
	"double":      "sql.NullFloat64",
	"real":        "sql.NullFloat64",
	"numeric":     "sql.NullFloat64",
	"decimal":     "sql.NullFloat64",
	"text":        "sql.NullString",
	"varchar":     "sql.NullString",
	"char":        "sql.NullString",
	"bit":         "sql.NullString",
	"interval":    "sql.NullString",
	"money":       "sql.NullString",
	"cidr":        "sql.NullString",
	"inet":        "sql.NullString",
	"macaddr":     "sql.NullString",
	"uuid":        "sql.NullString",
	"json":        "sql.NullString",
	"xml":         "sql.NullString",
	"year":        "sql.NullString",
	"date":        "typex.NullTime",
	"datetime":    "typex.NullTime",
	"timestamp":   "typex.NullTime",
	"time":        "typex.NullTime",
	"binary":      "typex.NullBytes",
	"varbinary":   "typex.NullBytes",
	"tinyblob":    "typex.NullBytes",
	"blob":        "typex.NullBytes",
	"mediumblob":  "typex.NullBytes",
	"longblob":    "typex.NullBytes",
	"bytea":       "typex.NullBytes",
}

type QueryType int8

const (
	QueryTypeInvalid QueryType = 0
	QueryTypeInsert  QueryType = 1
	QueryTypeDelete  QueryType = 2
	QueryTypeUpdate  QueryType = 3
	QueryTypeSelect  QueryType = 4
)

var queryTypes = map[string]QueryType{
	"INSERT": QueryTypeInsert,
	"DELETE": QueryTypeDelete,
	"UPDATE": QueryTypeUpdate,
	"SELECT": QueryTypeSelect,
}

type Syntax struct {
	DbFieldName string
	FieldName   string
	FieldType   string
}

type Statement struct {
	QueryType QueryType
	Query     string
	Args      []string
}
