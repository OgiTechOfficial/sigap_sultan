package common

type (
	CRUDService[T any] interface {
		Get(params T) (interface{}, error)
		Create(params T) (interface{}, error)
		Update(params T) (interface{}, error)
		Delete(params T) (interface{}, error)
	}
)
