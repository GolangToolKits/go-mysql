package simplemysql

import (
	"database/sql"
	"testing"
	"time"
)

func TestMyDbTx_Insert(t *testing.T) {
	var a []interface{}
	a = append(a, "test insert 22", "123 main st")

	var a2 []interface{}

	type fields struct {
		Tx *sql.Tx
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
			want:  true,
			want1: 0,
		},
		{
			name: "test 2",
			args: args{
				query: "insert into test (name, address) values(?, ??)",
				args:  a,
			},
			want:  false,
			want1: 1,
		},
		{
			name: "test 3",
			args: args{
				query: "insert into test (name, address) values(?, ?)",
				args:  a2,
			},
			want:  false,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			got, got1 := tr.Insert(tt.args.query, tt.args.args...)
			tr.Commit()
			if got != tt.want {
				t.Errorf("MyDbTx.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 == tt.want1 {
				t.Errorf("MyDbTx.Insert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMyDbTx_Update(t *testing.T) {
	var id = 112
	var a []interface{}
	var tstr = time.Now().UnixNano()
	a = append(a, tstr, "22222 main st", id)

	var a2 []interface{}

	type fields struct {
		Tx *sql.Tx
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
				query: "update test set name = ? , address = ? where id = ?/ ",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			if got := tr.Update(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDbTx.Update() = %v, want %v", got, tt.want)
			}
			tr.Commit()
		})
	}
}

func TestMyDbTx_Delete(t *testing.T) {

	var id = 140
	var a []interface{}
	a = append(a, id)

	var a2 []interface{}

	type fields struct {
		Tx *sql.Tx
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
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "delete from test where id = ? ",
				args:  a,
			},
			want: true,
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
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			if got := tr.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDbTx.Delete() = %v, want %v", got, tt.want)
			}
			tr.Commit()
		})
	}
}

func TestMyDbTx_Commit(t *testing.T) {
	type fields struct {
		Tx *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			tr.Commit()
			if got := tr.Commit(); got != tt.want {
				t.Errorf("MyDbTx.Commit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDbTx_Delete_Rollback(t *testing.T) {

	var id = 141
	var a []interface{}
	a = append(a, id)

	//var a2 []interface{}

	type fields struct {
		Tx *sql.Tx
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
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			args: args{
				query: "delete from test where id = ? ",
				args:  a,
			},
			want: true,
		},
		// {
		// 	name: "test 2",
		// 	args: args{
		// 		query: "delete from test where id = ?? ",
		// 		args:  a,
		// 	},
		// 	want: false,
		// },
		// {
		// 	name: "test 3",
		// 	args: args{
		// 		query: "delete from test where id = ? ",
		// 		args:  a2,
		// 	},
		// 	want: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			if got := tr.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDbTx.Delete() = %v, want %v", got, tt.want)
			}
			tr.Rollback()
		})
	}
}
func TestMyDbTx_Rollback(t *testing.T) {
	type fields struct {
		Tx *sql.Tx
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// tr := &MyDbTx{
			// 	Tx: tt.fields.Tx,
			// }
			mm := &MyDB{
				Host:     "localhost:3306",
				User:     "admin",
				Password: "admin",
				Database: "testdb",
				// db:       tt.fields.db,
				// err:      tt.fields.err,
			}
			m := mm.New()
			m.Connect()
			tr := m.BeginTransaction()
			tr.Rollback()
			if got := tr.Rollback(); got != tt.want {
				t.Errorf("MyDbTx.Rollback() = %v, want %v", got, tt.want)
			}
		})
	}
}
