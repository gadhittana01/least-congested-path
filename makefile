mockGraph:
	mockgen -package repository -source=./utils/graph.go -destination=./utils/mock/graph_mock.go

mockRepository:
	mockgen -package repository -source=./repository/repository.go -destination=./repository/mock/repository_mock.go

mockService:
	mockgen -package service -source=./service/service.go -destination=./service/mock/service_mock.go