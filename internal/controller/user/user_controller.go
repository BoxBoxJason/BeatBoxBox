package user_controller

import (
	"BeatBoxBox/pkg/utils"
	"path/filepath"
)

func init() {
	go utils.CheckDirExists(filepath.Join("data", "illustrations", "users"))
}
