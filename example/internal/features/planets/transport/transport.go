package transport

import (
	"context"
	"example/internal/features/planets"
	protoPlanets "example/protoc/planets"
	"fmt"

	"github.com/ihatiko/olymp/components/transports/cron"
	"github.com/ihatiko/olymp/components/transports/daemon"
)

type transport struct {
	service planets.IService
}

func New(service planets.IService) planets.ITransport {
	return &transport{}
}

func (t transport) UpdatePlanet(ctx context.Context, request *protoPlanets.UpdatePlanetRequest) (*protoPlanets.UpdatePlanetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (t transport) Load(request daemon.Request) error {
	fmt.Println(123)
	return nil
}

func (t transport) Update(request cron.Request) error {
	fmt.Println(123)
	return nil
}
