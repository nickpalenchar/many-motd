package parsers

import (
	"log"
	"many-motd/config"
	"many-motd/data"
	"os"
	"path"
	"time"
)

const (
	datasubdir     string = "remote"
	remoteDataFile string = "remotedata.json"
)

type RemoteEntry struct {
	URL         string `json:"url"`
	LocalFile   string `json:"localFile"`
	LastUpdated int64  `json:"lastUpdated"`
}

type RemoteData []RemoteEntry

func parseRemote(url string) []data.Message {
	initRemoteData()
	return make([]data.Message, 0)
}

func initRemoteData() {
	parserDatadir := path.Join(config.GetConfig().Src, config.Datadir, datasubdir)

	if _, err := os.Stat(parserDatadir); os.IsNotExist(err) {
		os.MkdirAll(parserDatadir, os.FileMode(0777))
	} else if err != nil {
		log.Fatal("could not setup remote data")
		return
	}

	// check if remotedata.json exists
	if _, err := os.Stat(path.Join(parserDatadir, remoteDataFile)); os.IsNotExist(err) {
		err := os.WriteFile(
			path.Join(parserDatadir, remoteDataFile),
			[]byte("{}"),
			0777)
		if err != nil {
			log.Fatalf("Could not create data files for remote parsing: %v", err)
		}
	}
}

func checkUrlUpdates(re RemoteEntry) {
	now := time.Now().UTC().UnixMilli()

	if now-re.LastUpdated > int64(time.Hour*24) {
		
	}
}
