// Package hello responds to "Say something" with "Hello World".
package hello

import (
	"log"

	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/nlp"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin
var sm *dt.StateMachine

func init() {
	// When Abot receives a message, it'll route the message to the correct
	// package. Doing that requires a trigger, which tells Abot to send the
	// response to this package when Commands include "say" and Objects
	// include "something", "hello", etc. Case should always be lowercase,
	// and the words will be stemmed automatically, so there's no need to
	// include variations like "cat" and "cats".
	trigger := &nlp.StructuredInput{
		Commands: []string{"say"},
		Objects:  []string{"something", "hello", "hi"},
	}

	// Tell Abot how we'll respond to first messages and follow up messages
	// from a user in a conversation.
	fns := &dt.PluginFns{Run: Run, FollowUp: FollowUp}

	// Create the package, setting it up to communicate with Abot through
	// the functions we specified.
	var err error
	p, err = plugin.New("github.com/itsabot/plugin_hello", trigger, fns)
	if err != nil {
		log.Fatalln("building", err)
	}

	// Abot includes a state machine designed to have conversations. This
	// is the simplest possible example, but we'll cover more advanced
	// cases with branching conversations, conditional next states, memory,
	// jumps and more in other guides.
	//
	// For more information on state machines in general, see:
	// https://en.wikipedia.org/wiki/Finite-state_machine
	sm = dt.NewStateMachine(p)
	sm.SetStates(
		[]dt.State{
			{
				OnEntry: func(in *dt.Msg) string {
					return "Hello world!"
				},
				OnInput: func(in *dt.Msg) {
				},
				Complete: func(in *dt.Msg) (bool, string) {
					return true, ""
				},
			},
		},
	)
}

// Run is called when the user first enters the package in a conversation. Here
// we simply reset the state of the state machine and pass it on to FollowUp().
func Run(in *dt.Msg) (string, error) {
	sm.Reset(in)
	return FollowUp(in)
}

// FollowUp is called as the user continues to message the package after the
// user's first message. Here we tell the state machine to move to the next
// available state and return the response from that move--since there's only
// one state, it will always be "Hello world!" from the OnEntry function.
func FollowUp(in *dt.Msg) (string, error) {
	return sm.Next(in), nil
}
