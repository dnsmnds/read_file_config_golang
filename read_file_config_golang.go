// Read file config - Go

// Copyright 2014 - Dênis Mendes. All rights reserved.
// Author: Dênis Mendes <denisffmendes@gmail.com>
// Use of this source code is governed by a BSD-style

package main

import (
    "fmt"
    "log"
    "github.com/alyu/configparser"
)

func main(){
  config, err := configparser.Read("config.ini")
  if err != nil {
    log.Fatal(err)
  }

  section, err := config.Section("mysql")

  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Printf("%s \n", section.ValueOf("database"))
  }
}
