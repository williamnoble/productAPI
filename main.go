package main

import (
	"context"
	"flag"
	handlers2 "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"productAPI/fs"
	"productAPI/handlers"
	middleware "productAPI/midddleware"
	"time"
)

var port = flag.String("bind_address", ":9090", "Bind address for HTTP Server API")

func main() {
	const BasePath = "./images"
	flag.Parse()

	l := log.New(os.Stdout, "HTTP-API", log.LstdFlags)

	permitCors := handlers2.CORS(handlers2.AllowedOrigins([]string{"*"}))

	store := fs.NewLocalFileSystem(BasePath, 1024*1000*5)
	productsHandler := handlers.NewProducts(l)
	filesHandler := handlers.NewFileHandler(l, store)
	routes := mux.NewRouter()


	routes.HandleFunc("/images/{id:[0-9]+}/{filename:[a-zA-Z]+\\.[a-z]{3}}", filesHandler.Upload)
	routes.Handle("/add", middleware.MiddlewareValidatorProduct(http.HandlerFunc(productsHandler.AddProduct))).Methods(http.MethodPost)
	routes.Handle("/update/{id}", middleware.MiddlewareValidatorProduct(http.HandlerFunc(productsHandler.UpdateProduct))).Methods(http.MethodPut)


	//routes.Handle("/update/{id}", middleware.MiddlewareValidatorProduct(http.HandlerFunc(productsHandler.UpdateProduct))).Methods(http.MethodPut)

	routes.HandleFunc("/", productsHandler.GetProducts).Methods((http.MethodGet))

	srv := http.Server{
		Addr:         *port,
		Handler:      permitCors(routes),
		ReadTimeout:   1 * time.Second,
		WriteTimeout:  1 * time.Second,
		ErrorLog:     l,
	}


	go func() {	err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <- c
	log.Println("Got Signal", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	srv.Shutdown(ctx)



}
