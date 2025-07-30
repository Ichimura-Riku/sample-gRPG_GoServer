package main

import (
	"fmt"
	"log"
	"net/http"

	"buf-demo/example/client/samplefunc"
)

func main() {
	http.HandleFunc("/call-grpc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "gRPCリクエストを送信します...")
		samplefunc.CallGRPC()
	})
	log.Println("HTTPサーバー起動: :18080 (curlで /call-grpc を叩いてください)")
	http.ListenAndServe(":18080", nil)
}
