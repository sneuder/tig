package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func CreateJSONFile(data interface{}, filename string, path string) error {
	// Marshal the struct to JSON with indentation for readability
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %w", err)
	}

	// Create or open the specified JSON file
	fullPathFile := filepath.Join(path, filename+".json")
	file, err := os.Create(fullPathFile)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	// Write the JSON data to the file
	if _, err = file.Write(jsonData); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func AddExtension(fileName string) string {
	return fileName + ".json"
}

func ReadJSONFile(fileName string, path string) ([]byte, error) {
	fullPathFile := filepath.Join(path, AddExtension(fileName))
	data, err := os.ReadFile(fullPathFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return data, nil
}

func RemoveJSONFile(fileName string, path string) bool {
	removed := true

	fullPathFile := filepath.Join(path, AddExtension(fileName))
	exists := CheckJSONFile(fileName, path)

	if exists {
		err := os.Remove(fullPathFile)
		removed = err != nil
	}

	return removed
}

func CheckJSONFile(filename string, path string) bool {
	fullPathFile := filepath.Join(path, AddExtension(filename))
	if _, err := os.Stat(fullPathFile); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	return true
}
