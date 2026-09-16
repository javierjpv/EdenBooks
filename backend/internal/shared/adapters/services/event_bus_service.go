package services

import (
	"fmt"
	"sync"

	"github.com/asaskevich/EventBus"
)

type event struct {
    topic string
    data  interface{}
}

type GoEventBus struct {
    bus        EventBus.Bus
    mu         sync.RWMutex
    eventQueue chan event
    closed     bool
}

func NewGoEventBus() *GoEventBus {
    eb := &GoEventBus{
        bus:        EventBus.New(),
        eventQueue: make(chan event, 100),
        closed:     false,
    }
    go eb.processEvents()
    return eb
}

func (g *GoEventBus) processEvents() {
    for evt := range g.eventQueue {
        g.mu.RLock()
        g.bus.Publish(evt.topic, evt.data)
        g.mu.RUnlock()
    }
}

func (g *GoEventBus) Publish(topic string, data interface{}) error {
    if g.closed {
        return fmt.Errorf("event bus is closed")
    }
    
    select {
    case g.eventQueue <- event{topic: topic, data: data}:
        return nil
    default:
        // Si el buffer está lleno, publicar directamente
        g.mu.RLock()
        defer g.mu.RUnlock()
        go g.bus.Publish(topic, data)
        return nil
    }
}

func (g *GoEventBus) Subscribe(topic string, handler func(data interface{})) error {
    g.mu.Lock()
    defer g.mu.Unlock()
    return g.bus.Subscribe(topic, handler)
}

func (g *GoEventBus) Close() error {
    g.mu.Lock()
    defer g.mu.Unlock()
    if !g.closed {
        g.closed = true
        close(g.eventQueue)
    }
    return nil
}