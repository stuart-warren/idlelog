package env

import (
  "flag"
  "fmt"
  "os"
  "time"
)

const (
  DATE_FORMAT = "2006-01-02"
  TIMESTAMP_FORMAT = time.RFC3339
  OPENTSDB_SOCKET = "127.0.0.1:8953"
  GRAPHITE_SOCKET = "127.0.0.1:2003"
  DEFAULT_IDLE = 5 * time.Minute
  SLEEP_DURATION = 1 * time.Minute
)

var (
  DateString = time.Now().Local().Format(DATE_FORMAT)
  TimestampString = time.Now().Local().Format(TIMESTAMP_FORMAT)
  EpochString = time.Now().Unix()
  TimeForIdle *time.Duration
  SleepDuration *time.Duration
  Debug *bool
  User string
  Host string
  OpenTsdb *bool
  OpenTsdbSocket *string
  Graphite *bool
  GraphiteSocket *string
  LogDir *string
  LogFile string
  LogJsonDir *string
  LogJsonFile string
  LogCsvDir *string
  LogCsvFile string
)

func init() {
  Debug = flag.Bool("d", false, "Debug (stdout)")
  LogDir = flag.String("log", "", "Existing directory to write logs")
  LogJsonDir = flag.String("json", "", "Existing directory to write JSON logs")
  LogCsvDir = flag.String("csv", "", "Existing directory to write CSV logs")
  OpenTsdb = flag.Bool("o", false, "OpenTSDB enabled?")
  OpenTsdbSocket = flag.String("os", OPENTSDB_SOCKET, "UDP Socket for OpenTSDB")
  Graphite = flag.Bool("g", false, "Graphite enabled?")
  GraphiteSocket = flag.String("gs", GRAPHITE_SOCKET, "UDP Socket for Graphite")
  TimeForIdle = flag.Duration("idle", DEFAULT_IDLE, "Length of inactivity considered idle")
  SleepDuration = flag.Duration("poll", SLEEP_DURATION, "Time between checks for activity")
  User = os.Getenv("USER")
  Host, _ = os.Hostname()
  flag.Parse()
  LogFile = fmt.Sprintf("%s/%s.log", *LogDir, DateString)
  LogJsonFile = fmt.Sprintf("%s/%s.json", *LogJsonDir, DateString)
  LogCsvFile = fmt.Sprintf("%s/%s.csv", *LogCsvDir, DateString)
  if len(os.Args) <= 1{
    fmt.Println("No other options selected, so enabling debug mode.")
    *Debug = true
  }
}
