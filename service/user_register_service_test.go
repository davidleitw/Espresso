package service

import (
	"Espresso/models"
	serial "Espresso/serialization"
	"reflect"
	"testing"
)

func TestUserRegisterCatcher_Register(t *testing.T) {
	type fields struct {
		UserMail        string
		UserName        string
		UserPass        string
		UserPassConfirm string
		UserRid         string
	}
	tests := []struct {
		name   string
		fields fields
		want   *models.Users
		want1  *serial.Response
	}{
		// TODO: Add test cases.
		{
			name: "test data_1",
			fields: fields{
				UserMail:        "davidleitw@gmail.com",
				UserName:        "UserTestData1",
				UserPass:        "UserPassTest1",
				UserPassConfirm: "UserPassTest1",
				UserRid:         "0975687944",
			},
			want: &models.Users{
				User_Id:    "davidleitw@gmail.com",
				User_Name:  "UserTestData1",
				PassWord:   "UserPassTest1",
				Mobile_Rid: "0975687944",
			},
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserRegisterCatcher{
				UserMail:        tt.fields.UserMail,
				UserName:        tt.fields.UserName,
				UserPass:        tt.fields.UserPass,
				UserPassConfirm: tt.fields.UserPassConfirm,
				UserRid:         tt.fields.UserRid,
			}
			got, got1 := service.Register()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRegisterCatcher.Register() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("UserRegisterCatcher.Register() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
