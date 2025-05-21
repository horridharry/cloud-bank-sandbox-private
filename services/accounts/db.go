package main

import (
    "sync"
)

var (
    accounts = make(map[string]Account)
    mu       = sync.Mutex{}
)
