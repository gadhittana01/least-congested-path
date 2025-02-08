package service

type IRepository interface {
	FindLeastCongestedPath(start, end string) string
	AddRoad(start, end string, congestion int)
}

type IService interface {
	FindLeastCongestedPath(start, end string) string
	AddRoad(start, end string, congestion int)
}

type service struct {
	repository IRepository
}

func NewService(repository IRepository) IService {
	return &service{
		repository: repository,
	}
}

func (h *service) FindLeastCongestedPath(start, end string) string {
	return h.repository.FindLeastCongestedPath(start, end)
}

func (h *service) AddRoad(start, end string, congestion int) {
	h.repository.AddRoad(start, end, congestion)
}
