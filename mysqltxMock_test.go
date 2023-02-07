package simplemysql

import (
	"database/sql"
	"testing"
)

func TestMyDbTxMock_Insert(t *testing.T) {
	type fields struct {
		Tx       *sql.Tx
		MyDBMock *MyDBMock
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
			fields: fields{
				MyDBMock: &MyDBMock{
					MockInsertSuccess1: true,
					MockInsertID1:      2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 2",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					MockInsertSuccess2:     true,
					MockInsertID2:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 3",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					MockInsertSuccess3:     true,
					MockInsertID3:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 4",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					mockInsertSuccess3Used: true,
					MockInsertSuccess4:     true,
					MockInsertID4:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 5",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					mockInsertSuccess3Used: true,
					mockInsertSuccess4Used: true,
					MockInsertSuccess5:     true,
					MockInsertID5:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 6",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					mockInsertSuccess3Used: true,
					mockInsertSuccess4Used: true,
					mockInsertSuccess5Used: true,
					MockInsertSuccess6:     true,
					MockInsertID6:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 7",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					mockInsertSuccess3Used: true,
					mockInsertSuccess4Used: true,
					mockInsertSuccess5Used: true,
					mockInsertSuccess6Used: true,
					MockInsertSuccess7:     true,
					MockInsertID7:          2,
				},
			},
			want:  true,
			want1: 2,
		},
		{
			name: "test 8",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockInsertSuccess1Used: true,
					mockInsertSuccess2Used: true,
					mockInsertSuccess3Used: true,
					mockInsertSuccess4Used: true,
					mockInsertSuccess5Used: true,
					mockInsertSuccess6Used: true,
					mockInsertSuccess7Used: true,
					MockInsertSuccess8:     true,
					MockInsertID8:          2,
				},
			},
			want:  true,
			want1: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &MyDbTxMock{
				Tx:       tt.fields.Tx,
				MyDBMock: tt.fields.MyDBMock,
			}
			got, got1 := tr.Insert(tt.args.query, tt.args.args...)
			if got != tt.want {
				t.Errorf("MyDbTxMock.Insert() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MyDbTxMock.Insert() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMyDbTxMock_Update(t *testing.T) {
	type fields struct {
		Tx       *sql.Tx
		MyDBMock *MyDBMock
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
			fields: fields{
				MyDBMock: &MyDBMock{
					MockUpdateSuccess1: true,
				},
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockUpdateSuccess1Used: true,
					MockUpdateSuccess2:     true,
				},
			},
			want: true,
		},
		{
			name: "test 3",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockUpdateSuccess1Used: true,
					mockUpdateSuccess2Used: true,
					MockUpdateSuccess3:     true,
				},
			},
			want: true,
		},
		{
			name: "test 4",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockUpdateSuccess1Used: true,
					mockUpdateSuccess2Used: true,
					mockUpdateSuccess3Used: true,
					MockUpdateSuccess4:     true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &MyDbTxMock{
				Tx:       tt.fields.Tx,
				MyDBMock: tt.fields.MyDBMock,
			}
			if got := tr.Update(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDbTxMock.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDbTxMock_Delete(t *testing.T) {
	type fields struct {
		Tx       *sql.Tx
		MyDBMock *MyDBMock
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
			fields: fields{
				MyDBMock: &MyDBMock{
					MockDeleteSuccess1: true,
				},
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					MockDeleteSuccess2:     true,
				},
			},
			want: true,
		},
		{
			name: "test 3",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					MockDeleteSuccess3:     true,
				},
			},
			want: true,
		},
		{
			name: "test 4",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					mockDeleteSuccess3Used: true,
					MockDeleteSuccess4:     true,
				},
			},
			want: true,
		},
		{
			name: "test 5",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					mockDeleteSuccess3Used: true,
					mockDeleteSuccess4Used: true,
					MockDeleteSuccess5:     true,
				},
			},
			want: true,
		},
		{
			name: "test 6",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					mockDeleteSuccess3Used: true,
					mockDeleteSuccess4Used: true,
					mockDeleteSuccess5Used: true,
					MockDeleteSuccess6:     true,
				},
			},
			want: true,
		},
		{
			name: "test 7",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					mockDeleteSuccess3Used: true,
					mockDeleteSuccess4Used: true,
					mockDeleteSuccess5Used: true,
					mockDeleteSuccess6Used: true,
					MockDeleteSuccess7:     true,
				},
			},
			want: true,
		},
		{
			name: "test 8",
			fields: fields{
				MyDBMock: &MyDBMock{
					mockDeleteSuccess1Used: true,
					mockDeleteSuccess2Used: true,
					mockDeleteSuccess3Used: true,
					mockDeleteSuccess4Used: true,
					mockDeleteSuccess5Used: true,
					mockDeleteSuccess6Used: true,
					mockDeleteSuccess7Used: true,
					MockDeleteSuccess8:     true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &MyDbTxMock{
				Tx:       tt.fields.Tx,
				MyDBMock: tt.fields.MyDBMock,
			}
			if got := tr.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDbTxMock.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDbTxMock_Commit(t *testing.T) {
	type fields struct {
		Tx       *sql.Tx
		MyDBMock *MyDBMock
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
				MyDBMock: &MyDBMock{
					MockCommitSuccess: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &MyDbTxMock{
				Tx:       tt.fields.Tx,
				MyDBMock: tt.fields.MyDBMock,
			}
			if got := tr.Commit(); got != tt.want {
				t.Errorf("MyDbTxMock.Commit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDbTxMock_Rollback(t *testing.T) {
	type fields struct {
		Tx       *sql.Tx
		MyDBMock *MyDBMock
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
				MyDBMock:&MyDBMock{
					MockRollbackSuccess: true,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &MyDbTxMock{
				Tx:       tt.fields.Tx,
				MyDBMock: tt.fields.MyDBMock,
			}
			if got := tr.Rollback(); got != tt.want {
				t.Errorf("MyDbTxMock.Rollback() = %v, want %v", got, tt.want)
			}
		})
	}
}
