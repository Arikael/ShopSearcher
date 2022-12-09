package search_config

import (
	"embed"
	"log"
	"os"
)

//go:embed data-embedded/*.txt
var embeddedDataFolder embed.FS

const embeddedFolderName = "data-embedded"
const localFolderName = "./data"

func fsExists(name string) (bool, error) {
	var _, err = os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			log.Fatal(err)
			return false, err
		}
	}

	return true, nil
}
