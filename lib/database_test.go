package lib

import (
	"database/sql"
	"linebot-server/stru"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAllUserFromDB(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		mock func()
		want []stru.User
	}{
		// TODO: Add test cases.
		{
			name: "user table",
			args: args{
				db: db,
			},
			mock: func() {
				rows := mock.NewRows([]string{"ID", "UID", "Name"}).AddRow(1, "hi", "hihi")
				mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
			},
			want: []stru.User{
				{
					ID:     1,
					UserID: "hi",
					Name:   "hihi",
				},
			},
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if got := AllUserFromDB(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DataFromDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
