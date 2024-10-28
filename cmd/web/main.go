package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strings"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func neuter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app := &application{errorLog: errorLog, infoLog: infoLog}

	// Router
	mux := http.NewServeMux()

	defaultAddr := os.Getenv("ADDR")

	addr := flag.String("addr", defaultAddr, "HTTP network address")
	flag.Parse()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// All URL paths that starts with "/static/"
	// Strip "/static" prefix before the request reachers the file server
	mux.Handle("/static/", http.StripPrefix("/static", neuter(fileServer)))

	f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// Regiter the other applications routes as normal
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()

	errorLog.Fatal(err)
}
