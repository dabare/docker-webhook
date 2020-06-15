package main

import "github.com/dabare/docker-webhook/pkg/webhooks"

func main() {
	r := webhooks.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
