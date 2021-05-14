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
	field := []string{"ID", "UserID", "DisplayName", "PictureURL"}
	query := `SELECT ID, UserID, DisplayName, PictureURL FROM Users WHERE DisplayName=?`
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
					AddRow(0, "userId", "hi", "url").
					AddRow(1, "userId1", "name1", "url")
				// actions := []driver.Value{"hi"}
				mock.ExpectQuery(query).WithArgs("hi").WillReturnRows(rows)
			},
			want: User{
				ID:          0,
				UserID:      "userId",
				DisplayName: "hi",
				PictureUrl:  "url",
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
					AddRow(0, "userId", "hi", "url")
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
	field := []string{"ID", "UserID", "DisplayName", "PictureURL"}
	query := `SELECT ID, UserID, DisplayName, PictureURL FROM Users`
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
					AddRow(1, "uid", "name", "url").
					AddRow(2, "uid2", "name2", "url2")
				mock.ExpectQuery(query).WillReturnRows(rows)
			},
			want: []User{
				{
					ID:          1,
					UserID:      "uid",
					DisplayName: "name",
					PictureUrl:  "url",
				},
				{
					ID:          2,
					UserID:      "uid2",
					DisplayName: "name2",
					PictureUrl:  "url2",
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
	db, mock := NewMock()
	type args struct {
		db      *sql.DB
		profile *linebot.UserProfileResponse
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create user",
			args: args{
				db: db,
				profile: &linebot.UserProfileResponse{
					UserID:      "uID",
					DisplayName: "name",
					PictureURL:  "url",
				},
			},
			mock: func() {
				prep := mock.ExpectPrepare("INSERT INTO Users")
				prep.ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.db, tt.args.profile); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserByAccessToken(t *testing.T) {
	type args struct {
		accessToken string
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Get user by access token: no access token",
			args: args{
				accessToken: "",
			},
			want:    User{},
			wantErr: true,
		},
		// Postpone
		// {
		// 	name: "Get user by access token: success",
		// 	args: args{
		// 		accessToken: "",
		// 	},
		// 	want:    User{},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserByAccessToken(tt.args.accessToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
