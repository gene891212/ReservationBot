package models

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/line/line-bot-sdk-go/linebot"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestGetUser(t *testing.T) {
	db, mock := NewMock()
	field := []string{"ID", "UserID", "DisplayName"}
	query := `SELECT ID, UserID, DisplayName FROM Users WHERE DisplayName=?`
	defer db.Close()

	type args struct {
		db   *sql.DB
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get user by name",
			args: args{
				db:   db,
				name: "hi",
			},
			mock: func() {
				rows := sqlmock.NewRows(field).
					AddRow(0, "userId", "hi").
					AddRow(1, "userId1", "name1")
				// actions := []driver.Value{"hi"}
				mock.ExpectQuery(query).WithArgs("hi").WillReturnRows(rows)
			},
			want: User{
				ID:          0,
				UserID:      "userId",
				DisplayName: "hi",
			},
			wantErr: false,
		},
		{
			name: "Get user by name no match",
			args: args{
				db:   db,
				name: "hi",
			},
			mock: func() {
				rows := sqlmock.NewRows(field).
					AddRow(0, "userId", "hi")
				mock.ExpectQuery(query).WithArgs("ok").WillReturnRows(rows)
			},
			want:    User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUser(tt.args.db, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	field := []string{"ID", "UserID", "DisplayName"}
	query := `SELECT ID, UserID, DisplayName FROM Users`
	defer db.Close()

	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "No data in Users",
			args: args{db},
			mock: func() {
				rows := mock.NewRows(field)
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want:    []User{},
			wantErr: false,
		},
		{
			name: "Get All Users",
			args: args{db},
			mock: func() {
				rows := mock.NewRows(field).
					AddRow(1, "uid", "name").
					AddRow(2, "uid2", "name2")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []User{
				{
					ID:          1,
					UserID:      "uid",
					DisplayName: "name",
				},
				{
					ID:          2,
					UserID:      "uid2",
					DisplayName: "name2",
				},
			},
			wantErr: false,
		},
		{
			name: "Query error",
			args: args{db},
			mock: func() {
				mock.ExpectQuery(query).WillReturnError(errors.New("Query error"))
			},
			want:    []User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUsers(tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	type args struct {
		db      *sql.DB
		profile *linebot.UserProfileResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.db, tt.args.profile); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
