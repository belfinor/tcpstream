package main


// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2017-08-12


import (
  "encoding/json"
  "fmt"
  "github.com/belfinor/Helium/log"
  "github.com/belfinor/Helium/daemon"
  "io/ioutil"
  "os"
)


type Config struct {
  Daemon daemon.Config `json:"daemon"`
  Listen string `json:"listen"`
  Log    log.Config `json:"log"`
  Proxy  []string `json:"proxy"`
}

var _config *Config


func LoadConfig( filename string ) *Config {

    var data []byte
    var err error

    if data, err = ioutil.ReadFile( filename ) ; err != nil {
        fmt.Println( "read " + filename + " error" )
        os.Exit(1)
    }

    var _con Config

    if err = json.Unmarshal( data, &_con ) ; err != nil {
        fmt.Println( err )
        os.Exit(1)
    }

    _config = &_con

    return &_con
}


func SetConfig( conf *Config ) {
    _config = conf
}


func GetConfig() *Config {
    return _config
}


