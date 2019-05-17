package danmu_rest 

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

//Logger ...log out web requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, err := os.OpenFile("rest.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		start := time.Now()

		inner.ServeHTTP(w, r)

		log.SetOutput(file)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
