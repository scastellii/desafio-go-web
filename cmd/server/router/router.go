package router

import (
	handler2 "github.com/bootcamp-go/desafio-go-web/cmd/handler"
	"github.com/bootcamp-go/desafio-go-web/internal/domain"
	"github.com/bootcamp-go/desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Handler *handler2.Service
	Route   *gin.Engine
}

func (r *Router) MapRoutes() {
	group := r.Route.Group("/tickets")
	group.GET("/getByCountry/:dest", r.Handler.GetTicketsByCountry())
	group.GET("/getAverage/:dest", r.Handler.AverageDestination())
}

func NewRouter(r *gin.Engine, list []domain.Ticket) Router {
	repo := tickets.NewRepository(list)
	service := tickets.NewService(repo)
	handler := handler2.NewService(service)
	return Router{
		Route:   r,
		Handler: handler,
	}
}
