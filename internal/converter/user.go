package converter

import (
	"github.com/biryanim/auth/internal/model"
	desc "github.com/biryanim/auth/pkg/user_api_v1"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamp.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     info.Name,
		Username: info.Username,
		Email:    info.Email,
		Role:     desc.Role(info.Role),
	}
}

func ToUpdatedUserInfoFromService(updateInfo model.UpdateUserInfo) *desc.UpdateUserInfo {
	return &desc.UpdateUserInfo{
		Name:  stringToWrapper(updateInfo.Name),
		Email: stringToWrapper(updateInfo.Email),
	}
}

func stringToWrapper(s *string) *wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	return &wrapperspb.StringValue{Value: *s}
}

func ToUserInfoFromDesc(userInfo *desc.UserInfo, password string) *model.UserInfo {
	return &model.UserInfo{
		Name:     userInfo.Name,
		Username: userInfo.Username,
		Email:    userInfo.Email,
		Role:     int32(userInfo.Role),
		Password: password,
	}
}

func ToUpdatedUserInfoFromDesc(updateInfo *desc.UpdateUserInfo) *model.UpdateUserInfo {
	var name, email *string
	if updateInfo.Name != nil {
		nameValue := updateInfo.Name.GetValue()
		name = &nameValue
	}
	if updateInfo.Email != nil {
		emailValue := updateInfo.Email.GetValue()
		email = &emailValue
	}
	return &model.UpdateUserInfo{
		Name:  name,
		Email: email,
	}
}
