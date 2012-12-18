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

/*
 * Go doesn't like the function pointers, like in
 * MediaInfo's headers, so we need to provide a
 * tiny wrapper.
 * */

#include <MediaInfoDLL/MediaInfoDLL.h>

#include "mediainfo_go.h"

void mediainfo_c_init()
{
    MediaInfoDLL_Load();
}

void *mediainfo_c_open(char *filename)
{
    void *handle;
    int mret;

    handle = MediaInfo_New();
    if (!handle)
        return NULL;

    mret = MediaInfo_Open(handle, filename);
    if (!mret)
        return NULL;

    return handle;
}

char *mediainfo_c_get(void *opaque, char *key, enum MediaInfo_stream_t type)
{
    return (char *) MediaInfo_Get(opaque, type, 0, key, MediaInfo_Info_Text,
                                  MediaInfo_Info_Name);
}

void mediainfo_c_close(void *opaque)
{
    MediaInfo_Close(opaque);
}
