package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"strconv"
	"time"
)

func main() {

	// Get if BODY is to be dumped from ENV
	dumpBody := true
	dumpEnv := os.Getenv("BODY")
	if dumpEnv == "0" {
		dumpBody = false
	}

	port := 80
	// Get port from ENV
	rawPort := os.Getenv("PORT")
	/* Further controls on port not needed
	http.ListenAndServe already does them */
	if rawPort != "" {
		var err error
		port, err = strconv.Atoi(rawPort)
		if err != nil {
			log.Fatalln("Invalid PORT specified")
		}
	}

	// Routes
	http.HandleFunc("/", mirror(MirrorHandler(), dumpBody))

	// Start the HTTP Server
	fmt.Fprintf(os.Stdout, "Listening on :%d\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}

// MirrorHandler handles the / path and returns Req+Resp in resp's body
func MirrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "", "")
	})
}

// Mirror Captures req and resp and mirror them back to the client
func mirror(h http.Handler, dumpBody bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Container for the Dumped request and response
		var wholeDump string

		// Save a copy of the request for debugging
		requestDump, err := httputil.DumpRequest(r, dumpBody)
		if err != nil {
			log.Println(err)
		}

		wholeDump = fmt.Sprintf("%s \n", "====================\n\n")

		// Add Client IP
		clientIp := r.RemoteAddr
		wholeDump += fmt.Sprintf("CLIENT: %q REQUEST \n", clientIp)
		wholeDump += fmt.Sprintf("%s\n", "--------------------------")
		// Format the Unixtimestamp into a string
		timeReq := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		wholeDump += fmt.Sprintf("%s \n", "Time: "+timeReq)
		wholeDump += fmt.Sprintf("%s \n", string(requestDump))

		// add extra br if there is also the body
		if dumpBody {
			wholeDump += fmt.Sprintf("\n")
		}

		// Setup response capturing
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, r)

		// Save a copy of the response for debugging
		dump, err := httputil.DumpResponse(rec.Result(), false)
		if err != nil {
			log.Fatal(err)
		}

		// Add hostname information into the dump
		hostname, _ := os.Hostname()
		wholeDump += fmt.Sprintf("SERVER: %s RESPONSE \n", hostname)
		wholeDump += fmt.Sprintf("%s\n", "--------------------------")
		// Format the Unixtimestamp into a string
		timeRes := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		wholeDump += fmt.Sprintf("%s \n", "Time: "+timeRes)
		wholeDump += fmt.Sprintf("%s \n", string(dump))

		wholeDump += fmt.Sprintf("%s", "====================")

		// Copy the captured response headers to the new response
		for k, v := range rec.Header() {
			w.Header()[k] = v
		}

		// Logs the req+resp on STDOUT
		fmt.Println(wholeDump)

		// Writes the req+resp in body of resp
		fmt.Fprintln(w, wholeDump)
	}
}
