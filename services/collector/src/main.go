// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ToDo : add code from https://github.com/privacysandbox/aggregation-service/blob/main/docs/collecting.md

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Print("starting server...")
	http.HandleFunc("/", handler)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	host := os.Getenv("COLLECTOR_HOST")
	if host == "" {
		host = "nowhere"
	}
	fmt.Fprintf(w, "Hello from %s !\n", host)

	detail := os.Getenv("COLLECTOR_DETAIL")
	if detail != "" {
		fmt.Fprintf(w, "My job is to %s\n", detail)
		fmt.Fprintf(w, "Looking forward to working with you !\n")
	}
}
