package assetView

import (
	"dataxWeb/models/util/lg"
	"dataxWeb/models/util/tool/db/dbConnection"
)

func libraryListQuery(databaseOpen dbConnection.DatabaseInterface) ([]libraryListJsonStruct, error) {
	err := databaseOpen.Open()
	if err != nil {
		lg.Error(err)
		return nil, err
	}

	defer databaseOpen.Close()

	// 先查库
	var schemataList []struct {
		SchemaName string `db:"SCHEMA_NAME"` //
	}
	querySQL := `select SCHEMA_NAME FROM information_schema.SCHEMATA where SCHEMA_NAME not in ('mysql','information_schema','performance_schema','sys');`
	if err := databaseOpen.Select(&schemataList, querySQL); err != nil {
		lg.Error(err)
		return nil, err
	}

	// 在查表
	var libraryListJson []libraryListJsonStruct
	for _, schemata := range schemataList {
		var schemataList []libraryListJsonStruct
		var columns []struct {
			TableName string `db:"TABLE_NAME"` //
		}

		querySQL = `SELECT TABLE_NAME FROM information_schema.TABLES where TABLE_SCHEMA = ? `
		if err := databaseOpen.Select(&columns, querySQL, schemata.SchemaName); err != nil {
			panic(err)
		}
		for _, column := range columns {
			schemataList = append(schemataList, libraryListJsonStruct{
				Type:       1,
				Id:         column.TableName,
				Label:      column.TableName,
				SchemaName: schemata.SchemaName,
				Children:   []libraryListJsonStruct{},
			})
		}

		libraryListJson = append(libraryListJson, libraryListJsonStruct{
			Type:     0,
			Id:       schemata.SchemaName,
			Label:    schemata.SchemaName,
			Children: schemataList,
		})
	}

	return libraryListJson, nil
}
