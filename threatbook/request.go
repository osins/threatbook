package threatbook

type TestStruct struct {
	V string
}

func (s *TestStruct) String() string {
	return "dsfdsfsdf"
}

type Parameter interface {
	String() string
}

type Request struct {
	Key string `json:"apikey"`
}
