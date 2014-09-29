/*
 * Go doesn't like the function pointers, like in
 * MediaInfo's headers, so we need to provide a
 * tiny wrapper.
 * */

#include <MediaInfoDLL/MediaInfoDLL.h>

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
