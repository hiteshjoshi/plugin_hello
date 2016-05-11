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
	// Create the plugin, setting it up to communicate with Abot through
	// the functions we specified.
	var err error
	p, err = plugin.New("github.com/itsabot/plugin_hello")
	if err != nil {
		log.Fatalln("failed to build plugin.", err)
	}

	// When Abot receives a message, it'll route the message to the correct
	// package. Doing that requires a trigger, which tells Abot to send the
	// response to this package when Commands include "say" and Objects
	// include "something", "hello", etc. Case should always be lowercase,
	// and the words will be stemmed automatically, so there's no need to
	// include variations like "cat" and "cats". plugin.AppendTrigger is
	// optional if you set KeywordHandlers (as described in the Building a
	// Plugin guide), but since we only have a state machine, we'll add
	// these words as triggers.
	plugin.AppendTrigger(p, &nlp.StructuredInput{
		Commands: []string{"say"},
		Objects:  []string{"something", "hello", "hi"},
	})

	// Abot includes a state machine designed to have conversations. This
	// is the simplest possible example, but we'll cover more advanced
	// cases with branching conversations, conditional next states, memory,
	// jumps and more in other guides.
	//
	// For more information on state machines in general, see:
	// https://en.wikipedia.org/wiki/Finite-state_machine
	plugin.SetStates(p, [][]dt.State{[]dt.State{
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
	}})
}
