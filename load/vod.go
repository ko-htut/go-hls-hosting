package load

import (
	"github.com/TakuSemba/go-media-hosting/parse"
	"strconv"
	"strings"
)

type VodLoader struct {
	DefaultLoader
	MasterPlaylist parse.MasterPlaylist
}

func NewVodLoader(original parse.MasterPlaylist) VodLoader {
	return VodLoader{
		DefaultLoader:  NewDefaultLoader(original),
		MasterPlaylist: original,
	}
}

func (v *VodLoader) LoadMediaPlaylist(index int) ([]byte, error) {
	var mediaPlaylist []byte
	var tsCount = 0
	for _, tag := range v.MasterPlaylist.MediaPlaylists[index].Tags {
		mediaPlaylist = append(mediaPlaylist, tag...)
		mediaPlaylist = append(mediaPlaylist, '\n')
		if strings.HasPrefix(tag, "#EXTINF") {
			// Consider mp4
			mediaPlaylist = append(mediaPlaylist, strconv.Itoa(tsCount)+".ts"...)
			mediaPlaylist = append(mediaPlaylist, '\n')
			tsCount += 1
		}
	}
	return mediaPlaylist, nil
}
