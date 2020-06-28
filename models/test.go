package models

type TestModel struct {
	//
}

func (test *TestModel) GetTestInfo(name string) string {
	name += "1"
	return name
}
