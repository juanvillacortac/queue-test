package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/juanvillacortac/bank-queue/pkg/handlers"
)

type Server struct {
	mux *chi.Mux
}

type ServerOptions struct {
	ServeFS        fs.FS
	FSPattern      string
	FSFallbackFile string
	ApiPrefix      string
}

func NewServer(options ServerOptions) Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Compress(5))
	r.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			if r.Method == "OPTIONS" {
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	})

	if options.ApiPrefix == "" {
		options.ApiPrefix = "/api"
	}
	if options.FSPattern == "" {
		options.FSPattern = "/*"
	}

	handlers.RegisterHandlers(r, options.ApiPrefix)

	if options.ServeFS != nil {
		r.Handle(options.FSPattern, SPAHandler(options.ServeFS, options.FSFallbackFile))
	}

	return Server{
		mux: r,
	}
}

func (s Server) Listen(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}

func SPAHandler(assets fs.FS, fallbackFile string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cleaned := strings.TrimPrefix(path.Clean(r.URL.Path), "/")
		cases := []string{
			cleaned,
			cleaned + ".html",
			path.Join(cleaned, "index.html"),
		}
		var found bool
		for _, file := range cases {
			f, err := assets.Open(file)
			if err == nil {
				defer f.Close()
			}
			found = f != nil && err == nil
			if found {
				r.URL.Path = path.Clean(fmt.Sprintf("/%s", file))
				if r.URL.Path == "" || r.URL.Path == "/index.html" {
					r.URL.Path = "/"
				}
				println(r.URL.Path)
				break
			}
		}
		if !found {
			if fallbackFile != "" {
				http.ServeFileFS(w, r, assets, fallbackFile)
				return
			} else {
				http.NotFound(w, r)
				return
			}
		}
		http.FileServer(http.FS(assets)).ServeHTTP(w, r)
	}
}
