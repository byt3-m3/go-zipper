package vars

import (
	"go-zip/internal/utils/common"
)

var (
	ZipFileLocation        = common.GetEnv("ZIP_FILE_LOCATION", "/app/uszips.csv")
	WebserverListenAddress = common.GetEnv("LISTEN_ADDRESS", "0.0.0.0:8080")
)
