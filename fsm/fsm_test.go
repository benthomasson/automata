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

    s := &fsm.State{Enter: noop1, Exit: noop1}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    c.ChangeState(s)
    describe(s)
    describe(c)
}

func TestMessage(t *testing.T) {

    s := &fsm.State{Enter: noop1, Exit: noop1}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    c.HandleMessage(fsm.SimpleMessage{})
    describe(s)
    describe(c)
}


func noop1(*fsm.State, *fsm.Controller) {
}

func noop2(*fsm.State, *fsm.Controller, fsm.Message) {
}


func TestReceiveMessages(t *testing.T) {

    s := &fsm.State{Enter: noop1, Exit: noop1}
    c := &fsm.Controller{State: s}
    c.ChangeState(s)
    ch := make(chan fsm.Message, 1)
    go c.ReceiveMessages(ch)
    ch <- fsm.SimpleMessage{}
    describe(s)
    describe(c)
    time.Sleep(1 * time.Millisecond)
}


