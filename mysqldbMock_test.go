package gomysql

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestMyDBMock_Get(t *testing.T) {
	var mr DbRow
	type fields struct {
		Host                   string
		User                   string
		Password               string
		Database               string
		db                     *sql.DB
		err                    error
		MockConnectSuccess     bool
		MockCloseSuccess       bool
		MockCommitSuccess      bool
		MockRollbackSuccess    bool
		MockInsertSuccess1     bool
		MockInsertSuccess2     bool
		MockInsertSuccess3     bool
		MockInsertSuccess4     bool
		MockInsertSuccess5     bool
		MockInsertSuccess6     bool
		MockInsertSuccess7     bool
		MockInsertSuccess8     bool
		mockInsertSuccess1Used bool
		mockInsertSuccess2Used bool
		mockInsertSuccess3Used bool
		mockInsertSuccess4Used bool
		mockInsertSuccess5Used bool
		mockInsertSuccess6Used bool
		mockInsertSuccess7Used bool
		mockInsertSuccess8Used bool
		MockInsertID1          int64
		MockInsertID2          int64
		MockInsertID3          int64
		MockInsertID4          int64
		MockInsertID5          int64
		MockInsertID6          int64
		MockInsertID7          int64
		MockInsertID8          int64
		MockUpdateSuccess1     bool
		MockUpdateSuccess2     bool
		MockUpdateSuccess3     bool
		MockUpdateSuccess4     bool
		mockUpdateSuccess1Used bool
		mockUpdateSuccess2Used bool
		mockUpdateSuccess3Used bool
		mockUpdateSuccess4Used bool
		MockDeleteSuccess1     bool
		MockDeleteSuccess2     bool
		MockDeleteSuccess3     bool
		MockDeleteSuccess4     bool
		MockDeleteSuccess5     bool
		MockDeleteSuccess6     bool
		MockDeleteSuccess7     bool
		MockDeleteSuccess8     bool
		mockDeleteSuccess1Used bool
		mockDeleteSuccess2Used bool
		mockDeleteSuccess3Used bool
		mockDeleteSuccess4Used bool
		mockDeleteSuccess5Used bool
		mockDeleteSuccess6Used bool
		mockDeleteSuccess7Used bool
		mockDeleteSuccess8Used bool
		MockTestRow            *DbRow
		mockRow1Used           bool
		mockRow2Used           bool
		mockRow3Used           bool
		mockRow4Used           bool
		mockRow5Used           bool
		mockRow6Used           bool
		mockRow7Used           bool
		mockRow8Used           bool
		MockRow1               *DbRow
		MockRow2               *DbRow
		MockRow3               *DbRow
		MockRow4               *DbRow
		MockRow5               *DbRow
		MockRow6               *DbRow
		MockRow7               *DbRow
		MockRow8               *DbRow
		mockRows1Used          bool
		mockRows2Used          bool
		mockRows3Used          bool
		mockRows4Used          bool
		mockRows5Used          bool
		mockRows6Used          bool
		mockRows7Used          bool
		mockRows8Used          bool
		MockRows1              *DbRows
		MockRows2              *DbRows
		MockRows3              *DbRows
		MockRows4              *DbRows
		MockRows5              *DbRows
		MockRows6              *DbRows
		MockRows7              *DbRows
		MockRows8              *DbRows
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
			fields: fields{
				MockRow1: &mr,
			},
			want: &mr,
		},
		{
			name: "test 2",
			fields: fields{
				mockRow1Used: true,
				MockRow2:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 3",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				MockRow3:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 4",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				mockRow3Used: true,
				MockRow4:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 5",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				mockRow3Used: true,
				mockRow4Used: true,
				MockRow5:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 6",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				mockRow3Used: true,
				mockRow4Used: true,
				mockRow5Used: true,
				MockRow6:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 7",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				mockRow3Used: true,
				mockRow4Used: true,
				mockRow5Used: true,
				mockRow6Used: true,
				MockRow7:     &mr,
			},
			want: &mr,
		},
		{
			name: "test 8",
			fields: fields{
				mockRow1Used: true,
				mockRow2Used: true,
				mockRow3Used: true,
				mockRow4Used: true,
				mockRow5Used: true,
				mockRow6Used: true,
				mockRow7Used: true,
				MockRow8:     &mr,
			},
			want: &mr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mm := &MyDBMock{
				Host:                   tt.fields.Host,
				User:                   tt.fields.User,
				Password:               tt.fields.Password,
				Database:               tt.fields.Database,
				db:                     tt.fields.db,
				err:                    tt.fields.err,
				MockConnectSuccess:     tt.fields.MockConnectSuccess,
				MockCloseSuccess:       tt.fields.MockCloseSuccess,
				MockCommitSuccess:      tt.fields.MockCommitSuccess,
				MockRollbackSuccess:    tt.fields.MockRollbackSuccess,
				MockInsertSuccess1:     tt.fields.MockInsertSuccess1,
				MockInsertSuccess2:     tt.fields.MockInsertSuccess2,
				MockInsertSuccess3:     tt.fields.MockInsertSuccess3,
				MockInsertSuccess4:     tt.fields.MockInsertSuccess4,
				MockInsertSuccess5:     tt.fields.MockInsertSuccess5,
				MockInsertSuccess6:     tt.fields.MockInsertSuccess6,
				MockInsertSuccess7:     tt.fields.MockInsertSuccess7,
				MockInsertSuccess8:     tt.fields.MockInsertSuccess8,
				mockInsertSuccess1Used: tt.fields.mockInsertSuccess1Used,
				mockInsertSuccess2Used: tt.fields.mockInsertSuccess2Used,
				mockInsertSuccess3Used: tt.fields.mockInsertSuccess3Used,
				mockInsertSuccess4Used: tt.fields.mockInsertSuccess4Used,
				mockInsertSuccess5Used: tt.fields.mockInsertSuccess5Used,
				mockInsertSuccess6Used: tt.fields.mockInsertSuccess6Used,
				mockInsertSuccess7Used: tt.fields.mockInsertSuccess7Used,
				mockInsertSuccess8Used: tt.fields.mockInsertSuccess8Used,
				MockInsertID1:          tt.fields.MockInsertID1,
				MockInsertID2:          tt.fields.MockInsertID2,
				MockInsertID3:          tt.fields.MockInsertID3,
				MockInsertID4:          tt.fields.MockInsertID4,
				MockInsertID5:          tt.fields.MockInsertID5,
				MockInsertID6:          tt.fields.MockInsertID6,
				MockInsertID7:          tt.fields.MockInsertID7,
				MockInsertID8:          tt.fields.MockInsertID8,
				MockUpdateSuccess1:     tt.fields.MockUpdateSuccess1,
				MockUpdateSuccess2:     tt.fields.MockUpdateSuccess2,
				MockUpdateSuccess3:     tt.fields.MockUpdateSuccess3,
				MockUpdateSuccess4:     tt.fields.MockUpdateSuccess4,
				mockUpdateSuccess1Used: tt.fields.mockUpdateSuccess1Used,
				mockUpdateSuccess2Used: tt.fields.mockUpdateSuccess2Used,
				mockUpdateSuccess3Used: tt.fields.mockUpdateSuccess3Used,
				mockUpdateSuccess4Used: tt.fields.mockUpdateSuccess4Used,
				MockDeleteSuccess1:     tt.fields.MockDeleteSuccess1,
				MockDeleteSuccess2:     tt.fields.MockDeleteSuccess2,
				MockDeleteSuccess3:     tt.fields.MockDeleteSuccess3,
				MockDeleteSuccess4:     tt.fields.MockDeleteSuccess4,
				MockDeleteSuccess5:     tt.fields.MockDeleteSuccess5,
				MockDeleteSuccess6:     tt.fields.MockDeleteSuccess6,
				MockDeleteSuccess7:     tt.fields.MockDeleteSuccess7,
				MockDeleteSuccess8:     tt.fields.MockDeleteSuccess8,
				mockDeleteSuccess1Used: tt.fields.mockDeleteSuccess1Used,
				mockDeleteSuccess2Used: tt.fields.mockDeleteSuccess2Used,
				mockDeleteSuccess3Used: tt.fields.mockDeleteSuccess3Used,
				mockDeleteSuccess4Used: tt.fields.mockDeleteSuccess4Used,
				mockDeleteSuccess5Used: tt.fields.mockDeleteSuccess5Used,
				mockDeleteSuccess6Used: tt.fields.mockDeleteSuccess6Used,
				mockDeleteSuccess7Used: tt.fields.mockDeleteSuccess7Used,
				mockDeleteSuccess8Used: tt.fields.mockDeleteSuccess8Used,
				MockTestRow:            tt.fields.MockTestRow,
				mockRow1Used:           tt.fields.mockRow1Used,
				mockRow2Used:           tt.fields.mockRow2Used,
				mockRow3Used:           tt.fields.mockRow3Used,
				mockRow4Used:           tt.fields.mockRow4Used,
				mockRow5Used:           tt.fields.mockRow5Used,
				mockRow6Used:           tt.fields.mockRow6Used,
				mockRow7Used:           tt.fields.mockRow7Used,
				mockRow8Used:           tt.fields.mockRow8Used,
				MockRow1:               tt.fields.MockRow1,
				MockRow2:               tt.fields.MockRow2,
				MockRow3:               tt.fields.MockRow3,
				MockRow4:               tt.fields.MockRow4,
				MockRow5:               tt.fields.MockRow5,
				MockRow6:               tt.fields.MockRow6,
				MockRow7:               tt.fields.MockRow7,
				MockRow8:               tt.fields.MockRow8,
				mockRows1Used:          tt.fields.mockRows1Used,
				mockRows2Used:          tt.fields.mockRows2Used,
				mockRows3Used:          tt.fields.mockRows3Used,
				mockRows4Used:          tt.fields.mockRows4Used,
				mockRows5Used:          tt.fields.mockRows5Used,
				mockRows6Used:          tt.fields.mockRows6Used,
				mockRows7Used:          tt.fields.mockRows7Used,
				mockRows8Used:          tt.fields.mockRows8Used,
				MockRows1:              tt.fields.MockRows1,
				MockRows2:              tt.fields.MockRows2,
				MockRows3:              tt.fields.MockRows3,
				MockRows4:              tt.fields.MockRows4,
				MockRows5:              tt.fields.MockRows5,
				MockRows6:              tt.fields.MockRows6,
				MockRows7:              tt.fields.MockRows7,
				MockRows8:              tt.fields.MockRows8,
			}
			m := mm.New()
			if got := m.Get(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyDBMock.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDBMock_GetList(t *testing.T) {
	var mrs DbRows
	type fields struct {
		Host                   string
		User                   string
		Password               string
		Database               string
		db                     *sql.DB
		err                    error
		MockConnectSuccess     bool
		MockCloseSuccess       bool
		MockCommitSuccess      bool
		MockRollbackSuccess    bool
		MockInsertSuccess1     bool
		MockInsertSuccess2     bool
		MockInsertSuccess3     bool
		MockInsertSuccess4     bool
		MockInsertSuccess5     bool
		MockInsertSuccess6     bool
		MockInsertSuccess7     bool
		MockInsertSuccess8     bool
		mockInsertSuccess1Used bool
		mockInsertSuccess2Used bool
		mockInsertSuccess3Used bool
		mockInsertSuccess4Used bool
		mockInsertSuccess5Used bool
		mockInsertSuccess6Used bool
		mockInsertSuccess7Used bool
		mockInsertSuccess8Used bool
		MockInsertID1          int64
		MockInsertID2          int64
		MockInsertID3          int64
		MockInsertID4          int64
		MockInsertID5          int64
		MockInsertID6          int64
		MockInsertID7          int64
		MockInsertID8          int64
		MockUpdateSuccess1     bool
		MockUpdateSuccess2     bool
		MockUpdateSuccess3     bool
		MockUpdateSuccess4     bool
		mockUpdateSuccess1Used bool
		mockUpdateSuccess2Used bool
		mockUpdateSuccess3Used bool
		mockUpdateSuccess4Used bool
		MockDeleteSuccess1     bool
		MockDeleteSuccess2     bool
		MockDeleteSuccess3     bool
		MockDeleteSuccess4     bool
		MockDeleteSuccess5     bool
		MockDeleteSuccess6     bool
		MockDeleteSuccess7     bool
		MockDeleteSuccess8     bool
		mockDeleteSuccess1Used bool
		mockDeleteSuccess2Used bool
		mockDeleteSuccess3Used bool
		mockDeleteSuccess4Used bool
		mockDeleteSuccess5Used bool
		mockDeleteSuccess6Used bool
		mockDeleteSuccess7Used bool
		mockDeleteSuccess8Used bool
		MockTestRow            *DbRow
		mockRow1Used           bool
		mockRow2Used           bool
		mockRow3Used           bool
		mockRow4Used           bool
		mockRow5Used           bool
		mockRow6Used           bool
		mockRow7Used           bool
		mockRow8Used           bool
		MockRow1               *DbRow
		MockRow2               *DbRow
		MockRow3               *DbRow
		MockRow4               *DbRow
		MockRow5               *DbRow
		MockRow6               *DbRow
		MockRow7               *DbRow
		MockRow8               *DbRow
		mockRows1Used          bool
		mockRows2Used          bool
		mockRows3Used          bool
		mockRows4Used          bool
		mockRows5Used          bool
		mockRows6Used          bool
		mockRows7Used          bool
		mockRows8Used          bool
		MockRows1              *DbRows
		MockRows2              *DbRows
		MockRows3              *DbRows
		MockRows4              *DbRows
		MockRows5              *DbRows
		MockRows6              *DbRows
		MockRows7              *DbRows
		MockRows8              *DbRows
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
			fields: fields{
				MockRows1: &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 2",
			fields: fields{
				mockRows1Used: true,
				MockRows2:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 3",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				MockRows3:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 4",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				mockRows3Used: true,
				MockRows4:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 5",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				mockRows3Used: true,
				mockRows4Used: true,
				MockRows5:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 6",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				mockRows3Used: true,
				mockRows4Used: true,
				mockRows5Used: true,
				MockRows6:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 7",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				mockRows3Used: true,
				mockRows4Used: true,
				mockRows5Used: true,
				mockRows6Used: true,
				MockRows7:     &mrs,
			},
			want: &mrs,
		},
		{
			name: "test 8",
			fields: fields{
				mockRows1Used: true,
				mockRows2Used: true,
				mockRows3Used: true,
				mockRows4Used: true,
				mockRows5Used: true,
				mockRows6Used: true,
				mockRows7Used: true,
				MockRows8:     &mrs,
			},
			want: &mrs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDBMock{
				Host:                   tt.fields.Host,
				User:                   tt.fields.User,
				Password:               tt.fields.Password,
				Database:               tt.fields.Database,
				db:                     tt.fields.db,
				err:                    tt.fields.err,
				MockConnectSuccess:     tt.fields.MockConnectSuccess,
				MockCloseSuccess:       tt.fields.MockCloseSuccess,
				MockCommitSuccess:      tt.fields.MockCommitSuccess,
				MockRollbackSuccess:    tt.fields.MockRollbackSuccess,
				MockInsertSuccess1:     tt.fields.MockInsertSuccess1,
				MockInsertSuccess2:     tt.fields.MockInsertSuccess2,
				MockInsertSuccess3:     tt.fields.MockInsertSuccess3,
				MockInsertSuccess4:     tt.fields.MockInsertSuccess4,
				MockInsertSuccess5:     tt.fields.MockInsertSuccess5,
				MockInsertSuccess6:     tt.fields.MockInsertSuccess6,
				MockInsertSuccess7:     tt.fields.MockInsertSuccess7,
				MockInsertSuccess8:     tt.fields.MockInsertSuccess8,
				mockInsertSuccess1Used: tt.fields.mockInsertSuccess1Used,
				mockInsertSuccess2Used: tt.fields.mockInsertSuccess2Used,
				mockInsertSuccess3Used: tt.fields.mockInsertSuccess3Used,
				mockInsertSuccess4Used: tt.fields.mockInsertSuccess4Used,
				mockInsertSuccess5Used: tt.fields.mockInsertSuccess5Used,
				mockInsertSuccess6Used: tt.fields.mockInsertSuccess6Used,
				mockInsertSuccess7Used: tt.fields.mockInsertSuccess7Used,
				mockInsertSuccess8Used: tt.fields.mockInsertSuccess8Used,
				MockInsertID1:          tt.fields.MockInsertID1,
				MockInsertID2:          tt.fields.MockInsertID2,
				MockInsertID3:          tt.fields.MockInsertID3,
				MockInsertID4:          tt.fields.MockInsertID4,
				MockInsertID5:          tt.fields.MockInsertID5,
				MockInsertID6:          tt.fields.MockInsertID6,
				MockInsertID7:          tt.fields.MockInsertID7,
				MockInsertID8:          tt.fields.MockInsertID8,
				MockUpdateSuccess1:     tt.fields.MockUpdateSuccess1,
				MockUpdateSuccess2:     tt.fields.MockUpdateSuccess2,
				MockUpdateSuccess3:     tt.fields.MockUpdateSuccess3,
				MockUpdateSuccess4:     tt.fields.MockUpdateSuccess4,
				mockUpdateSuccess1Used: tt.fields.mockUpdateSuccess1Used,
				mockUpdateSuccess2Used: tt.fields.mockUpdateSuccess2Used,
				mockUpdateSuccess3Used: tt.fields.mockUpdateSuccess3Used,
				mockUpdateSuccess4Used: tt.fields.mockUpdateSuccess4Used,
				MockDeleteSuccess1:     tt.fields.MockDeleteSuccess1,
				MockDeleteSuccess2:     tt.fields.MockDeleteSuccess2,
				MockDeleteSuccess3:     tt.fields.MockDeleteSuccess3,
				MockDeleteSuccess4:     tt.fields.MockDeleteSuccess4,
				MockDeleteSuccess5:     tt.fields.MockDeleteSuccess5,
				MockDeleteSuccess6:     tt.fields.MockDeleteSuccess6,
				MockDeleteSuccess7:     tt.fields.MockDeleteSuccess7,
				MockDeleteSuccess8:     tt.fields.MockDeleteSuccess8,
				mockDeleteSuccess1Used: tt.fields.mockDeleteSuccess1Used,
				mockDeleteSuccess2Used: tt.fields.mockDeleteSuccess2Used,
				mockDeleteSuccess3Used: tt.fields.mockDeleteSuccess3Used,
				mockDeleteSuccess4Used: tt.fields.mockDeleteSuccess4Used,
				mockDeleteSuccess5Used: tt.fields.mockDeleteSuccess5Used,
				mockDeleteSuccess6Used: tt.fields.mockDeleteSuccess6Used,
				mockDeleteSuccess7Used: tt.fields.mockDeleteSuccess7Used,
				mockDeleteSuccess8Used: tt.fields.mockDeleteSuccess8Used,
				MockTestRow:            tt.fields.MockTestRow,
				mockRow1Used:           tt.fields.mockRow1Used,
				mockRow2Used:           tt.fields.mockRow2Used,
				mockRow3Used:           tt.fields.mockRow3Used,
				mockRow4Used:           tt.fields.mockRow4Used,
				mockRow5Used:           tt.fields.mockRow5Used,
				mockRow6Used:           tt.fields.mockRow6Used,
				mockRow7Used:           tt.fields.mockRow7Used,
				mockRow8Used:           tt.fields.mockRow8Used,
				MockRow1:               tt.fields.MockRow1,
				MockRow2:               tt.fields.MockRow2,
				MockRow3:               tt.fields.MockRow3,
				MockRow4:               tt.fields.MockRow4,
				MockRow5:               tt.fields.MockRow5,
				MockRow6:               tt.fields.MockRow6,
				MockRow7:               tt.fields.MockRow7,
				MockRow8:               tt.fields.MockRow8,
				mockRows1Used:          tt.fields.mockRows1Used,
				mockRows2Used:          tt.fields.mockRows2Used,
				mockRows3Used:          tt.fields.mockRows3Used,
				mockRows4Used:          tt.fields.mockRows4Used,
				mockRows5Used:          tt.fields.mockRows5Used,
				mockRows6Used:          tt.fields.mockRows6Used,
				mockRows7Used:          tt.fields.mockRows7Used,
				mockRows8Used:          tt.fields.mockRows8Used,
				MockRows1:              tt.fields.MockRows1,
				MockRows2:              tt.fields.MockRows2,
				MockRows3:              tt.fields.MockRows3,
				MockRows4:              tt.fields.MockRows4,
				MockRows5:              tt.fields.MockRows5,
				MockRows6:              tt.fields.MockRows6,
				MockRows7:              tt.fields.MockRows7,
				MockRows8:              tt.fields.MockRows8,
			}
			if got := m.GetList(tt.args.query, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MyDBMock.GetList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDBMock_Delete(t *testing.T) {
	type fields struct {
		Host                   string
		User                   string
		Password               string
		Database               string
		db                     *sql.DB
		err                    error
		MockConnectSuccess     bool
		MockCloseSuccess       bool
		MockCommitSuccess      bool
		MockRollbackSuccess    bool
		MockInsertSuccess1     bool
		MockInsertSuccess2     bool
		MockInsertSuccess3     bool
		MockInsertSuccess4     bool
		MockInsertSuccess5     bool
		MockInsertSuccess6     bool
		MockInsertSuccess7     bool
		MockInsertSuccess8     bool
		mockInsertSuccess1Used bool
		mockInsertSuccess2Used bool
		mockInsertSuccess3Used bool
		mockInsertSuccess4Used bool
		mockInsertSuccess5Used bool
		mockInsertSuccess6Used bool
		mockInsertSuccess7Used bool
		mockInsertSuccess8Used bool
		MockInsertID1          int64
		MockInsertID2          int64
		MockInsertID3          int64
		MockInsertID4          int64
		MockInsertID5          int64
		MockInsertID6          int64
		MockInsertID7          int64
		MockInsertID8          int64
		MockUpdateSuccess1     bool
		MockUpdateSuccess2     bool
		MockUpdateSuccess3     bool
		MockUpdateSuccess4     bool
		mockUpdateSuccess1Used bool
		mockUpdateSuccess2Used bool
		mockUpdateSuccess3Used bool
		mockUpdateSuccess4Used bool
		MockDeleteSuccess1     bool
		MockDeleteSuccess2     bool
		MockDeleteSuccess3     bool
		MockDeleteSuccess4     bool
		MockDeleteSuccess5     bool
		MockDeleteSuccess6     bool
		MockDeleteSuccess7     bool
		MockDeleteSuccess8     bool
		mockDeleteSuccess1Used bool
		mockDeleteSuccess2Used bool
		mockDeleteSuccess3Used bool
		mockDeleteSuccess4Used bool
		mockDeleteSuccess5Used bool
		mockDeleteSuccess6Used bool
		mockDeleteSuccess7Used bool
		mockDeleteSuccess8Used bool
		MockTestRow            *DbRow
		mockRow1Used           bool
		mockRow2Used           bool
		mockRow3Used           bool
		mockRow4Used           bool
		mockRow5Used           bool
		mockRow6Used           bool
		mockRow7Used           bool
		mockRow8Used           bool
		MockRow1               *DbRow
		MockRow2               *DbRow
		MockRow3               *DbRow
		MockRow4               *DbRow
		MockRow5               *DbRow
		MockRow6               *DbRow
		MockRow7               *DbRow
		MockRow8               *DbRow
		mockRows1Used          bool
		mockRows2Used          bool
		mockRows3Used          bool
		mockRows4Used          bool
		mockRows5Used          bool
		mockRows6Used          bool
		mockRows7Used          bool
		mockRows8Used          bool
		MockRows1              *DbRows
		MockRows2              *DbRows
		MockRows3              *DbRows
		MockRows4              *DbRows
		MockRows5              *DbRows
		MockRows6              *DbRows
		MockRows7              *DbRows
		MockRows8              *DbRows
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
				MockDeleteSuccess1: true,
			},
			want: true,
		},
		{
			name: "test 2",
			fields: fields{
				mockDeleteSuccess1Used: true,
				MockDeleteSuccess2:     true,
			},
			want: true,
		},
		{
			name: "test 3",
			fields: fields{
				mockDeleteSuccess1Used: true,
				mockDeleteSuccess2Used: true,
				MockDeleteSuccess3:     true,
			},
			want: true,
		},
		{
			name: "test 4",
			fields: fields{
				mockDeleteSuccess1Used: true,
				mockDeleteSuccess2Used: true,
				mockDeleteSuccess3Used: true,
				MockDeleteSuccess4:     true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDBMock{
				Host:                   tt.fields.Host,
				User:                   tt.fields.User,
				Password:               tt.fields.Password,
				Database:               tt.fields.Database,
				db:                     tt.fields.db,
				err:                    tt.fields.err,
				MockConnectSuccess:     tt.fields.MockConnectSuccess,
				MockCloseSuccess:       tt.fields.MockCloseSuccess,
				MockCommitSuccess:      tt.fields.MockCommitSuccess,
				MockRollbackSuccess:    tt.fields.MockRollbackSuccess,
				MockInsertSuccess1:     tt.fields.MockInsertSuccess1,
				MockInsertSuccess2:     tt.fields.MockInsertSuccess2,
				MockInsertSuccess3:     tt.fields.MockInsertSuccess3,
				MockInsertSuccess4:     tt.fields.MockInsertSuccess4,
				MockInsertSuccess5:     tt.fields.MockInsertSuccess5,
				MockInsertSuccess6:     tt.fields.MockInsertSuccess6,
				MockInsertSuccess7:     tt.fields.MockInsertSuccess7,
				MockInsertSuccess8:     tt.fields.MockInsertSuccess8,
				mockInsertSuccess1Used: tt.fields.mockInsertSuccess1Used,
				mockInsertSuccess2Used: tt.fields.mockInsertSuccess2Used,
				mockInsertSuccess3Used: tt.fields.mockInsertSuccess3Used,
				mockInsertSuccess4Used: tt.fields.mockInsertSuccess4Used,
				mockInsertSuccess5Used: tt.fields.mockInsertSuccess5Used,
				mockInsertSuccess6Used: tt.fields.mockInsertSuccess6Used,
				mockInsertSuccess7Used: tt.fields.mockInsertSuccess7Used,
				mockInsertSuccess8Used: tt.fields.mockInsertSuccess8Used,
				MockInsertID1:          tt.fields.MockInsertID1,
				MockInsertID2:          tt.fields.MockInsertID2,
				MockInsertID3:          tt.fields.MockInsertID3,
				MockInsertID4:          tt.fields.MockInsertID4,
				MockInsertID5:          tt.fields.MockInsertID5,
				MockInsertID6:          tt.fields.MockInsertID6,
				MockInsertID7:          tt.fields.MockInsertID7,
				MockInsertID8:          tt.fields.MockInsertID8,
				MockUpdateSuccess1:     tt.fields.MockUpdateSuccess1,
				MockUpdateSuccess2:     tt.fields.MockUpdateSuccess2,
				MockUpdateSuccess3:     tt.fields.MockUpdateSuccess3,
				MockUpdateSuccess4:     tt.fields.MockUpdateSuccess4,
				mockUpdateSuccess1Used: tt.fields.mockUpdateSuccess1Used,
				mockUpdateSuccess2Used: tt.fields.mockUpdateSuccess2Used,
				mockUpdateSuccess3Used: tt.fields.mockUpdateSuccess3Used,
				mockUpdateSuccess4Used: tt.fields.mockUpdateSuccess4Used,
				MockDeleteSuccess1:     tt.fields.MockDeleteSuccess1,
				MockDeleteSuccess2:     tt.fields.MockDeleteSuccess2,
				MockDeleteSuccess3:     tt.fields.MockDeleteSuccess3,
				MockDeleteSuccess4:     tt.fields.MockDeleteSuccess4,
				MockDeleteSuccess5:     tt.fields.MockDeleteSuccess5,
				MockDeleteSuccess6:     tt.fields.MockDeleteSuccess6,
				MockDeleteSuccess7:     tt.fields.MockDeleteSuccess7,
				MockDeleteSuccess8:     tt.fields.MockDeleteSuccess8,
				mockDeleteSuccess1Used: tt.fields.mockDeleteSuccess1Used,
				mockDeleteSuccess2Used: tt.fields.mockDeleteSuccess2Used,
				mockDeleteSuccess3Used: tt.fields.mockDeleteSuccess3Used,
				mockDeleteSuccess4Used: tt.fields.mockDeleteSuccess4Used,
				mockDeleteSuccess5Used: tt.fields.mockDeleteSuccess5Used,
				mockDeleteSuccess6Used: tt.fields.mockDeleteSuccess6Used,
				mockDeleteSuccess7Used: tt.fields.mockDeleteSuccess7Used,
				mockDeleteSuccess8Used: tt.fields.mockDeleteSuccess8Used,
				MockTestRow:            tt.fields.MockTestRow,
				mockRow1Used:           tt.fields.mockRow1Used,
				mockRow2Used:           tt.fields.mockRow2Used,
				mockRow3Used:           tt.fields.mockRow3Used,
				mockRow4Used:           tt.fields.mockRow4Used,
				mockRow5Used:           tt.fields.mockRow5Used,
				mockRow6Used:           tt.fields.mockRow6Used,
				mockRow7Used:           tt.fields.mockRow7Used,
				mockRow8Used:           tt.fields.mockRow8Used,
				MockRow1:               tt.fields.MockRow1,
				MockRow2:               tt.fields.MockRow2,
				MockRow3:               tt.fields.MockRow3,
				MockRow4:               tt.fields.MockRow4,
				MockRow5:               tt.fields.MockRow5,
				MockRow6:               tt.fields.MockRow6,
				MockRow7:               tt.fields.MockRow7,
				MockRow8:               tt.fields.MockRow8,
				mockRows1Used:          tt.fields.mockRows1Used,
				mockRows2Used:          tt.fields.mockRows2Used,
				mockRows3Used:          tt.fields.mockRows3Used,
				mockRows4Used:          tt.fields.mockRows4Used,
				mockRows5Used:          tt.fields.mockRows5Used,
				mockRows6Used:          tt.fields.mockRows6Used,
				mockRows7Used:          tt.fields.mockRows7Used,
				mockRows8Used:          tt.fields.mockRows8Used,
				MockRows1:              tt.fields.MockRows1,
				MockRows2:              tt.fields.MockRows2,
				MockRows3:              tt.fields.MockRows3,
				MockRows4:              tt.fields.MockRows4,
				MockRows5:              tt.fields.MockRows5,
				MockRows6:              tt.fields.MockRows6,
				MockRows7:              tt.fields.MockRows7,
				MockRows8:              tt.fields.MockRows8,
			}
			if got := m.Delete(tt.args.query, tt.args.args...); got != tt.want {
				t.Errorf("MyDBMock.Delete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyDBMock_Close(t *testing.T) {
	type fields struct {
		Host                   string
		User                   string
		Password               string
		Database               string
		db                     *sql.DB
		err                    error
		MockConnectSuccess     bool
		MockCloseSuccess       bool
		MockCommitSuccess      bool
		MockRollbackSuccess    bool
		MockInsertSuccess1     bool
		MockInsertSuccess2     bool
		MockInsertSuccess3     bool
		MockInsertSuccess4     bool
		MockInsertSuccess5     bool
		MockInsertSuccess6     bool
		MockInsertSuccess7     bool
		MockInsertSuccess8     bool
		mockInsertSuccess1Used bool
		mockInsertSuccess2Used bool
		mockInsertSuccess3Used bool
		mockInsertSuccess4Used bool
		mockInsertSuccess5Used bool
		mockInsertSuccess6Used bool
		mockInsertSuccess7Used bool
		mockInsertSuccess8Used bool
		MockInsertID1          int64
		MockInsertID2          int64
		MockInsertID3          int64
		MockInsertID4          int64
		MockInsertID5          int64
		MockInsertID6          int64
		MockInsertID7          int64
		MockInsertID8          int64
		MockUpdateSuccess1     bool
		MockUpdateSuccess2     bool
		MockUpdateSuccess3     bool
		MockUpdateSuccess4     bool
		mockUpdateSuccess1Used bool
		mockUpdateSuccess2Used bool
		mockUpdateSuccess3Used bool
		mockUpdateSuccess4Used bool
		MockDeleteSuccess1     bool
		MockDeleteSuccess2     bool
		MockDeleteSuccess3     bool
		MockDeleteSuccess4     bool
		MockDeleteSuccess5     bool
		MockDeleteSuccess6     bool
		MockDeleteSuccess7     bool
		MockDeleteSuccess8     bool
		mockDeleteSuccess1Used bool
		mockDeleteSuccess2Used bool
		mockDeleteSuccess3Used bool
		mockDeleteSuccess4Used bool
		mockDeleteSuccess5Used bool
		mockDeleteSuccess6Used bool
		mockDeleteSuccess7Used bool
		mockDeleteSuccess8Used bool
		MockTestRow            *DbRow
		mockRow1Used           bool
		mockRow2Used           bool
		mockRow3Used           bool
		mockRow4Used           bool
		mockRow5Used           bool
		mockRow6Used           bool
		mockRow7Used           bool
		mockRow8Used           bool
		MockRow1               *DbRow
		MockRow2               *DbRow
		MockRow3               *DbRow
		MockRow4               *DbRow
		MockRow5               *DbRow
		MockRow6               *DbRow
		MockRow7               *DbRow
		MockRow8               *DbRow
		mockRows1Used          bool
		mockRows2Used          bool
		mockRows3Used          bool
		mockRows4Used          bool
		mockRows5Used          bool
		mockRows6Used          bool
		mockRows7Used          bool
		mockRows8Used          bool
		MockRows1              *DbRows
		MockRows2              *DbRows
		MockRows3              *DbRows
		MockRows4              *DbRows
		MockRows5              *DbRows
		MockRows6              *DbRows
		MockRows7              *DbRows
		MockRows8              *DbRows
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
				MockCloseSuccess: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MyDBMock{
				Host:                   tt.fields.Host,
				User:                   tt.fields.User,
				Password:               tt.fields.Password,
				Database:               tt.fields.Database,
				db:                     tt.fields.db,
				err:                    tt.fields.err,
				MockConnectSuccess:     tt.fields.MockConnectSuccess,
				MockCloseSuccess:       tt.fields.MockCloseSuccess,
				MockCommitSuccess:      tt.fields.MockCommitSuccess,
				MockRollbackSuccess:    tt.fields.MockRollbackSuccess,
				MockInsertSuccess1:     tt.fields.MockInsertSuccess1,
				MockInsertSuccess2:     tt.fields.MockInsertSuccess2,
				MockInsertSuccess3:     tt.fields.MockInsertSuccess3,
				MockInsertSuccess4:     tt.fields.MockInsertSuccess4,
				MockInsertSuccess5:     tt.fields.MockInsertSuccess5,
				MockInsertSuccess6:     tt.fields.MockInsertSuccess6,
				MockInsertSuccess7:     tt.fields.MockInsertSuccess7,
				MockInsertSuccess8:     tt.fields.MockInsertSuccess8,
				mockInsertSuccess1Used: tt.fields.mockInsertSuccess1Used,
				mockInsertSuccess2Used: tt.fields.mockInsertSuccess2Used,
				mockInsertSuccess3Used: tt.fields.mockInsertSuccess3Used,
				mockInsertSuccess4Used: tt.fields.mockInsertSuccess4Used,
				mockInsertSuccess5Used: tt.fields.mockInsertSuccess5Used,
				mockInsertSuccess6Used: tt.fields.mockInsertSuccess6Used,
				mockInsertSuccess7Used: tt.fields.mockInsertSuccess7Used,
				mockInsertSuccess8Used: tt.fields.mockInsertSuccess8Used,
				MockInsertID1:          tt.fields.MockInsertID1,
				MockInsertID2:          tt.fields.MockInsertID2,
				MockInsertID3:          tt.fields.MockInsertID3,
				MockInsertID4:          tt.fields.MockInsertID4,
				MockInsertID5:          tt.fields.MockInsertID5,
				MockInsertID6:          tt.fields.MockInsertID6,
				MockInsertID7:          tt.fields.MockInsertID7,
				MockInsertID8:          tt.fields.MockInsertID8,
				MockUpdateSuccess1:     tt.fields.MockUpdateSuccess1,
				MockUpdateSuccess2:     tt.fields.MockUpdateSuccess2,
				MockUpdateSuccess3:     tt.fields.MockUpdateSuccess3,
				MockUpdateSuccess4:     tt.fields.MockUpdateSuccess4,
				mockUpdateSuccess1Used: tt.fields.mockUpdateSuccess1Used,
				mockUpdateSuccess2Used: tt.fields.mockUpdateSuccess2Used,
				mockUpdateSuccess3Used: tt.fields.mockUpdateSuccess3Used,
				mockUpdateSuccess4Used: tt.fields.mockUpdateSuccess4Used,
				MockDeleteSuccess1:     tt.fields.MockDeleteSuccess1,
				MockDeleteSuccess2:     tt.fields.MockDeleteSuccess2,
				MockDeleteSuccess3:     tt.fields.MockDeleteSuccess3,
				MockDeleteSuccess4:     tt.fields.MockDeleteSuccess4,
				MockDeleteSuccess5:     tt.fields.MockDeleteSuccess5,
				MockDeleteSuccess6:     tt.fields.MockDeleteSuccess6,
				MockDeleteSuccess7:     tt.fields.MockDeleteSuccess7,
				MockDeleteSuccess8:     tt.fields.MockDeleteSuccess8,
				mockDeleteSuccess1Used: tt.fields.mockDeleteSuccess1Used,
				mockDeleteSuccess2Used: tt.fields.mockDeleteSuccess2Used,
				mockDeleteSuccess3Used: tt.fields.mockDeleteSuccess3Used,
				mockDeleteSuccess4Used: tt.fields.mockDeleteSuccess4Used,
				mockDeleteSuccess5Used: tt.fields.mockDeleteSuccess5Used,
				mockDeleteSuccess6Used: tt.fields.mockDeleteSuccess6Used,
				mockDeleteSuccess7Used: tt.fields.mockDeleteSuccess7Used,
				mockDeleteSuccess8Used: tt.fields.mockDeleteSuccess8Used,
				MockTestRow:            tt.fields.MockTestRow,
				mockRow1Used:           tt.fields.mockRow1Used,
				mockRow2Used:           tt.fields.mockRow2Used,
				mockRow3Used:           tt.fields.mockRow3Used,
				mockRow4Used:           tt.fields.mockRow4Used,
				mockRow5Used:           tt.fields.mockRow5Used,
				mockRow6Used:           tt.fields.mockRow6Used,
				mockRow7Used:           tt.fields.mockRow7Used,
				mockRow8Used:           tt.fields.mockRow8Used,
				MockRow1:               tt.fields.MockRow1,
				MockRow2:               tt.fields.MockRow2,
				MockRow3:               tt.fields.MockRow3,
				MockRow4:               tt.fields.MockRow4,
				MockRow5:               tt.fields.MockRow5,
				MockRow6:               tt.fields.MockRow6,
				MockRow7:               tt.fields.MockRow7,
				MockRow8:               tt.fields.MockRow8,
				mockRows1Used:          tt.fields.mockRows1Used,
				mockRows2Used:          tt.fields.mockRows2Used,
				mockRows3Used:          tt.fields.mockRows3Used,
				mockRows4Used:          tt.fields.mockRows4Used,
				mockRows5Used:          tt.fields.mockRows5Used,
				mockRows6Used:          tt.fields.mockRows6Used,
				mockRows7Used:          tt.fields.mockRows7Used,
				mockRows8Used:          tt.fields.mockRows8Used,
				MockRows1:              tt.fields.MockRows1,
				MockRows2:              tt.fields.MockRows2,
				MockRows3:              tt.fields.MockRows3,
				MockRows4:              tt.fields.MockRows4,
				MockRows5:              tt.fields.MockRows5,
				MockRows6:              tt.fields.MockRows6,
				MockRows7:              tt.fields.MockRows7,
				MockRows8:              tt.fields.MockRows8,
			}
			if got := m.Close(); got != tt.want {
				t.Errorf("MyDBMock.Close() = %v, want %v", got, tt.want)
			}
		})
	}
}
