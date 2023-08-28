package main

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// // 本物のDBにつなぐ
// func Test_GetHeadUserNameByID(t *testing.T) {
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want int
// 	}{
// 		{
// 			name: "success",
// 			args: 1,
// 			want: "20文字太郎20文字",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := GetHeadUserNameByID(tt.args.i); got != tt.want {
// 				t.Errorf("Add() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_FakeGetHeadUserNameByID(t *testing.T) {
// 	//イメージが微妙
// 	type args struct {
// 		id int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		{
// 			name: "ちゃんとしたユーザ",
// 			args: 1,
// 			want: "20文字太郎20文字",
// 		},
// 		{
// 			name: "nil",
// 			args: 2,
// 			want: nil,
// 		},
// 		{
// 			name: "空文字",
// 			args: 2,
// 			want: "",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := FakeGetHeadUserNameByID(tt.args.id); got != tt.want {
// 				t.Errorf("Add() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// // mockを使う
// func GetHeadUserNameByID_Mock() {}

func Test_WithRawDB(t *testing.T) {
	tests := []struct {
		name   string
		userID int
		want   string
	}{
		{
			name:   "ちゃんとしたユーザ",
			userID: 1,
			want:   "20testtest",
		},
		{
			name:   "空文字",
			userID: 2,
			want:   "",
		},
		{
			name:   "nameが100文字くらいある",
			userID: 3,
			want:   "100testtes",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// usecaseの生成だけ変えるイメージ
			uc := UserUsecase{
				Repository: &userRepository{},
			}
			got, err := uc.GetHeadUserNameByID(tt.userID)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_WithFake(t *testing.T) {
	tests := []struct {
		name   string
		userID int
		want   string
	}{
		{
			name:   "ちゃんとしたユーザ",
			userID: 1,
			want:   "20testtest",
		},
		{
			name:   "空文字",
			userID: 2,
			want:   "",
		},
		{
			name:   "nameが100文字くらいある",
			userID: 3,
			want:   "100testtes",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc := UserUsecase{
				Repository: &fakeUserRepository{},
			}
			got, err := uc.GetHeadUserNameByID(tt.userID)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

type UserMock struct{}

func Test_WithMock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockUser := NewMockUserRepository(mockCtrl)
	uc := UserUsecase{
		Repository: mockUser,
	}
	tests := []struct {
		name     string
		userID   int
		userName string
		want     string
	}{
		{
			name:     "ちゃんとしたユーザ",
			userID:   1,
			userName: "20testtest1",
			want:     "20testtest",
		},
		{
			name:     "空文字",
			userID:   2,
			userName: "",
			want:     "",
		},
		{
			name:     "nameが100文字くらいある",
			userID:   3,
			userName: "100testtesaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			want:     "100testtes",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			mockUser.EXPECT().GetUserNameByID(tt.userID).Return(tt.userName, nil)
			got, err := uc.GetHeadUserNameByID(tt.userID)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
