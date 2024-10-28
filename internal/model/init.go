package model

var Models = make([]interface{}, 0)

func RegisterModels(m ...interface{}) {
	Models = append(Models, m...)
}
