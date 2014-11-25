package main

import (
    "log"
    "fmt"
    "github.com/BurntSushi/xgb"
    "github.com/BurntSushi/xgb/screensaver"
    "github.com/BurntSushi/xgb/xproto"
    "github.com/stuart-warren/idlelog/env"
    "github.com/stuart-warren/idlelog/file"
    "github.com/stuart-warren/idlelog/net"
    "time"
)

var (
  timeSinceInput time.Duration
)

func init() {

}

func debugLog(isIdle bool) {
  if isIdle {
    fmt.Printf("%s %d %s@%s is idle\n", env.TimestampString, env.EpochString, env.User, env.Host)
  } else {
    fmt.Printf("%s %d %s@%s is active\n", env.TimestampString, env.EpochString, env.User, env.Host)
  }
}

func main() {
  x, err := xgb.NewConn()
  if err != nil {
      log.Fatal(err)
  }
  defer x.Close()
  screen := xproto.Setup(x).DefaultScreen(x)

  err = screensaver.Init(x)
  if err != nil {
      log.Fatal(err)
  }

  for true {

    env.DateString = time.Now().Local().Format(env.DATE_FORMAT)
    env.TimestampString = time.Now().Local().Format(env.TIMESTAMP_FORMAT)
    env.EpochString = time.Now().Unix()
    env.LogFile = fmt.Sprintf("%s/%s.log", *env.LogDir, env.DateString)
    env.LogJsonFile = fmt.Sprintf("%s/%s.json", *env.LogJsonDir, env.DateString)
    env.LogCsvFile = fmt.Sprintf("%s/%s.csv", *env.LogCsvDir, env.DateString)

    info, err := screensaver.QueryInfo(x, xproto.Drawable(screen.Root)).Reply()
    if err != nil {
        log.Fatal(err)
    }

    timeSinceInput, _ = time.ParseDuration(fmt.Sprintf("%dms", info.MsSinceUserInput))

    var isIdle bool
    if timeSinceInput > *env.TimeForIdle {
      isIdle = true
    }

    // call reporting methods here
    if *env.Debug {
      debugLog(isIdle)
    }
    if *env.LogDir != "" {
      file.WriteLog(isIdle)
    }
    if *env.LogJsonDir != "" {
      file.WriteJsonLog(isIdle)
    }
    if *env.LogCsvDir != "" {
      file.WriteCsvLog(isIdle)
    }
    if *env.Graphite {
      net.SendGraphite(isIdle)
    }
    if *env.OpenTsdb {
      net.SendOpenTsdb(isIdle)
    }

    time.Sleep(*env.SleepDuration)
  }
}
