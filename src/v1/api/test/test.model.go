package test

type TestModel struct{}

func (t *TestModel) NewModel() *TestModel {
	return &TestModel{}
}
