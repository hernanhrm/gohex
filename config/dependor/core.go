package dependor

var container map[string]any

func Init() {
	container = make(map[string]any)
}

func Set[T any](name string, value T) {
	if container == nil {
		container = make(map[string]any)
	}

	container[name] = value
}

func Get[T any](name string) T {
	return container[name].(T)
}
