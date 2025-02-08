package main

import (
	"fmt"

	"github.com/gadhittana01/least-congested-path/handler"
	"github.com/gadhittana01/least-congested-path/repository"
	"github.com/gadhittana01/least-congested-path/service"
	"github.com/gadhittana01/least-congested-path/utils"
)

func main() {
	graph := utils.NewGraph()
	repository := repository.NewRepository(graph)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	handler.AddRoad("NE 42nd Way", "NE 42nd St", 1)
	handler.AddRoad("NE 42nd St", "NE 39th St", 2)
	handler.AddRoad("NE 39th St", "206th Pl NE", 2)
	handler.AddRoad("NE 39th St", "204th Ave NE", 2)
	handler.AddRoad("NE 42nd St", "203rd Ave NE", 2)
	handler.AddRoad("203rd Ave NE", "204th Ave NE", 1)
	handler.AddRoad("204th Ave NE", "206th Pl NE", 1)

	path := handler.FindLeastCongestedPath("NE 42nd Way", "206th Pl NE")
	fmt.Println("Recommended Route:", path)
}
