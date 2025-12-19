// Code generated manually. DO NOT EDIT.
// source: auth/v1/auth.proto

package v1

import "google.golang.org/protobuf/types/known/timestamppb"

type SignUpRequest struct {
	Email    string
	Password string
	Name     string
}

func (x *SignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *SignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SignUpResponse struct {
	UserId string
}

func (x *SignUpResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type SignInRequest struct {
	Email    string
	Password string
}

func (x *SignInRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SignInResponse struct {
	AccessToken string
	ExpiresAt   *timestamppb.Timestamp
}

func (x *SignInResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type ValidateTokenRequest struct {
	AccessToken string
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type ValidateTokenResponse struct {
	Valid     bool
	UserId    string
	ExpiresAt *timestamppb.Timestamp
}

func (x *ValidateTokenResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateTokenResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ValidateTokenResponse) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}
