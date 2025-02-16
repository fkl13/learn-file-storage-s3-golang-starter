package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0755)
	}
	return nil
}

func buildFileName(videoID uuid.UUID, fileExtension string) string {
	return fmt.Sprintf("%s%s", videoID, fileExtension)
}

func (cfg apiConfig) buildAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(filePath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, filePath)
}
