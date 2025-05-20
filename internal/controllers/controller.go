package controllers

type ControllerShared struct {
	Name string
}

func (me *ControllerShared) GetName() string {
	return me.Name
}
