package main

import (
	"fmt"
	"time"
)

type Agent struct {
	rootTopic string
	current   string
	memory    *Memory
	running   bool
	stopChan  chan struct{}
}

func NewAgent(topic string) *Agent {
	return &Agent{
		rootTopic: topic,
		current:   topic,
		memory:    NewMemory(),
		stopChan:  make(chan struct{}),
	}
}

func (a *Agent) Run() {
	a.running = true

	go func() {

		for {

			select {

			case <-a.stopChan:
				fmt.Println("Agent stopped")
				return

			default:

				fmt.Println("\n=================================")
				fmt.Println("Researching:", a.current)

				urls, err := Search(a.current)

				if err != nil {
					fmt.Println(err)
					time.Sleep(10 * time.Second)
					continue
				}

				var collected []string
				var lastURL string
				var lastSummary string

				for _, url := range urls {

					fmt.Println("Scraping:", url)

					text, err := Scrape(url)

					if err != nil {
						continue
					}

					summary := text

					if len(summary) > 500 {
						summary = summary[:500]
					}

					lastURL = url
					lastSummary = summary

					keywords := ExtractKeywords(text)
					collected = append(collected, keywords...)
				}

				a.memory.Save(
					a.current,
					lastURL,
					lastSummary,
					collected,
				)

				if err := a.memory.SaveToFile(); err != nil {
					fmt.Println("Failed to save memory:", err)
				}

				next := PickNextTopic(collected, a.memory)

				if next != "" {
					a.current = next
				}

				fmt.Println("Next topic:", a.current)

				time.Sleep(15 * time.Second)
			}
		}
	}()
}

func (a *Agent) Stop() {
	close(a.stopChan)
}
