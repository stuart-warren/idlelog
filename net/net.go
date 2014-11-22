package net

import (
  "fmt"
  "github.com/stuart-warren/idlelog/env"
  "log"
  "net"
)

func SendOpenTsdb(isIdle bool) {
  if isIdle {
    send("udp", env.OpenTsdbSocket, fmt.Sprintf("put user.active %d 0 host=%s user=%s\n", env.EpochString, env.Host, env.User))
  } else {
    send("udp", env.OpenTsdbSocket, fmt.Sprintf("put user.active %d 1 host=%s user=%s\n", env.EpochString, env.Host, env.User))
  }
}

func SendGraphite(isIdle bool) {
  if isIdle {
    send("udp", env.GraphiteSocket, fmt.Sprintf("user.active.%s.%s 0 %d\n", env.Host, env.User, env.EpochString))
  } else {
    send("udp", env.GraphiteSocket, fmt.Sprintf("user.active.%s.%s 1 %d\n", env.Host, env.User, env.EpochString))
  }
}

func send(connType string, socket *string, msg string) {
  conn, err := net.Dial(connType, *socket)
  if err != nil {
  	log.Fatalf("Could not connect to %s socket %s", connType, socket)
  }
  defer conn.Close()
  _, err = fmt.Fprintf(conn, msg)
  if err != nil {
    log.Fatal("Error sending message")
  }
}
