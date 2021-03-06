package core

import (
	"sync"

	sherlock "github.com/kcmerrill/sherlock/core"
)

// Q holds all of the topics along with a lock to access said topics
type Q struct {
	lock   *sync.Mutex
	topics map[string]*Topic
	Stat   *sherlock.Sherlock
}

// CreateQ inits our Q
func CreateQ() *Q {
	q := &Q{
		Stat: sherlock.New(100),
	}
	q.topics = make(map[string]*Topic)
	q.lock = &sync.Mutex{}
	return q
}

// NewMessage creates a new message based on topic key and value
func (q *Q) NewMessage(topic, key, value string) *Message {
	q.lock.Lock()
	var m *Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].NewMessage(topic, key, value)
	} else {
		q.topics[topic] = CreateTopic(topic, q)
		m = q.topics[topic].NewMessage(topic, key, value)
	}
	q.lock.Unlock()
	return m
}

// NewRawMessage creates a new message based on a raw message
func (q *Q) NewRawMessage(msg *Message) *Message {
	q.lock.Lock()
	var m *Message
	if _, exists := q.topics[msg.Topic]; exists {
		m = q.topics[msg.Topic].NewRawMessage(msg)
	} else {
		q.topics[msg.Topic] = CreateTopic(msg.Topic, q)
		m = q.topics[msg.Topic].NewRawMessage(msg)
	}
	q.lock.Unlock()
	return m
}

// Message returns a message from a given topic
func (q *Q) Message(topic string) *Message {
	q.lock.Lock()
	var m *Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].Message()
	} else {
		q.topics[topic] = CreateTopic(topic, q)
		m = q.topics[topic].Message()
	}
	q.lock.Unlock()
	return m
}

// Messages returns multiple messages from a given topic
func (q *Q) Messages(topic string, count int) []*Message {
	q.lock.Lock()
	var m []*Message
	if _, exists := q.topics[topic]; exists {
		m = q.topics[topic].Messages(count)
	} else {
		q.topics[topic] = CreateTopic(topic, q)
		m = q.topics[topic].Messages(count)
	}
	q.lock.Unlock()
	return m
}

// Complete finishes the work for a given message id
func (q *Q) Complete(topic, id string) {
	q.lock.Lock()
	if _, exists := q.topics[topic]; exists {
		q.topics[topic].CompleteMessage(topic, id)
	} else {
		q.topics[topic] = CreateTopic(topic, q)
		q.topics[topic].CompleteMessage(topic, id)
	}
	q.lock.Unlock()
}

// Delete a given message off of a topic
func (q *Q) Delete(topic, id string) {
	q.lock.Lock()
	if _, exists := q.topics[topic]; exists {
		q.topics[topic].DeleteMessage(id)
	} else {
		q.topics[topic] = CreateTopic(topic, q)
		q.topics[topic].DeleteMessage(id)
	}
	q.lock.Unlock()
}
