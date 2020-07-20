package fsm_test

import (
    "testing"
    "automata_rpg/fsm"
    "fmt"
    "reflect"
    "time"
)

func describe(i interface{}) {
    fmt.Printf("(%v, %#v, %T)\n", i, i, i)
}

type HelloFunc func(string)

func SayHello(to string) {
    fmt.Printf("Hello, %s \n", to)
}

func noop1(*fsm.State, *fsm.Controller) {
}

func noop2(*fsm.State, *fsm.Controller, fsm.Message) {
}

var value int = 1;

func setValue1(*fsm.State, *fsm.Controller) {
    value = 4
}

func setValue2(*fsm.State, *fsm.Controller, fsm.Message) {
    value = 5
}

func Test1(t *testing.T) {
    ch := make(chan fsm.Message, 2)
    ch <- fsm.SimpleMessage{MessageTypeName: "foo1"}
    ch <- fsm.DataMessage{MessageTypeName: "foo2"}
    fmt.Printf("%s\n", (<-ch).MessageType())
    fmt.Printf("%s\n", reflect.TypeOf((<-ch)).String())
}

func Test2(t *testing.T) {
    SayHello("me")
    hf := SayHello
    hf("me")

    var hf2 HelloFunc

    hf2 = SayHello

    hf2("me")

    describe(hf)

}


func TestState(t *testing.T) {

    value = 1

    s := &fsm.State{Enter: noop1, Exit: setValue1}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    if c.State != s {
        t.Errorf("Wrong state")
    }
    if value != 4 {
        t.Errorf("Exit handler not run")
    }
}

func TestMessage(t *testing.T) {

    value = 1
    handlers := map[string]func(*fsm.State, *fsm.Controller, fsm.Message){
        "SimpleMessage": setValue2,
    }

    s := &fsm.State{Enter: setValue1, Exit: noop1, Handlers: handlers}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    if c.State != s {
        t.Errorf("Wrong state")
    }

    if value != 4 {
        t.Errorf("Enter handler not run")
    }

    c.HandleMessage(fsm.SimpleMessage{MessageTypeName: "SimpleMessage"})

    if value != 5 {
        t.Errorf("Message handler not run")
    }
}


func TestReceiveMessages(t *testing.T) {

    value = 1
    handlers := map[string]func(*fsm.State, *fsm.Controller, fsm.Message){
        "SimpleMessage": setValue2,
    }

    s := &fsm.State{Enter: setValue1, Exit: noop1, Handlers: handlers}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    if c.State != s {
        t.Errorf("Wrong state")
    }
    if value != 4 {
        t.Errorf("Enter handler not run")
    }
    ch := make(chan fsm.Message, 1)
    go c.ReceiveMessages(ch)
    ch <- fsm.SimpleMessage{MessageTypeName: "SimpleMessage"}
    time.Sleep(1 * time.Millisecond)
    if value != 5 {
        t.Errorf("Message handler not run")
    }
}

func TestValue(t *testing.T) {

    value = 1 

    if value != 1 {
        t.Errorf("wrong value")
    }
    value = 2
    if value != 2 {
        t.Errorf("wrong value")
    }
}


