package scw

import "context"

type RequestOption func(*ScalewayRequest)

func WithContext(ctx context.Context) RequestOption {
	return func(s *ScalewayRequest) {
		s.Ctx = ctx
	}
}
