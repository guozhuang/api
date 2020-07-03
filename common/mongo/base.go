package mongo

type IMongo interface {
	//Insert(map[string]string)
	//Update(map[string]string, map[string]string)
	//GetCount(map[string]string) int
	//GetList(map[string]string, int, int) []map[string]string
	//GetAll(map[string]string) []IEditor
	FindOne(map[string]string) map[string]string
	//Remove(map[string]string)
}

type ApplyFunc func()

func (f *ApplyFunc) FindOne(map[string]string) {
	//需要中转实现么？
}
