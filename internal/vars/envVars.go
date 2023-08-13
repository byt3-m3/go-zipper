package vars

import (
	"github.com/byt3-m3/go-zipper/internal/utils/common"
)

var (
	ZipFileLocation        = common.GetEnv("ZIP_FILE_LOCATION", "/uszips.csv")
	WebserverListenAddress = common.GetEnv("LISTEN_ADDRESS", "0.0.0.0:8080")
)
