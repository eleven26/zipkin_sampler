package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	_ "zipkin_sampler/sampler"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "9422", "zipkin sampler port.")
	flag.Parse()

	http.HandleFunc("/traces", func(w http.ResponseWriter, r *http.Request) {
		// todo panic 处理
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			logrus.Error("read body error.")
			logrus.Fatal(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(([]byte)("read body error."))

			return
		}

		fmt.Println(string(body))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(([]byte)("ok"))
	})

	logrus.Fatal(http.ListenAndServe(":"+port, nil))
}
