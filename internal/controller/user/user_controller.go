package user_controller

import (
	file_utils "BeatBoxBox/pkg/utils/fileutils"
	"path/filepath"
)

func init() {
	go file_utils.CheckDirExists(filepath.Join("data", "illustrations", "users"))
}
