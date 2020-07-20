package fsm

import (
    "log"
)


type Message interface {
    MessageType() string
}

type SimpleMessage struct {
    MessageTypeName string
}

func (msg SimpleMessage) MessageType() string {
    return msg.MessageTypeName
}

type DataMessage struct {
    MessageTypeName string
    Data int
}


func (msg DataMessage) MessageType() string {
    return msg.MessageTypeName
}

type State struct {
    Name string
    Handlers map[string]func(*State, *Controller, Message)
    DefaultHandler func(*State, *Controller, Message)
    Enter func(*State, *Controller)
    Exit func(*State, *Controller)
}

func LogEnter(s *State, a *Controller, msg Message) {
    log.Printf("Enter")
}

func LogExit(s *State, a *Controller, msg Message) {
    log.Printf("Exit")
}

type Controller struct {
    State *State
    Handlers map[string]func(*Controller, Message)
    DefaultHandler func(*Controller, Message)
}

func (c *Controller) ChangeState(s *State) {
    if c.State != nil {
        c.State.Exit(c.State, c)
    }
    log.Printf("ChangeState to %s", s.Name)
    c.State = s
    c.State.Enter(c.State, c)
}

func (c *Controller) HandleMessage(msg Message) {

    if c.State == nil {
        // Do nothing. The FSM isn't turned on yet.
    } else if handler, ok := c.State.Handlers[msg.MessageType()]; ok {
        handler(c.State, c, msg)
    } else if c.State.DefaultHandler != nil {
        c.State.DefaultHandler(c.State, c, msg)
    } else if c.DefaultHandler != nil {
        c.DefaultHandler(c, msg)
    } else {
        log.Printf("No handler for %#v", msg)
    }
}

func (c *Controller) ReceiveMessages(inbox chan Message) {

    for message := range inbox {
        c.HandleMessage(message)
    }
}
