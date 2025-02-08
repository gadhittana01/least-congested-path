package handler

import (
	"testing"

	mocksvc "github.com/gadhittana01/least-congested-path/service/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initHandler(
	ctrl *gomock.Controller,
) (IService, *mocksvc.MockIService) {
	mockSvc := mocksvc.NewMockIService(ctrl)
	handler := NewHandler(mockSvc)
	return handler, mockSvc
}

func TestFindLeastCongestedPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handlerMock, mockSvc := initHandler(ctrl)

	t.Run("success - find least congested path", func(t *testing.T) {
		start := "A"
		end := "B"
		expectedPath := "A -> B"

		mockSvc.EXPECT().FindLeastCongestedPath(start, end).Return(expectedPath).Times(1)

		path := handlerMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, expectedPath, path)
	})

	t.Run("failure - no path found", func(t *testing.T) {
		start := "A"
		end := "C"
		noPathFound := "No Path Found"

		mockSvc.EXPECT().FindLeastCongestedPath(start, end).Return(noPathFound).Times(1)

		path := handlerMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, noPathFound, path)
	})
}

func TestAddRoad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	handlerMock, mockRepo := initHandler(ctrl)

	t.Run("success - add road", func(t *testing.T) {
		start := "A"
		end := "B"
		congestion := 5
		mockRepo.EXPECT().AddRoad(start, end, congestion).Times(1)

		handlerMock.AddRoad(start, end, congestion)
	})
}
