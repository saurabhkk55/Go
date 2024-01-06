package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/hello", getHello)

	server := http.Server{
		Addr:    ":3333",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error", err)
	}
}

####################

package main

import (
	"fmt"
	"io"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
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

#######################

// Inspecting a Request’s Query String

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

	fmt.Printf("got / request. first(%t)=%s, second(%t)=%s\n", hasFirst, first, hasSecond, second)
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
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

#######################

// Inspecting a Request’s Query String

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
	io.WriteString(w, "Hello, HTTP!\n")
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

RUN: curl -X POST -d 'This is the body' 'http://localhost:3333?first=1&second='

######################

// Retrieving Form Data

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
	// io.WriteString(w, "Hello, HTTP!\n")

	// $ curl -X POST -F 'myName=Saurabh' 'http://localhost:3333/hello'
	// Hello, Saurabh!
	// $ curl -X POST 'http://localhost:3333/hello'
	// $ curl -X POST -F 'myName=' 'http://localhost:3333/hello'
	// $ curl 'http://localhost:3333/hello'
	// Hello, HTTP!
	myName := r.PostFormValue("myName")

	// $ curl -X POST 'http://localhost:3333/hello?myName=Super'
	// Hello, Super!
	// 	$ curl -X POST 'http://localhost:3333/hello'
	// Hello, HTTP!
	// myName := r.FormValue("myName")
	if myName == "" {
		myName = "HTTP"
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

// The r.PostFormValue method you’re using in the getHello function
// to retrieve the myName form value is a special method that only includes
// values posted from a form in the body of a request. However, an r.FormValue
// method is also available that includes both the form body and any values in
// the query string. So, if you used r.FormValue("myName") you could also remove
// the -F option and include myName=Sammy in the query string to see Sammy returned as well.
// If you did that without the change to r.FormValue, though, you’d see the
// default HTTP response for the name. Being careful about where you’re retrieving these
// values from can avoid potential conflicts in names or bugs that are hard to track down.
// It’s useful to be more strict and use the r.PostFormValue unless you want the flexibility
// to put it in the query string as well.


##########################

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