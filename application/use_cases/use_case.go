package use_cases

type UseCase interface {
	Execute(input interface{}) (*interface{}, error)
}
