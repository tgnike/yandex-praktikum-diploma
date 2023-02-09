package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
	"github.com/tgnike/yandex-praktikum-diploma/internal/accruals"
	"github.com/tgnike/yandex-praktikum-diploma/internal/authservice"
	"github.com/tgnike/yandex-praktikum-diploma/internal/authservice/authrepository"
	"github.com/tgnike/yandex-praktikum-diploma/internal/balanceservice"
	"github.com/tgnike/yandex-praktikum-diploma/internal/balanceservice/balancerepository"
	"github.com/tgnike/yandex-praktikum-diploma/internal/config"
	"github.com/tgnike/yandex-praktikum-diploma/internal/orderservice"
	"github.com/tgnike/yandex-praktikum-diploma/internal/orderservice/ordersrepository"
	"github.com/tgnike/yandex-praktikum-diploma/internal/postgres"
	"github.com/tgnike/yandex-praktikum-diploma/internal/router"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
	"github.com/tgnike/yandex-praktikum-diploma/internal/storage"
)

func main() {

	cfg := config.New()
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()

	err = cfg.Check()

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgx := &postgres.Storage{DataSourceName: cfg.DSN}
	pgx.Init()

	accruals := accruals.New(cfg.AccrualSystemAddress)
	go accruals.Start(ctx)

	server := makeServer(ctx, pgx, accruals)

	r := router.NewRouter(server)

	// запуск сервера
	log.Fatal(http.ListenAndServe(cfg.RunAddress, r))

}

func makeServer(ctx context.Context, pgx *postgres.Storage, accruals orderservice.Accruals) *server.Server {

	storage := storage.NewStore(pgx)

	users := authservice.New(authrepository.New(storage))
	orders := orderservice.New(ordersrepository.New(storage), accruals)
	go orders.UpdateAccrualInformation(ctx)
	balance := balanceservice.New(balancerepository.New(storage))

	return server.New(users, orders, balance)

}
