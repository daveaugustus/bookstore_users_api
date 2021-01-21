package services

var (
	ItemsServices itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
