package threatbook

// TestStruct type
type TestStruct struct {
	V string
}

// String inner method
func (s *TestStruct) String() string {
	return "dsfdsfsdf"
}

// Parameter interface
type Parameter interface {
	String() string
}

// Request type
type Request struct {
	Key string `json:"APIkey"`
}
