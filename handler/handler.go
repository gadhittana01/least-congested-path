package handler

type IService interface {
	FindLeastCongestedPath(start, end string) string
	AddRoad(start, end string, congestion int)
}

type IHandler interface {
	FindLeastCongestedPath(start, end string) string
	AddRoad(start, end string, congestion int)
}

type handler struct {
	service IService
}

func NewHandler(service IService) IHandler {
	return handler{
		service: service,
	}
}

func (h handler) FindLeastCongestedPath(start, end string) string {
	return h.service.FindLeastCongestedPath(start, end)
}

func (h handler) AddRoad(start, end string, congestion int) {
	h.service.AddRoad(start, end, congestion)
}
