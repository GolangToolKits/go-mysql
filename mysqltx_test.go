package simplemysql

import (
	"database/sql"
	"testing"
)

func TestMyDbTx_Insert(t *testing.T) {
	var a []interface{}
	a = append(a, "test insert 1", "123 main st")


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
			want: true,
			want1: 0,
		},
		{
			name: "test 2",
			args: args{
				query: "insert into test (name, address) values(?, ??)",
				args:  a,
			},
			want: false,
			want1: 1,
		},
		{
			name: "test 3",
			args: args{
				query: "insert into test (name, address) values(?, ?)",
				args:  a2,
			},
			want: false,
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
			tr:= m.BeginTransaction()
			got, got1 := tr.Insert(tt.args.query, tt.args.args...)
			if got != tt.want {
				t.Errorf("MyDbTx.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 == tt.want1 {
				t.Errorf("MyDbTx.Insert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
