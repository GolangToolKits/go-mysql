package gomysql

import (
	"database/sql"
	"reflect"
	"testing"
	"time"

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
			mm := &MyDB{
				Host:     tt.fields.Host,
				User:     tt.fields.User,
				Password: tt.fields.Password,
				Database: tt.fields.Database,
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
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
			if got := m.New(); !reflect.DeepEqual(got, tt.want) {
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
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			// if got := m.BeginTransaction(); !reflect.DeepEqual(got, tt.want) {
			if got := m.BeginTransaction(); got == nil {
				t.Errorf("MyDB.BeginTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test(t *testing.T) {
	var a []any
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
		args  []any
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
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			if got := m.Test(tt.args.query, tt.args.args...); len(got.Row) == 0 {
				t.Errorf("MyDB.Test() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDB_Insert(t *testing.T) {
	var a []any
	a = append(a, "test insert 1", "123 main st")

	var a2 []any
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
		args  []any
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
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
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

func TestMyDB_Update(t *testing.T) {
	var id = 112
	var a []any
	var tstr = time.Now().UnixNano()
	a = append(a, tstr, "11111 main st", id)

	var a2 []any

	var a3 []any
	a3 = append(a3, "test insert 2", "123456 main st", 0)

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
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "update test set name = ? , address = ? where id = ? ",
				args:  a,
			},
			want: true,
		},
		{
			name: "test 2",
			args: args{
				query: "update test set name = ? , address = ? where id = ?? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				query: "update test set name = ? , address = ? where id = ? ",
				args:  a2,
			},
			want: false,
		},
		{
			name: "test 4",
			args: args{
				query: "update test set name = ? , address = ? where id = ? ",
				args:  a3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			if got := m.Update(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDB.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDB_Get(t *testing.T) {
	var id = 112
	var a []any
	a = append(a, id)

	var a2 []any

	var id3 = 172
	var a3 []any
	a3 = append(a3, id3)

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
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DbRow
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "select * from test where id = ? ",
				args:  a,
			},
			//want: true,
		},
		{
			name: "test 2",
			args: args{
				query: "select * from test where id = ?? ",
				args:  a,
			},
			//want: true,
		},
		{
			name: "test 3",
			args: args{
				query: "select * from test where id = ? ",
				args:  a2,
			},
			//want: true,
		},
		{
			name: "test 4",
			args: args{
				query: "select * from test where id = ? ",
				args:  a3,
			},
			//want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			// if got := m.Get(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
			got := m.Get(tt.args.query, tt.args.args...)
			{
				//t.Errorf("MyDB.Get() = %v, want %v", got, tt.want)
				if tt.name == "test 1" && len(got.Row) == 0 {
					t.Fail()
				}
				if tt.name == "test 2" && got.Row != nil {
					t.Fail()
				}
				if tt.name == "test 3" && len(got.Row) != 0 {
					t.Fail()
				}
			}
		})
	}
}

func TestMyDB_GetList(t *testing.T) {
	var a []any
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
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *DbRows
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "select * from test ",
				args:  a,
			},
			//want: true,
		},
		{
			name: "test 2",
			args: args{
				query: "select * from test ?",
				args:  a,
			},
			//want: true,
		},
		{
			name: "test 3",
			args: args{
				query: "select * from test where name = ?",
				args:  a,
			},
			//want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			// if got := m.GetList(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
			got := m.GetList(tt.args.query, tt.args.args...)
			if tt.name == "test 1" && len(got.Rows) == 0 {
				t.Fail()
			}

			if tt.name == "test 2" && len(got.Rows) != 0 {
				t.Fail()
			}

			if tt.name == "test 3" && len(got.Rows) != 0 {
				t.Fail()
			}
			//t.Errorf("MyDB.GetList() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestMyDB_Delete(t *testing.T) {
	var id = 130
	var a []any
	a = append(a, id)

	var a2 []any

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
		args  []any
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "delete from test where id = ? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 2",
			args: args{
				query: "delete from test where id = ?? ",
				args:  a,
			},
			want: false,
		},
		{
			name: "test 3",
			args: args{
				query: "delete from test where id = ? ",
				args:  a2,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			// if got := m.Delete(tt.args.query, tt.args.args...); got != tt.want {
			// 	t.Errorf("MyDB.Delete() = %v, want %v", got, tt.want)
			// }
			got := m.Delete(tt.args.query, tt.args.args...)
			if tt.name == "test 1" && tt.want != got {
				t.Fail()
			}

			if tt.name == "test 2" && tt.want != got {
				t.Fail()
			}

			if tt.name == "test 3" && tt.want != got {
				t.Fail()
			}

		})
	}
}

func TestMyDB_Close(t *testing.T) {
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
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				db:       tt.fields.db,
				err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			if got := m.Close(); got != tt.want {
				t.Errorf("MyDB.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
