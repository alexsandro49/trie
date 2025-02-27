package main

import (
	"encoding/json"
	"log"
	"os"
)

type Storage[T any] struct {
	fileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName}
}

func (s *Storage[T]) Save(data *T) {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		log.Fatal("Error: Problem in data processing")
	}

	if os.WriteFile(s.fileName, fileData, 0644) != nil {
		log.Fatal("Error: Unable to save data")
	}
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile(s.fileName)

	if err != nil {
		return err
	}

	if json.Unmarshal(fileData, data) != nil {
		log.Fatal("Error: Unable to read saved data")
	}

	return nil
}
