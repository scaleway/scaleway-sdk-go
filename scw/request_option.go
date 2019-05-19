package scw

import "context"

// RequestOption is a function that applies options to a ScalewayRequest.
type RequestOption func(*ScalewayRequest)

// WithContext request option sets the context of a ScalewayRequest
func WithContext(ctx context.Context) RequestOption {
	return func(s *ScalewayRequest) {
		s.Ctx = ctx
	}
}

func WithPage(page uint) RequestOption {
	return func(s *ScalewayRequest) {
		s.Pagination.Page = page
	}
}
