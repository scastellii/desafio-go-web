package tickets

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetTotalTickets(c *gin.Context, dest string) (int, error)
	AverageDestination(c *gin.Context, dest string) (float64, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}

}

type service struct {
	repository Repository
}

func (s *service) GetTotalTickets(c *gin.Context, dest string) (total int, err error) {
	tickets, err := s.repository.GetTicketByDestination(c, dest)
	if err != nil {
		return
	}
	if len(tickets) == 0 {
		err = errors.New("No se encontraron tickets")
		return
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(c *gin.Context, dest string) (average float64, err error) {
	tickets, err := s.repository.GetTicketByDestination(c, dest)
	if err != nil {
		return
	}
	if len(tickets) == 0 {
		err = errors.New("No se encontraron tickets")
		return
	}

	all, err := s.repository.GetAll(c)
	if err != nil {
		return
	}

	average = float64(len(tickets)) / float64(len(all))
	return
}
