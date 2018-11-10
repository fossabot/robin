# robin
[![Go Report Card](https://goreportcard.com/badge/github.com/jiansoft/robin)](https://goreportcard.com/report/github.com/jiansoft/robin)
[![Build Status](https://travis-ci.org/jiansoft/robin.svg?branch=master)](https://travis-ci.org/jiansoft/robin)
### Features

Fiber  
-------
* GoroutineSingle - a fiber backed by a dedicated goroutine. Every job is executed by a goroutine.
* GoroutineMulti - a fiber backed by more goroutine. Each job is executed by a new goroutine.

Channels
-------
* Channels callback is executed for each message received.

Cron
-------
Golang job scheduling for humans. It is inspired by [schedule](<https://github.com/dbader/schedule>).
  


Usage
================

### Quick Start

1.Instsall
~~~
  go get github.com/jiansoft/robin
~~~

2.Use examples
~~~ golang
import (
    "log"
    "time"
    
    "github.com/jiansoft/robin"
)

func main() {
    //The method is going to execute only once after 2000 ms.
    robin.Delay(2000).Do(runCron, "a Delay 2000 ms")
    
    minute := 11
    second := 50
    
    //Every Friday is going to execute once at 14:11:50 (HH:mm:ss).
    robin.EveryFriday().At(14, minute, second).Do(runCron, "Friday")

    //Every N day  is going to execute once at 14:11:50(HH:mm:ss)
    robin.Every(1).Days().At(14, minute, second).Do(runCron, "Days")

    //Every N hours is going to execute once at 11:50:00(HH:mm:ss).
    robin.Every(1).Hours().At(0, minute, second).Do(runCron, "Every 1 Hours")

    //Every N minutes is going to execute once at 50(ss).
    robin.Every(1).Minutes().At(0, 0, second).Do(runCron, "Every 1 Minutes")

    //Every N seconds is going to execute once
    robin.Every(10).Seconds().Do(runCron, "Every 10 Seconds")
}

func runCron(s string) {
    log.Printf("I am %s CronTest %v\n", s, time.Now())
}
~~~
[More example](<https://github.com/jiansoft/robin/blob/master/example/main.go>)

## License

Copyright (c) 2017

Released under the MIT license:

- [www.opensource.org/licenses/MIT](http://www.opensource.org/licenses/MIT)
