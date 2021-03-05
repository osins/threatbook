package domain

// Result type
type Result struct {
	Success string
	Code    int32
	Message string
}

// SetSuccess inner method
func (s *Result) SetSuccess(val string) *Result {
	s.Success = val
	return s
}

// SetCode inner method
func (s *Result) SetCode(val int32) *Result {
	s.Code = val
	return s
}

// SetMessage inner method
func (s *Result) SetMessage(val string) *Result {
	s.Message = val
	return s
}
