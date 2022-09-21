package main

import "github.com/PavelDonchenko/40projects/go-bookstore/cmd/server"

func main() {
	server.Run()
	//var wait time.Duration
	//flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	//flag.Parse()
	//
	//router := mux.NewRouter()
	//routes.RegisterBookStoreRoutes(router)
	//
	//srv := &http.Server{
	//	Addr:         "localhost:6666",
	//	WriteTimeout: time.Second * 15,
	//	ReadTimeout:  time.Second * 15,
	//	IdleTimeout:  time.Second * 60,
	//	Handler:      router,
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil {
	//		log.Println(err)
	//	}
	//}()
	//
	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt)
	//
	//<-c
	//
	//ctx, cancel := context.WithTimeout(context.Background(), wait)
	//defer cancel()
	//srv.Shutdown(ctx)
	//log.Println("shutting down")
	//os.Exit(0)
}
