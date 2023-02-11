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

type Connector struct {
	Address  string
	register chan models.OrderNumber
	check    chan models.OrderNumber
	updates  chan *models.AccrualInformation
}

const checkInterval time.Duration = 1 * time.Second

func (c *Connector) Start(ctx context.Context) {

	ctxRegister, cancelRegister := context.WithCancel(ctx)
	ctxUpdate, cancelUpdate := context.WithCancel(ctx)

	go c.registration(ctxRegister)
	go c.update(ctxUpdate)

	for {
		select {
		case <-ctx.Done():
			cancelRegister()
			cancelUpdate()
		}
	}

}

func (c *Connector) GetUpdates() chan *models.AccrualInformation {
	return c.updates
}

func (c *Connector) registration(ctx context.Context) {

	for {
		select {
		case order := <-c.register:
			err := c.postOrder(order)

			if err != nil {
				log.Print(err)
				continue
			}

			c.check <- order

		case <-ctx.Done():
			return

		}

	}

}

func (c *Connector) update(ctx context.Context) {

	for {
		select {
		case order := <-c.check:
			orderinfo, err := c.getOrder(order)

			if err != nil {
				log.Print(err)

			}
			if orderinfo.Status == models.REGISTERED {
				c.Check(orderinfo.Order)
				continue
			}
			c.updates <- orderinfo
		case <-ctx.Done():
			return

		}

	}

}

func New(address string) *Connector {
	return &Connector{Address: address,
		register: make(chan models.OrderNumber),
		check:    make(chan models.OrderNumber),
		updates:  make(chan *models.AccrualInformation)}
}

func (c *Connector) Register(order models.OrderNumber) {

	c.register <- order

}

func (c *Connector) Check(order models.OrderNumber) {

	c.check <- order

}

func (c *Connector) postOrder(order models.OrderNumber) error {

	entryPoint := fmt.Sprintf("http://%s/api/orders", c.Address)

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`{"order":"%s"}`, order)).
		Post(entryPoint)

	if err != nil {
		return err
	}

	if resp.StatusCode() != 202 {

		return errors.New("wrong")

	}

	return nil
}

func (c *Connector) getOrder(order models.OrderNumber) (*models.AccrualInformation, error) {

	entryPoint := fmt.Sprintf("http://%s/api/orders/%s", c.Address, string(order))

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
