package service

import (
	"testing"

	mockrepo "github.com/gadhittana01/least-congested-path/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initService(
	ctrl *gomock.Controller,
) (IService, *mockrepo.MockIRepository) {
	mockRepo := mockrepo.NewMockIRepository(ctrl)
	service := NewService(mockRepo)
	return service, mockRepo
}

func TestFindLeastCongestedPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock, mockRepo := initService(ctrl)

	t.Run("success - find least congested path", func(t *testing.T) {
		start := "A"
		end := "B"
		expectedPath := "A -> B"

		mockRepo.EXPECT().FindLeastCongestedPath(start, end).Return(expectedPath).Times(1)

		path := serviceMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, expectedPath, path)
	})

	t.Run("failure - no path found", func(t *testing.T) {
		start := "A"
		end := "C"
		noPathFound := "No Path Found"

		mockRepo.EXPECT().FindLeastCongestedPath(start, end).Return(noPathFound).Times(1)

		path := serviceMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, noPathFound, path)
	})
}

func TestAddRoad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock, mockRepo := initService(ctrl)

	t.Run("success - add road", func(t *testing.T) {
		start := "A"
		end := "B"
		congestion := 5
		mockRepo.EXPECT().AddRoad(start, end, congestion).Times(1)

		serviceMock.AddRoad(start, end, congestion)
	})
}
