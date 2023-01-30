package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tgnike/yandex-praktikum-diploma/internal/authservice/authhandlers"
	"github.com/tgnike/yandex-praktikum-diploma/internal/orderservice/ordershandlers"
	"github.com/tgnike/yandex-praktikum-diploma/internal/server"
)

func NewRouter(server *server.Server) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5, "application/json", "html/text", "text/html", ""))
	r.Use(middleware.Recoverer)

	// POST /api/user/register — регистрация пользователя;
	// POST /api/user/login — аутентификация пользователя;
	// POST /api/user/orders — загрузка пользователем номера заказа для расчёта;
	// GET /api/user/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
	// GET /api/user/balance — получение текущего баланса счёта баллов лояльности пользователя;
	// POST /api/user/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
	// GET /api/user/withdrawals

	r.Route("/api/user", func(r chi.Router) {

		r.Post("/register", handle(&authhandlers.RegistationHandler{Service: server.Users}))
		r.Post("/login", handle(&authhandlers.LoginHandler{Service: server.Users}))

		r.Post("/orders", handle(&ordershandlers.PostOrderHandler{Service: server.Orders}))
		r.Get("/orders", handle(&ordershandlers.GetOrdersHandler{Service: server.Orders}))

		r.Post("/balance/withdraw", mockOk())
		r.Get("/balance", mockOk())

		r.Get("/withdrawals", mockOk())

	})

	return r
}

// Хэлпер
func handle(handler http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})

}
