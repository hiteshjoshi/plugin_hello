// Package hello responds to "Say something" with "Hello World".
package hello

import (
	"log"

	"github.com/itsabot/abot/shared/datatypes"
	"github.com/itsabot/abot/shared/plugin"
)

var p *dt.Plugin

func init() {
	// Create the plugin, setting it up to communicate with Abot through
	// the functions we specified.
	var err error
	p, err = plugin.New("github.com/hiteshjoshi/plugin_hello")
	if err != nil {
		log.Fatalln("failed to build plugin.", err)
	}

	// When Abot hears these Trigger words together, e.g. "Say something"
	// or "Tell me something," Abot will route to this plugin and run the
	// associated function (Fn), sending the returned string "Hello world!"
	// to the user.
	plugin.SetKeywords(p, dt.KeywordHandler{
		Fn: func(in *dt.Msg) string {
			return "Hello world!"
		},
		Trigger: &dt.StructuredInput{
			Commands: []string{"say", "tell"},
			Objects:  []string{"something"},
		},
	})

	if err = plugin.Register(p); err != nil {
		p.Log.Fatal("failed to register plugin.", err)
	}
}
