package main

import (
    "fmt"
    "encoding/json"
    "bytes"
)

type SeedInfo struct {
    fpeKey string
    tweak string
    sessionKey string
    verifyPin string
    rate int
    ts int
    seed int
    qrCode string 
}  

func main(){
    s := SeedInfo{fpeKey: "zanpPlQmig7I1CqWTsKHWw==",tweak: "7DE78013B7DD0884", sessionKey: "DqI8WknxDLM+Xr4ZtEoZgQ==", verifyPin: "bXQyMDE4MDExMTE2MTgxNA==", rate: 5, ts:1515658859, seed: 12704240038, qrCode: "6210391175135012924"}
    b := new(bytes.Buffer)
    json.NewEncoder(b).Encode(s)
    fmt.Println(b)
}
