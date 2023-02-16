package service

type human struct {
	name string
}

func NewHuman(name string) human {
	var localHuman human
	localHuman.name = name
	return localHuman
}

func (h human) GetHumanName() string {
	return h.name
}
