package student

import "net/http"

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Here are all the students data"))
		w.WriteHeader(http.StatusOK)
	}
}
