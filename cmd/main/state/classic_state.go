package state

import "fmt"

type Switch struct {
	state State
}

func NewSwitch(on bool) *Switch{
	if on {
		return &Switch{state: NewOnState()}
	} else {
		return &Switch{state: NewOffState()}
	}
}

func (s *Switch) On() {
	s.state.On(s)
}

func (s *Switch) Off() {
	s.state.Off(s)
}

type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

type BaseState struct {

}

func (s *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (s *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}



type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	return &OnState{BaseState: BaseState{}}
}

type OffState struct {
	BaseState
}

func (s *OnState) Off(sw *Switch) {
	fmt.Println("Switching light off")
	sw.state = NewOffState()
}

func NewOffState() *OffState {
	return &OffState{BaseState: BaseState{}}
}

func (s *OffState) On(sw *Switch) {
	fmt.Println("Switchling light on")
	sw.state = NewOnState()
}

func RunClassicState() {
	sw := NewSwitch(false)
	sw.Off()
	sw.On()
	sw.On()
	sw.Off()
}



