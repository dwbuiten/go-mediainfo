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

package mediainfo

/*
 #cgo LDFLAGS: -ldl
 #include <stdlib.h>
 #include "c/mediainfo_wrapper.c"
*/
import "C"

import (
    "errors"
    "unsafe"
)

const (
    General = C.MediaInfo_Stream_General
    Video   = C.MediaInfo_Stream_Video
    Audio   = C.MediaInfo_Stream_Audio
    Image   = C.MediaInfo_Stream_Image
)

/* Dont expose this ugliness. */
type MediaInfo struct {
    ptr unsafe.Pointer
}

/* Loas the shared library. */
func Init() {
    C.mediainfo_c_init()
}

/*
 * Opens and parses the file.
 *
 * Takes a full path or reltaive path as an argument,
 * and returns a MediaInfo handler.
 */
func Open(file string) (MediaInfo, error) {
    var ret MediaInfo

    cfile := C.CString(file)
    defer C.free(unsafe.Pointer(cfile))

    cptr   := C.mediainfo_c_open(cfile)
    ret.ptr = cptr
    if cptr == nil {
        return ret, errors.New("Cannot open file.")
    }

    return ret, nil
}

/*
 * Get audio or video info for a key.
 *
 * Matches up with the list available via:
 *     mediainfo --Info-Parameters
 *
 * Only handles one video and audio stream currently.
 *
 * Takes a key, and a stream type as an arguments. Valid
 * stream types are mediainfo.Video and mediainfo.Audio.
 */
func (handle MediaInfo) Get(key string, typ uint32) (string, error) {
    ckey  := C.CString(key)
    cptr  := unsafe.Pointer(handle.ptr)
    defer C.free(unsafe.Pointer(ckey))

    cret := C.mediainfo_c_get(cptr, ckey, typ)
    ret  := C.GoString(cret)
    if len(ret) == 0 {
        return "", errors.New("Cannot get value for key.")
    }

    return ret, nil
}

/* Close a handle. */
func (handle MediaInfo) Close() {
    cptr  := unsafe.Pointer(handle.ptr)

    C.mediainfo_c_close(cptr)
}
