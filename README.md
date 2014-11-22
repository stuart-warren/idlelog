idlelog
=======

Quick tool to record idleness - in progress (Linux only - Requires X11)

Usage:
------
```
$ idlelog -h
Usage of idlelog:
  -csv="": Existing directory to write CSV logs
  -d=false: Debug (stdout)
  -g=false: Graphite enabled?
  -gs="127.0.0.1:2003": UDP Socket for Graphite
  -idle=5m0s: Length of inactivity considered idle
  -json="": Existing directory to write JSON logs
  -log="": Existing directory to write logs
  -o=false: OpenTSDB enabled?
  -os="127.0.0.1:8953": UDP Socket for OpenTSDB
  -poll=1m0s: Time between checks for activity
```

Try sending metrics to services or logging to a file

Example:
--------

Send metrics to local OpenTDSB tcollector and log to json files in /home/user/activitylog
```
$ mkdir -p /home/user/activitylog
$ idlelog -o -json /home/user/activitylog &
```
