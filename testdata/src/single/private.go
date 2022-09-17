package single

type service struct {
	s string
}

func new() *service {
	return &service{} // want `s is missing in new service`
}
