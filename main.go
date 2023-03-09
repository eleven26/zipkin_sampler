package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/eleven26/zipkin_sampler/sampler"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	cmd := &cobra.Command{
		Use:   "",
		Short: "zipkin sampler",
		Long:  "zipkin sampler",
		Run: func(cmd *cobra.Command, args []string) {
			port, _ := cmd.Flags().GetString("port")
			endpoint, _ := cmd.Flags().GetString("endpoint")

			fmt.Println("listen on port:", port)
			fmt.Println("zipkin endpoint:", endpoint)

			serveHTTP(port, endpoint)
		},
		Args: cobra.ExactArgs(0),
	}

	cmd.Flags().StringP("port", "p", "9422", "zipkin sampler port.")
	cmd.Flags().StringP("endpoint", "e", "http://192.168.2.168:9411/api/v2/spans", "zipkin collector endpoint.")

	//_ = cmd.MarkFlagRequired("endpoint")

	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func serveHTTP(port, endpoint string) {
	reporter := sampler.NewReporter(
		endpoint, // zipkin server
		sampler.NewTimeBaseSampler(time.Second*5), // 采集超过5秒的trace
	)
	collector := sampler.NewCollector(
		time.Minute*10, // 超过 30分钟，没有等到 root trace 就丢弃
		reporter,
	)

	http.HandleFunc("/traces", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic error.")
				log.Error(err)

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write(([]byte)("panic error."))
			}
		}()

		err := collector.Collect(r.Body)
		if err != nil {
			log.Error("collect error.")
			log.Error(err)

			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write(([]byte)("collect error."))

			return
		}

		w.WriteHeader(http.StatusAccepted)
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
