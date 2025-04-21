package static

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	Dir      string
	NotFound http.Handler
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(s.Dir, r.URL.Path)
	fi, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s.NotFound.ServeHTTP(w, r)
			return
		} else {
			panic(err)
		}
	}
	if fi.IsDir() {
		s.NotFound.ServeHTTP(w, r)
		return
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(w, f)
	if err != nil {
		panic(err)
	}
}
