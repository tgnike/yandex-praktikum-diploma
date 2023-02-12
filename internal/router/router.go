package router

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tgnike/yandex-praktikum-diploma/internal/authservice/authhandlers"
	"github.com/tgnike/yandex-praktikum-diploma/internal/balanceservice/balancehandlers"
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

		r.With(AuthMiddleware(server.Users)).Route("/orders", func(r chi.Router) {
			r.Post("/", handle(&ordershandlers.PostOrderHandler{Service: server.Orders}))
			r.Get("/", handle(&ordershandlers.GetOrdersHandler{Service: server.Orders}))
		})

		r.With(AuthMiddleware(server.Users)).Route("/balance", func(r chi.Router) {
			r.Post("/withdraw", handle(&balancehandlers.WithdrawRequestHandler{Service: server.Balance}))
			r.Get("/", mockOk())
		})

		r.With(AuthMiddleware(server.Users)).Route("/withdrawals", func(r chi.Router) {
			r.Get("/", mockOk())
		})

	})

	return r
}

func AuthMiddleware(auth server.Users) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			tokenFull := r.Header.Get("Authorization")
			words := strings.Split(tokenFull, " ")

			userID, err := auth.CheckAuthToken(words[len(words)-1])

			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), server.UserContext, userID)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// Хэлпер
func handle(handler http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	})

}
