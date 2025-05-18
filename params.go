package runner

type ParamsInterface interface {}

func GetParams[T any](params *ParamsInterface) T {
	return (*params).(T)
}

func SetParams[T any](params *ParamsInterface, t T) {
	*params = t
}