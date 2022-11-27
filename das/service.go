package das

import (
	"database/sql"
	"log"
	"strconv"
	"strings"

	core "github.com/mt1976/ebEstimates/core"
	logs "github.com/mt1976/ebEstimates/logs"
)

func Query(db *sql.DB, query string) ([]map[string]interface{}, int, error) {

	//log.Println("Query:", query)
	//if query containts dbo.transalationStore then dont do anything
	if !(strings.Contains(query, "dbo.translationStore")) {
		logs.Query(query)
	}
	rows, _ := db.Query(query) // Note: Ignoring errors for brevity
	cols, _ := rows.Columns()
	noResults := 0
	recs := []map[string]interface{}{}
	defer rows.Close()

	for rows.Next() {
		noResults++
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		m := make(map[string]interface{})
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, 0, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		//m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		//fmt.Println(m)
		//log.Println("Append", m)
		recs = append(recs, m)
	}
	//rows.Close()
	//log.Println("Recs:", recs)
	//spew.Dump(recs)
	//log.Println("Query:", m)
	//log.Println("No Results:", noResults)
	if !(strings.Contains(query, "dbo.translationStore")) {
		logs.Result(query, strconv.Itoa(noResults))
	}
	return recs, noResults, nil
}

func Poke(DB *sql.DB) error {
	errordb := DB.Ping()
	if errordb != nil {
		log.Println(errordb.Error())
	}
	return errordb
}

// TODO: implement
func Execute(tsql string) {

	logs.Database("Execute :", tsql)

	_, err := core.ApplicationDB.Exec(tsql)
	if err != nil {
		logs.Panic("Execution Error", err)
	}

}

// TODO: implement
func Process(tsql string) sql.Result {

	logs.Database("Execute :", tsql)

	res, err := core.ApplicationDB.Exec(tsql)
	if err != nil {
		logs.Panic("Execution Error", err)
		return nil
	}
	return res
}
