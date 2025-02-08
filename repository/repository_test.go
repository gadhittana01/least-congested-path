package repository

import (
	"testing"

	mockGraph "github.com/gadhittana01/least-congested-path/utils/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func initRepository(
	ctrl *gomock.Controller,
) (IRepository, *mockGraph.MockIGraph) {
	mockGraph := mockGraph.NewMockIGraph(ctrl)
	repository := NewRepository(mockGraph)
	return repository, mockGraph
}

func TestFindLeastCongestedPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repositoryMock, mockGraph := initRepository(ctrl)

	t.Run("success - find least congested path", func(t *testing.T) {
		start := "A"
		end := "B"
		expectedPath := "A -> B"

		mockGraph.EXPECT().FindLeastCongestedPath(start, end).Return(expectedPath).Times(1)

		path := repositoryMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, expectedPath, path)
	})

	t.Run("failure - no path found", func(t *testing.T) {
		start := "A"
		end := "C"
		noPathFound := "No Path Found"

		mockGraph.EXPECT().FindLeastCongestedPath(start, end).Return(noPathFound).Times(1)

		path := repositoryMock.FindLeastCongestedPath(start, end)

		assert.Equal(t, noPathFound, path)
	})
}

func TestAddRoad(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock, mockGraph := initRepository(ctrl)

	t.Run("success - add road", func(t *testing.T) {
		start := "A"
		end := "B"
		congestion := 5
		mockGraph.EXPECT().AddEdge(start, end, congestion).Times(1)

		serviceMock.AddRoad(start, end, congestion)
	})
}
