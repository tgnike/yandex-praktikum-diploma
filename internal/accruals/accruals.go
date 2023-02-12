package accruals

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/tgnike/yandex-praktikum-diploma/internal/models"
)

type Request struct {
	order *models.OrderNumber
	user  *models.UserID
}

type Connector struct {
	Address string
	//register chan models.OrderNumber
	check   chan *Request
	updates chan *models.AccrualInformation
}

const checkInterval time.Duration = 1 * time.Second

func (c *Connector) Start(ctx context.Context) {

	ctxUpdate, cancelUpdate := context.WithCancel(ctx)
	defer cancelUpdate()

	go c.update(ctxUpdate)

	<-ctx.Done()

}

func (c *Connector) GetUpdates() chan *models.AccrualInformation {
	return c.updates
}

func (c *Connector) update(ctx context.Context) {

	for {
		select {
		case request := <-c.check:
			orderinfo, err := c.getOrder(string(*request.order))

			if err != nil {
				log.Print(err)

			}
			if orderinfo.Status == models.REGISTERED {
				c.Check(request.order, request.user)
				continue
			}
			orderinfo.User = request.user
			c.updates <- orderinfo
		case <-ctx.Done():
			return

		}

	}

}

func New(address string) *Connector {
	return &Connector{Address: address,
		check:   make(chan *Request),
		updates: make(chan *models.AccrualInformation)}
}

func (c *Connector) Check(order *models.OrderNumber, user *models.UserID) {

	c.check <- &Request{user: user, order: order}

}

func (c *Connector) getOrder(order string) (*models.AccrualInformation, error) {

	entryPoint := fmt.Sprintf("%s/api/orders/%s", c.Address, string(order))

	oi := &models.AccrualInformation{}

	resp, err := resty.New().R().SetResult(oi).Get(entryPoint)

	if err != nil {
		return &models.AccrualInformation{}, err
	}

	if resp.StatusCode() != 200 {

		return &models.AccrualInformation{}, errors.New("wrong")

	}

	return oi, nil
}
