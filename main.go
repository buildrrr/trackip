// package main

// import (
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"strings"
// )

// func hello(w http.ResponseWriter, r *http.Request) {

// 	fmt.Fprint(w, "Hello Boss")

// }

// func getClientIP(r *http.Request) string {
// 	// Check X-Forwarded-For header first (standard for proxies)
// 	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
// 		// The first IP in the list is the original client
// 		parts := strings.Split(forwarded, ",")
// 		return strings.TrimSpace(parts[0])
// 	}

// 	// Check X-Real-IP header (common alternative)
// 	if realIP := r.Header.Get("X-Real-IP"); realIP != "" {
// 		return realIP
// 	}

// 	// Fallback to RemoteAddr, stripping the port
// 	ip, _, err := net.SplitHostPort(r.RemoteAddr)
// 	if err != nil {
// 		return r.RemoteAddr
// 	}
// 	return ip
// }

// func myip(w http.ResponseWriter, r *http.Request) {
// 	// 1. Get IP from the incoming request (Local/Proxy view)
// 	reqIPStr := getClientIP(r)

// 	reqIP := net.ParseIP(reqIPStr)
// 	reqVersion := "Unknown"
// 	if reqIP != nil {
// 		if reqIP.To4() != nil {
// 			reqVersion = "IPv4"
// 		} else {
// 			reqVersion = "IPv6"
// 		}
// 	}

// 	fmt.Fprintf(w, "My IP Address is:\n\n%s: %s\n", reqVersion, reqIPStr)
// }

// func inspect(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "--- Request Details ---")
// 	fmt.Fprintf(w, "Method: %s\n", r.Method)
// 	fmt.Fprintf(w, "URL: %s\n", r.URL.String())
// 	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
// 	fmt.Fprintf(w, "Host: %s\n", r.Host)
// 	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)

// 	fmt.Fprintln(w, "\n--- Headers ---")
// 	for name, values := range r.Header {
// 		for _, value := range values {
// 			fmt.Fprintf(w, "%s: %s\n", name, value)
// 		}
// 	}

// 	fmt.Fprintln(w, "\n--- Specific Interesting Headers ---")
// 	fmt.Fprintf(w, "User-Agent: %s\n", r.Header.Get("User-Agent"))
// 	fmt.Fprintf(w, "Accept-Language: %s\n", r.Header.Get("Accept-Language"))
// }

// func main() {

// 	http.HandleFunc("/h", hello)
// 	http.HandleFunc("/myip", myip)
// 	http.HandleFunc("/inspect", inspect)

// 	http.ListenAndServe(":8080", nil)

/// }

package main

import (
	"fmt"
	"net/http"
)

func myip(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "The method used is: %s\n", r.Method)
	fmt.Fprintf(w, "Url requested is:%s\n", r.URL.Path)
	fmt.Fprintf(w, "The protocol used is:%s\n", r.Proto)
	fmt.Fprintf(w, "The ip address used is:%s\n", r.RemoteAddr)
	fmt.Fprintf(w, "The hostname used is:%s\n", r.Host)
	fmt.Fprintf(w, "The url query is:%s\n", r.URL.Query())
	fmt.Fprintf(w, "The header is:%s\n", r.Header)
	fmt.Fprintf(w, "The body is:%s\n", r.Body)
	fmt.Fprintf(w, "The context is:%s\n", r.Context())
}

func main() {

	http.HandleFunc("/", myip)
	http.ListenAndServe(":8080", nil)

}
