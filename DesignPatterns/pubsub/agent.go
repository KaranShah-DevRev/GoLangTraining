package main

import (
	"fmt"
	"sync"
)

// Agent represents a basic publish-subscribe agent.
type Agent struct {
	mu          sync.Mutex               // Mutex for synchronization
	subscribers map[string][]chan string // Map to store subscribers for different topics
	quit        chan struct{}            // Channel to signal when the agent is closed
	closed      bool                     // Boolean flag to track whether the agent is closed
}

// NewAgent initializes and returns a new Agent instance.
func NewAgent() *Agent {
	return &Agent{
		subscribers: make(map[string][]chan string),
		quit:        make(chan struct{}),
	}
}

// Publish publishes a message to all subscribers of a given topic.
func (a *Agent) Publish(topic string, msg string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return
	}
	for _, ch := range a.subscribers[topic] {
		// Launch a goroutine to send the message to the subscriber
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}
}

// Subscribe subscribes to a given topic and returns a channel for receiving messages on that topic.
func (a *Agent) Subscribe(topic string) <-chan string {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return nil
	}
	ch := make(chan string, 1)
	a.subscribers[topic] = append(a.subscribers[topic], ch)
	return ch
}

// Close closes the agent, signaling to subscribers that they should stop listening.
func (a *Agent) Close() {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.closed {
		return
	}
	a.closed = true
	close(a.quit)
	// Close all subscriber channels
	for _, subs := range a.subscribers {
		for _, ch := range subs {
			close(ch)
		}
	}
}

func main() {
	// Create a new agent
	agent := NewAgent()
	// Defer closing the agent when the main function exits
	defer agent.Close()
	// Subscribe to "topic1"
	subscribe := agent.Subscribe("topic1")
	// Publish a message to "topic1"
	agent.Publish("topic1", "hello")
	// Receive and print the message from the subscriber
	fmt.Println(<-subscribe)
}
