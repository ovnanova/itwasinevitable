package main

import (
	"os"
	"path/filepath"
)

func cleanup() {
	// Enable only conversation in the log.
	if err := os.WriteFile("df_linux/data/init/announcements.txt", announcementstxt, 0644); err != nil {
		panic(err)
	}

	// Remove any old log we might have.
	if err := os.Remove("df_linux/gamelog.txt"); err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	// Remove any old worlds.
	if err := clearSaves(); err != nil {
		panic(err)
	}

	// Set up df-ai.
	if err := os.WriteFile("df_linux/dfhack-config/df-ai.json", dfaijson, 0644); err != nil {
		panic(err)
	}
}

func clearSaves() error {
	files, err := os.ReadDir("df_linux/data/save")
	if err != nil {
		return err
	}

	for _, file := range files {
		if err = os.RemoveAll(filepath.Join("df_linux/data/save", file.Name())); err != nil && !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}
