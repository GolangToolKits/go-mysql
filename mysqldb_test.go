package simplemysql

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMyDB_Connect(t *testing.T) {
	type fields struct {
		Host     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			fields: fields{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin2",
				Database: "testdb",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDB{
				Host:     tt.fields.Host,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			if got := m.Connect(); got != tt.want {
				t.Errorf("MyDB.Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDB_GetNewDatabase(t *testing.T) {
	var db MyDB

	type fields struct {
		Host     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   Database
	}{
		// TODO: Add test cases.
		{
			name:   "test 2",
			fields: fields{},
			want:   &db,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDB{
				Host:     tt.fields.Host,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			if got := m.GetNewDatabase(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyDB.GetNewDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDB_BeginTransaction(t *testing.T) {

	var tx MyDbTx
	type fields struct {
		Host     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	tests := []struct {
		name   string
		fields fields
		want   Transaction
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{},
			want:   &tx,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m.Connect()
			// if got := m.BeginTransaction(); !reflect.DeepEqual(got, tt.want) {
			if got := m.BeginTransaction(); got == nil {
				t.Errorf("MyDB.BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	var a []interface{}
	type fields struct {
		Host     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DbRow
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{},
			args: args{
				query: "select count(*) from test ",
				args:  a,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m.Connect()
			if got := m.Test(tt.args.query, tt.args.args...); len(got.Row) == 0 {
				t.Errorf("MyDB.Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDB_Insert(t *testing.T) {
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")

	var a2 []interface{}
	type fields struct {
		Host     string
		User     string
		Password string
		Database string
		db       *sql.DB
		err      error
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  int64
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "insert into test (name, address) values(?, ?)",
				args:  a,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				query: "insert into test (name, address) values(?, ?1)",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				query: "insert into test (name, address) values(?, ?)",
				args:  a2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m.Connect()
			got, got1 := m.Insert(tt.args.query, tt.args.args...)
			if got != tt.want {
				t.Errorf("MyDB.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 == 0 {
				t.Errorf("MyDB.Insert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
