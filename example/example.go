/*
 * Copyright 2012 Derek Buitenhuis
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
    handle, err := mediainfo.Open(args[0])
    if err != nil {
        fmt.Println(err)
        return
    }
    defer mediainfo.Close(handle)

    /* Get the info. */
    val, err := mediainfo.Get(handle, args[1], mediainfo.Video)
    if err != nil {
        fmt.Println(err);
        return
    }

    fmt.Println(val)
}
