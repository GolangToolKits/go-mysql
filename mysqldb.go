package gomysql

import (
	"database/sql"
	"log"

	//Required
	_ "github.com/go-sql-driver/mysql"
)

// MyDB MyDB
type MyDB struct {
	Host     string
	User     string
	Password string
	Database string
	db       *sql.DB
	err      error
}

// Connect Connect
func (m *MyDB) Connect() bool {
	var rtn = false
	var conStr = m.User + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Database
	m.db, m.err = sql.Open("mysql", conStr)
	if m.err == nil {
		m.err = m.db.Ping()
		if m.err != nil {
			log.Println("Database Connect Error:", m.err.Error())
		} else {
			rtn = true
		}
	}
	return rtn
}

// New NewDatabase
func (m *MyDB) New() Database {
	return m
}

// BeginTransaction BeginTransaction
func (m *MyDB) BeginTransaction() Transaction {
	var trans Transaction
	var myTrans MyDbTx
	tx, err := m.db.Begin()
	if err == nil && tx != nil {
		myTrans.Tx = tx
	}
	trans = &myTrans
	return trans
}

// Test Test
func (m *MyDB) Test(query string, args ...any) *DbRow {
	return m.Get(query, args...)
}

// Insert Insert
func (m *MyDB) Insert(query string, args ...any) (bool, int64) {
	var success = false
	var id int64 = -1
	var stmtIns *sql.Stmt
	stmtIns, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtIns.Close()
		res, err := stmtIns.Exec(args...)
		if err != nil {
			log.Println("Insert Exec err:", err.Error())
		} else {
			id, err = res.LastInsertId()
			affectedRows, _ := res.RowsAffected()
			if err == nil && affectedRows > 0 {
				success = true
			}
		}
	}
	return success, id
}

// Update Update
func (m *MyDB) Update(query string, args ...any) bool {
	var success = false
	var stmtUp *sql.Stmt
	stmtUp, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmtUp.Close()
		res, err := stmtUp.Exec(args...)
		if err != nil {
			log.Println("Update Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records updated")
			} else {
				success = true
			}
		}
	}
	return success
}

// Get Get
func (m *MyDB) Get(query string, args ...any) *DbRow {
	var rtn DbRow
	stmtGet, err := m.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		if err != nil {
			log.Println("Get err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					// err = rows.Scan(scanArgs...)
					// if err != nil {
					// 	log.Println("Error:", err.Error())
					// }
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rtn.Row = append(rtn.Row, value)
					}
				}
				// if err = rows.Err(); err != nil {
				// 	log.Println("Error:", err.Error())
				// }
			}
		}
	}
	return &rtn
}

// GetList GetList
func (m *MyDB) GetList(query string, args ...any) *DbRows {
	var rtn DbRows
	stmtGet, err := m.db.Prepare(query)
	if err != nil {
		log.Println("Error:", err.Error())
	} else {
		defer stmtGet.Close()
		rows, err := stmtGet.Query(args...)
		// defer rows.Close()
		if err != nil {
			log.Println("GetList err: ", err)
		} else {
			defer rows.Close()
			columns, err := rows.Columns()
			if err == nil {
				rtn.Columns = columns
				rowValues := make([]sql.RawBytes, len(columns))
				scanArgs := make([]any, len(rowValues))
				for i := range rowValues {
					scanArgs[i] = &rowValues[i]
				}
				for rows.Next() {
					var rowValuesStr []string
					// err = rows.Scan(scanArgs...)
					// if err != nil {
					// 	log.Println("Error:", err.Error())
					// }
					rows.Scan(scanArgs...)

					for _, col := range rowValues {
						var value string
						if col == nil {
							value = "NULL"
						} else {
							value = string(col)
						}
						rowValuesStr = append(rowValuesStr, value)
					}
					rtn.Rows = append(rtn.Rows, rowValuesStr)
				}
				// if err = rows.Err(); err != nil {
				// 	log.Println("Error:", err.Error())
				// }
			}
		}
	}
	return &rtn
}

// Delete Delete
func (m *MyDB) Delete(query string, args ...any) bool {
	var success = false
	var stmt *sql.Stmt
	stmt, m.err = m.db.Prepare(query)
	if m.err != nil {
		log.Println("Error:", m.err.Error())
	} else {
		defer stmt.Close()
		res, err := stmt.Exec(args...)
		if err != nil {
			log.Println("Delete Exec err:", err.Error())
		} else {
			affectedRows, _ := res.RowsAffected()
			if affectedRows == 0 {
				log.Println("Error: No records deleted")
			} else {
				success = true
			}
		}
	}
	return success
}

// Close Close
func (m *MyDB) Close() bool {
	var rtn = false
	err := m.db.Close()
	if err == nil {
		rtn = true
	}
	return rtn
}
