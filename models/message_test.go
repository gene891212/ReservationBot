package models

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateMessage(t *testing.T) {
	db, mock := NewMock()
	sender := User{
		ID:          0,
		DisplayName: "name0",
	}
	reciver := User{
		ID:          1,
		DisplayName: "name1",
	}
	now := time.Now()
	type args struct {
		db  *sql.DB
		msg Message
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Create Message",
			args: args{
				db: db,
				msg: Message{
					Content: "content",
					Sender:  sender,
					Reciver: reciver,
					Time:    now,
				},
			},
			mock: func() {
				// rows := sqlmock.NewRows([]string{"Content", "Sender", "Reciver", "Time"}).
				// 	AddRow("content", sender.ID, reciver.ID, now)
				prep := mock.ExpectPrepare(`INSERT INTO Messages`)
				prep.ExpectExec().WithArgs("content", 0, 1, now).WillReturnResult(sqlmock.NewResult(100, 1))
				// mock.ExpectExec(`INSERT INTO Messages`).
				// 	WithArgs("content", 0, 1, now).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.mock()
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateMessage(tt.args.db, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("CreateMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
