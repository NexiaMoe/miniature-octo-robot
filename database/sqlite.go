package database

import (
	"database/sql"
	"fmt"
	"helloion/entity"
	"helloion/utils"

	sq "github.com/Masterminds/squirrel"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Init(c *utils.Config) *sql.DB {
	// Initialize Sqlite3 database
	db, err := sql.Open(c.DBDriver, c.DBSource)

	utils.CheckErr(err)
	// Check if table exists
	_, err = db.Exec("SELECT * FROM log")

	if err != nil {
		fmt.Println("table log is not exists, creating table log")
		// Create table if not exists
		// id, ip, status_code, X-Amzn-Trace-Id
		_, err = db.Exec("CREATE TABLE log (id INTEGER PRIMARY KEY, ip TEXT, status_code INTEGER, x_amzn_trace_id TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
		utils.CheckErr(err)
	}

	DB = db

	return nil
}

func GetDB() *sql.DB {
	return DB
}

func SaveLog(l *entity.IP) error {
	fmt.Println(l)
	// Insert data to table log

	query := sq.Insert("log").Columns("ip", "status_code", "x_amzn_trace_id").Values(l.Origin, l.StatusCode, l.Headers.XAmznTraceID)

	sql, args, err := query.ToSql()

	utils.CheckErr(err)

	stmt, err := DB.Prepare(sql)

	utils.CheckErr(err)

	res, err := stmt.Exec(args...)

	utils.CheckErr(err)

	id, err := res.LastInsertId()

	utils.CheckErr(err)

	fmt.Println("Last Inserted ID : ", id)

	return nil
}

func GetLog() (logs []*entity.IP, err error) {
	rows, err := DB.Query("SELECT * FROM log")

	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		var log entity.IP
		var headers entity.Header
		err = rows.Scan(&log.Id,
			&log.Origin,
			&log.StatusCode,
			&headers.XAmznTraceID,
			&log.CreatedAt)

		utils.CheckErr(err)

		log.Headers = &headers

		logs = append(logs, &log)
	}

	return logs, nil
}
