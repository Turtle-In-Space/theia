package scanners

type ServiceScanner interface {
	run()
}

var ServiceRegistry = make(map[string]ServiceScanner)

func Register(name string, scanner ServiceScanner) {
	ServiceRegistry[name] = scanner
}
