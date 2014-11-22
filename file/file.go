package file

import (
  "encoding/json"
  "fmt"
  "github.com/stuart-warren/idlelog/env"
  "log"
  "os"
)

type Log struct {
  Timestamp string `json:"@timestamp"`
  User string `json:"user"`
  Host string `json:"host"`
  Active int `json:"isActive"`
}

func WriteLog(isIdle bool) {
  if isIdle {
    write(env.LogFile, fmt.Sprintf("%s %s@%s is idle", env.TimestampString, env.User, env.Host))
  } else {
    write(env.LogFile, fmt.Sprintf("%s %s@%s is active", env.TimestampString, env.User, env.Host))
  }
}

func WriteCsvLog(isIdle bool) {
  if isIdle {
    write(env.LogCsvFile, fmt.Sprintf("%s,%s,%s,0", env.TimestampString, env.User, env.Host))
  } else {
    write(env.LogCsvFile, fmt.Sprintf("%s,%s,%s,1", env.TimestampString, env.User, env.Host))
  }
}

func WriteJsonLog(isIdle bool) {
  l := &Log{
    Timestamp: env.TimestampString,
    User: env.User,
    Host: env.Host,
  }
  if isIdle {
    l.Active = 0
  } else {
    l.Active = 1
  }
  msg, _ := json.Marshal(l)
  write(env.LogJsonFile, string(msg))
}

func write(file string, msg string) {
  f, err := os.OpenFile(file, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    log.Fatalf("error opening file: %v", err)
  }
  defer f.Close()

  log.SetOutput(f)
  log.SetFlags(0)
  log.Println(msg)
}
