// Responding with Headers and a Status Code
package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}

	fmt.Printf("got / request. first(%t)=%s, second(%t)=%s, body=%s\n",
		hasFirst,
		first,
		hasSecond,
		second,
		body,
	)
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")

	myName := r.PostFormValue("myName")
	if myName == "" {
		w.Header().Set("x-missing-field", "myName")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", myName))
}

func startServer(ip_port string, mux *http.ServeMux) {
	server := http.Server{
		Addr:    ip_port,
		Handler: mux,
	}

	server.ListenAndServe()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	go startServer(":3333", mux)

	// select {} is used as an infinite loop that is used to keep the program alive and prevent it from exiting until it is manually terminated.
	select {} // same as infinity-for-loop
}

/*
saura@DESKTOP-GC3SDTN MINGW64 ~/OneDrive/Desktop/GO (main)
$ curl -v -X POST 'http://localhost:3333/hello'
*   Trying 127.0.0.1:3333...
* Connected to localhost (127.0.0.1) port 3333 (#0)
> POST /hello HTTP/1.1
> Host: localhost:3333
> User-Agent: curl/7.85.0
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 400 Bad Request
< X-Missing-Field: myName
< Date: Sat, 06 Jan 2024 11:29:57 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact

saura@DESKTOP-GC3SDTN MINGW64 ~/OneDrive/Desktop/GO (main)
$ curl -v -X POST -F 'myName=Super' 'http://localhost:3333/hello'
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1:3333...
* Connected to localhost (127.0.0.1) port 3333 (#0)
> POST /hello HTTP/1.1
> Host: localhost:3333
> User-Agent: curl/7.85.0
> Accept: */*
> Content-Length: 146
> Content-Type: multipart/form-data; boundary=------------------------bd275c1207a738c5   
>
* We are completely uploaded and fine
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Date: Sat, 06 Jan 2024 11:30:26 GMT
< Content-Length: 14
< Content-Type: text/plain; charset=utf-8
<
Hello, Super!
* Connection #0 to host localhost left intact
*/