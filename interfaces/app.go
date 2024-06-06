package interfaces

type AppKind int

const (
	AppTypeDesktop AppKind = iota
	AppTypeWebc
	AppTypeCLI
)

type IApp interface {
	Exec()
}

type IFactoryApp interface {
	MakeApp(ak AppKind) IApp
}
