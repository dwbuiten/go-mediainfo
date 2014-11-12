package main

import (
    "flag"
    "fmt"
    "mediainfo"
)

func main() {

    flag.Parse()
    args := flag.Args()
    if len(args) < 2 {
        fmt.Println("Not enough arguments.")
        return
    }

    /* Load the shared library. */
    mediainfo.Init()

    /* Open and parse the file. */
    info, err := mediainfo.Open(args[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer info.Close()

    /* Get the info. */
    val, err := info.Get(args[1], 0, mediainfo.Video)
    if err != nil {
        fmt.Println(err);
        return
    }

    fmt.Println(val)
}
