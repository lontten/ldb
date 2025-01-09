package ldb

import "strings"

func gen(num int) string {
	var queryArr []string
	for i := 0; i < num; i++ {
		queryArr = append(queryArr, "?")
	}
	return strings.Join(queryArr, ",")
}

var dangNamesMap = map[string]bool{
	"accessible":                    true,
	"add":                           true,
	"all":                           true,
	"alter":                         true,
	"analyze":                       true,
	"and":                           true,
	"as":                            true,
	"asc":                           true,
	"asensitive":                    true,
	"before":                        true,
	"between":                       true,
	"bigint":                        true,
	"binary":                        true,
	"blob":                          true,
	"both":                          true,
	"by":                            true,
	"call":                          true,
	"cascade":                       true,
	"case":                          true,
	"change":                        true,
	"char":                          true,
	"character":                     true,
	"check":                         true,
	"collate":                       true,
	"column":                        true,
	"condition":                     true,
	"constraint":                    true,
	"continue":                      true,
	"convert":                       true,
	"create":                        true,
	"cross":                         true,
	"cube":                          true,
	"current_date":                  true,
	"current_time":                  true,
	"current_user":                  true,
	"cursor":                        true,
	"database":                      true,
	"databases":                     true,
	"day_hour":                      true,
	"day_microsecond":               true,
	"day_minute":                    true,
	"day_second":                    true,
	"dec":                           true,
	"decimal":                       true,
	"declare":                       true,
	"default":                       true,
	"delayed":                       true,
	"delete":                        true,
	"desc":                          true,
	"describe":                      true,
	"deterministic":                 true,
	"distinct":                      true,
	"distinctrow":                   true,
	"div":                           true,
	"double":                        true,
	"drop":                          true,
	"dual":                          true,
	"each":                          true,
	"else":                          true,
	"elseif":                        true,
	"enclosed":                      true,
	"escaped":                       true,
	"exists":                        true,
	"exit":                          true,
	"explain":                       true,
	"false":                         true,
	"fetch":                         true,
	"float":                         true,
	"float4":                        true,
	"float8":                        true,
	"for":                           true,
	"force":                         true,
	"foreign":                       true,
	"from":                          true,
	"fulltext":                      true,
	"generated":                     true,
	"get":                           true,
	"grant":                         true,
	"group":                         true,
	"having":                        true,
	"high_priority":                 true,
	"hour_microsecond":              true,
	"hour_minute":                   true,
	"hour_second":                   true,
	"if":                            true,
	"ignore":                        true,
	"in":                            true,
	"index":                         true,
	"infile":                        true,
	"inner":                         true,
	"inout":                         true,
	"insensitive":                   true,
	"insert":                        true,
	"int":                           true,
	"int1":                          true,
	"int2":                          true,
	"int3":                          true,
	"int4":                          true,
	"int8":                          true,
	"integer":                       true,
	"interval":                      true,
	"into":                          true,
	"io_after_gtids":                true,
	"io_before_gtids":               true,
	"is":                            true,
	"iterate":                       true,
	"join":                          true,
	"key":                           true,
	"keys":                          true,
	"kill":                          true,
	"leading":                       true,
	"leave":                         true,
	"left":                          true,
	"like":                          true,
	"limit":                         true,
	"linear":                        true,
	"lines":                         true,
	"load":                          true,
	"localtime":                     true,
	"localtimestamp":                true,
	"lock":                          true,
	"long":                          true,
	"longblob":                      true,
	"longtext":                      true,
	"loop":                          true,
	"low_priority":                  true,
	"master_bind":                   true,
	"master_ssl_verify_server_cert": true,
	"match":                         true,
	"maxvalue":                      true,
	"mediumblob":                    true,
	"mediumint":                     true,
	"mediumtext":                    true,
	"middleint":                     true,
	"minute_microsecond":            true,
	"minute_second":                 true,
	"mod":                           true,
	"modifies":                      true,
	"natural":                       true,
	"not":                           true,
	"no_write_to_binlog":            true,
	"null":                          true,
	"numeric":                       true,
	"on":                            true,
	"optimize":                      true,
	"option":                        true,
	"optionally":                    true,
	"or":                            true,
	"order":                         true,
	"out":                           true,
	"outer":                         true,
	"outfile":                       true,
	"partition":                     true,
	"precision":                     true,
	"primary":                       true,
	"procedure":                     true,
	"purge":                         true,
	"range":                         true,
	"read":                          true,
	"reads":                         true,
	"read_write":                    true,
	"real":                          true,
	"references":                    true,
	"regexp":                        true,
	"release":                       true,
	"rename":                        true,
	"repeat":                        true,
	"replace":                       true,
	"require":                       true,
	"resignal":                      true,
	"restrict":                      true,
	"return":                        true,
	"revoke":                        true,
	"right":                         true,
	"rlike":                         true,
	"row":                           true,
	"rows":                          true,
	"schema":                        true,
	"schemas":                       true,
	"second_microsecond":            true,
	"select":                        true,
	"sensitive":                     true,
	"separator":                     true,
	"set":                           true,
	"show":                          true,
	"signal":                        true,
	"smallint":                      true,
	"spatial":                       true,
	"specific":                      true,
	"sql":                           true,
	"sqlexception":                  true,
	"sqlstate":                      true,
	"sqlwarning":                    true,
	"sql_big_result":                true,
	"sql_calc_found_rows":           true,
	"sql_small_result":              true,
	"ssl":                           true,
	"starting":                      true,
	"stored":                        true,
	"straight_join":                 true,
	"table":                         true,
	"terminated":                    true,
	"then":                          true,
	"tinyblob":                      true,
	"tinyint":                       true,
	"tinytext":                      true,
	"to":                            true,
	"trailing":                      true,
	"trigger":                       true,
	"true":                          true,
	"undo":                          true,
	"union":                         true,
	"unique":                        true,
	"unlock":                        true,
	"unsigned":                      true,
	"update":                        true,
	"usage":                         true,
	"use":                           true,
	"using":                         true,
	"utc_date":                      true,
	"utc_time":                      true,
	"utc_timestamp":                 true,
	"values":                        true,
	"varbinary":                     true,
	"varchar":                       true,
	"varcharacter":                  true,
	"varying":                       true,
	"virtual":                       true,
	"when":                          true,
	"where":                         true,
	"while":                         true,
	"with":                          true,
	"write":                         true,
	"xor":                           true,
	"year_month":                    true,
	"zerofill":                      true,
}

func genSelectCols(list []string) string {
	var list2 []string
	for _, s := range list {
		_, has := dangNamesMap[s]
		if has {
			list2 = append(list2, "`"+s+"`")
		} else {
			list2 = append(list2, s)
		}
	}
	return strings.Join(list2, ",")
}