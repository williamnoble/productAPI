package middleware

import (
	"context"
	"fmt"
	"net/http"
	"productAPI/data"
)

type productKey string
const ProductKey productKey = "Key"

func MiddlewareValidatorProduct(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := data.Product{}
		err := p.FromJSON(r.Body) // not R as we need a reader
		if err != nil {
			fmt.Println("ERROR in middleware: ", err)
			http.Error(w, "error in unmarshalling data", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), ProductKey, p)
		response := r.WithContext(ctx)
		next.ServeHTTP(w,response)


	})

}