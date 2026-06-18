package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

type Memory struct {
	mu       sync.Mutex
	Visited  map[string]int
	Research []ResearchRecord
}

func NewMemory() *Memory {

	m := &Memory{
		Visited: make(map[string]int),
	}

	_ = m.LoadFromFile()

	fmt.Println(
		"Loaded records:",
		len(m.Research),
	)

	return m
}

func (m *Memory) Save(topic string, url string, summary string, keywords []string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Visited[topic]++

	m.Research = append(m.Research, ResearchRecord{
		Topic:     topic,
		URL:       url,
		Summary:   summary,
		Keywords:  keywords,
		Timestamp: time.Now().Unix(),
	})

}

func (m *Memory) HasVisited(topic string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.Visited[topic] > 0
}

func (m *Memory) LoadFromFile() error {

	data, err := os.ReadFile("memory.json")

	if err != nil {
		return err
	}

	var records []ResearchRecord

	err = json.Unmarshal(data, &records)

	if err != nil {
		return err
	}

	m.Research = records

	for _, record := range records {
		m.Visited[record.Topic]++
	}

	return nil
}

func (m *Memory) SaveToFile() error {

	m.mu.Lock()
	defer m.mu.Unlock()

	data, err := json.MarshalIndent(
		m.Research,
		"",
		"  ",
	)

	if err != nil {
		return err
	}

	return os.WriteFile(
		"memory.json",
		data,
		0644,
	)
}
