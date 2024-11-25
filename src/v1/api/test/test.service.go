package test

type TestService struct {
	model *TestModel
}

func (s *TestService) NewService() *TestService {
	model := TestModel{}
	return &TestService{
		model: model.NewModel(),
	}
}
