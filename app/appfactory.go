package app

import "github.com/Giovanny472/gMiniPlanner/interfaces"

type factoryApp struct{}

var fa *factoryApp

func NewFactoryApp() interfaces.IFactoryApp {

	if fa == nil {
		fa = &factoryApp{}
	}

	return fa
}

func (a *factoryApp) MakeApp(ak interfaces.AppKind) interfaces.IApp {

	switch ak {
	case interfaces.AppTypeDesktop:
		return NewAppDesktop()
	default:
		return NewAppDesktop()
	}

}
