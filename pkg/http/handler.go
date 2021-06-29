package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zubayear/bruce-almighty/external/repo"
	myEndpoint "github.com/Zubayear/bruce-almighty/pkg/endpoint"
	"github.com/Zubayear/bruce-almighty/pkg/service"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	r      = repo.NewUserRepository()
	router = mux.NewRouter()
)

func NewHTTPServer(ctx context.Context, endpoints myEndpoint.Endpoints) http.Handler {
	router.Use(commonMiddleware)
	router.Methods("POST").Path("/create-user").Handler(
		httptransport.NewServer(
			endpoints.CreateUser,
			decodeCreateUserRequest,
			encodeResponse,
		))

	router.Methods("GET").Path("/get-user/{id}").Handler(
		httptransport.NewServer(
			endpoints.GetUser,
			decodeGetUserRequest,
			encodeResponse,
		))
	return router
}

func Run() {
	svc := service.New(r)
	e := myEndpoint.MakeEndpoints(svc)
	h := NewHTTPServer(context.Background(), e)
	http.ListenAndServe(":9000", h)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request myEndpoint.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request myEndpoint.GetUserRequest
	vars := mux.Vars(r)
	request = myEndpoint.GetUserRequest{
		Id: vars["id"],
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
