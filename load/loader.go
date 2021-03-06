package load

import (
	"github.com/TakuSemba/go-hls-hosting/parse"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type Loader interface {
	LoadMasterPlaylist() ([]byte, error)
	LoadMediaPlaylist(index int) ([]byte, error)
	LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error)
}

type DefaultLoader struct {
	MasterPlaylist parse.MasterPlaylist
}

func NewDefaultLoader(original parse.MasterPlaylist) DefaultLoader {
	return DefaultLoader{MasterPlaylist: original}
}

func (v *DefaultLoader) LoadMasterPlaylist() ([]byte, error) {
	var masterPlaylist []byte
	var mediaPlaylistIndex = 0
	for _, tag := range v.MasterPlaylist.Tags {
		masterPlaylist = append(masterPlaylist, tag...)
		masterPlaylist = append(masterPlaylist, '\n')
		if strings.HasPrefix(tag, "#EXT-X-STREAM-INF") {
			masterPlaylist = append(masterPlaylist, strconv.Itoa(mediaPlaylistIndex)+"/playlist.m3u8"...)
			masterPlaylist = append(masterPlaylist, '\n')
			mediaPlaylistIndex += 1
		}
	}
	return masterPlaylist, nil
}
func (v *DefaultLoader) LoadSegment(mediaPlaylistIndex, segmentIndex int) ([]byte, error) {
	mediaPlaylistPath := v.MasterPlaylist.MediaPlaylists[mediaPlaylistIndex].Path
	segmentPath := v.MasterPlaylist.MediaPlaylists[mediaPlaylistIndex].Segments[segmentIndex].Path
	segment, err := ioutil.ReadFile(filepath.Join(filepath.Dir(mediaPlaylistPath), segmentPath))
	if err != nil {
		return []byte{}, nil
	}
	return segment, nil
}
