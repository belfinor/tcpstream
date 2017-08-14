package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2017-08-13


import (
  "flag"
  "github.com/belfinor/Helium/daemon"
  "github.com/belfinor/Helium/log"
)


var STREAM *Stream


func main() {

    conf   := ""
    is_daemon := false

    flag.StringVar( &conf, "c", "lnode.json", "config file name" )
    flag.BoolVar( &is_daemon, "d", false, "run as daemon" )

    flag.Parse()

    cfg := LoadConfig( conf )

    if is_daemon {
        daemon.Run( &cfg.Daemon )
    }

    log.Init( &cfg.Log )

    if is_daemon {
        log.Info( "start application as daemon" )
    } else {
        log.Info( "start application" )
    }

    STREAM = NewStream( cfg.Proxy )

    srv := &Server{
        Addr:     cfg.Listen,
    }

    srv.Start()
}

