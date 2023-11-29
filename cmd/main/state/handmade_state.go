package state

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type State2 int

const (
	OffHook State2 = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State2) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

type TriggerResult struct {
	Trigger Trigger
	State   State2
}

var Rules = map[State2][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{CallConnected, Connected},
		{PlacedOnHold, OnHold},
	},
	Connected: {
		{HungUp, OnHook},
		{LeftMessage, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

func RunHandmadeState() {
	state, exitState := OffHook, OnHook
	for state != exitState {
		fmt.Printf("[%s] Select a trigger\n", state)
		for i := 0; i < len(Rules[state]); i++ {
			triggerResult := Rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", triggerResult.Trigger)
		}
		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))
		state = Rules[state][i].State
	}
	fmt.Println("Hung up")
}
