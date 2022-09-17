package single

type Service struct {
	Config string
	svc    *service
}

func New() *Service {
	return &Service{ // want `Config is missing in new Service`
		svc: new(),
	}
}
