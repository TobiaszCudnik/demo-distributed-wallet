package types

import (
	pb "github.com/TobiaszCudnik/demo-distributed-wallet/api/wallet_pb"
)

type WalletId string

type ApiError struct {
	Code pb.ErrCode
}

func (err ApiError) Error() string {
	return pb.ErrCode_name[int32(err.Code)]
}

func NewApiError(code pb.ErrCode) error {
	return ApiError{code}
}
