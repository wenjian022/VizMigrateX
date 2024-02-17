package dbConnection

import "strings"

// IsSelectQuery
//
//	@Description: 判断是否为 SELECT 查询
//	@param query
//	@return bool
func IsSelectQuery(query string) bool {
	query = strings.TrimSpace(strings.ToLower(query))
	return strings.HasPrefix(query, "select ")
}
