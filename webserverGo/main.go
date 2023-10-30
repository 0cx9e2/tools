package main


import (
  "fmt"
  "log"
  "net"
  "net/http"
)


func main() {

    http.HandleFunc("/", func (w http.ResponseWriter, 
        r *http.Request) {
  
        // when receive the request, print the greeting meassage
        fmt.Fprint(w, "webserverGo by 0xc9e2")
  
      })

  log.Println("Starting web server")

  listener, err := net.Listen("tcp", ":8080")

  if err != nil {
    log.Fatal(err)
  }

  http.Serve(listener, nil)

}
