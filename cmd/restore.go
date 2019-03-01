package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type fileNames []string

func (f fileNames) contains(file string) bool {
	for _, v := range f {
		if file == v {
			return true
		}
	}
	return false
}

// A file of the same name exists in the current directory.
func fileExistsCurrentDir(name string) (bool, error) {
	wd, err := os.Getwd()
	if err != nil {
		return true, err
	}
	filesInCurrentDir, err := getFileNames(wd)
	if err != nil {
		return true, err
	}
	if fileNames(filesInCurrentDir).contains(name) {
		return true, nil
	}
	return false, nil
}

// Remove a character string what given when moving to the trash can.
func removeAffix(org string) string {
	return org[:strings.LastIndex(org, "_")] + getExt(org)
}

// Choose the file to restore.
func chooseTarget(trashCanPath string) (string, string, error) {
	for {
		date, err := chooseFile(trashCanPath)
		if err != nil {
			return "", "", err
		} else if date == "" {
			return "", "", fmt.Errorf("Cannot get date")
		}

		fileName, err := chooseFile(filepath.Join(trashCanPath, date))
		if err != nil {
			return "", "", err
		} else if fileName != "" {
			return date, fileName, nil
		}
	}
}

func getTarget() (string, string, error) {
	trashCanPath := getTrashCanPath()
	date, fileName, err := chooseTarget(trashCanPath)
	if err != nil {
		return "", "", err
	}

	oldFilePath := filepath.Join(trashCanPath, date, fileName)
	if _, err := os.Stat(oldFilePath); err != nil {
		return "", "", err
	}

	return oldFilePath, removeAffix(fileName), nil
}

// Restore chose file or directory.
func restore(_ *cobra.Command, _ []string) error {
	oldFilePath, newFilePath, err := getTarget()
	if err != nil {
		return err
	}

	if exists, err := fileExistsCurrentDir(newFilePath); err != nil {
		return err
	} else if exists {
		return fmt.Errorf("A file with the same name already exists.")
	}

	if err := os.Rename(oldFilePath, newFilePath); err != nil {
		return err
	}

	return nil
}

func restoreCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "restore",
		Short: "Move files in the trash can to the current directory",
		RunE:  restore,
	}

	return cmd
}
