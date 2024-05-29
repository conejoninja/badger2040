package main

import "golang.org/toolchain/src/encoding/base64"

// Replace with your data by using -ldflags like this:
//
// tinygo flash -target badger2040 -ldflags="-X main.YourName=CONEJO -X main.YourTitle='Amazing human' -X main.YourSocial='@conejo@social.tinygo.org'"
//
// See Makefile for more info.
var (
	YourName, YourCompany, YourTitle, YourSocial, ProfilePic string
)

var profileImg []uint8
var profileErr error

const (
	DefaultCompany = "TinyGo"
	DefaultName    = "CONEJO"
	DefaultTitle   = "Technologist for hire"
	DefaultSocial  = "@conejo@social.tinygo.org"
)

type Talk struct {
	startHour, endHour, line1, line2, line3 string
}

type Day struct {
	title string
	talks []Talk
}

var scheduleData = []Day{
	{"Monday June 17 (TZ GMT+2)",
		[]Talk{
			{"10:00", "13:00", "", "Self-booked social activities", ""},
			{"14:00", "18:00", "Contributors summit", "(Roundtable discussiones open to all)", ""},
		},
	},
	{"Tuesday June 18 (TZ GMT+2)",
		[]Talk{
			{"09:00", "18:00", "", "Workshops at the venue", "(dedicated tocket required)"},
			{"14:00", "17:00", "", "HARDWARE HACKING WITH GO", ""},
		},
	},
	{"Wednesday June 19 (TZ GMT+2)",
		[]Talk{
			{"09:00", "10:00", "", "Breakfast", ""},
			{"10:00", "10:30", "", "Opening words", ""},
			{"10:30", "11:00", "", "The business case for Go", "Cameron Balahan"},
			{"11:00", "11:15", "", "15 min break", ""},
			{"11:15", "11:45", "Pack Your Bytes, We're Building: Memory", "Optimization Through Structure Packing", "Diana Shevchenko - TRACK 1"},
			{"11:15", "11:45", "Rethinking Domain-Driven Design in Go:", "From Myths to Reduced Project Complexity", "Robert Laszczak - TRACK 2"},
			{"11:45", "12:00", "", "15 min break", ""},
			{"12:00", "12:30", "", "A deep dive into de DB connection pool", "Agniva De Sarker - TRACK 1"},
			{"12:00", "12:30", "From Java to Go: I have a hammer", "and see nails everywhere", "Martin Gallauner - TRACK 2"},
			{"12:30", "14:00", "", "Lunch break", ""},
			{"14:00", "14:30", "", "Building AI applications in Go", "Travis Cline"},
			{"14:30", "14:45", "", "15 min break", ""},
			{"14:45", "15:15", "Leveraging Go for efficient", "infrastructure and data handling", "Chioma Onyekpere - TRACK 1"},
			{"14:45", "15:15", "Using formal reasoning to build", "concurrent Go systems", "Raghav Roy - TRACK 2"},
			{"15:30", "16:15", "Panel with the Go team", "Cameron Balahan, Alice Merrick", "Jonathan Amsterdam, Zvonimir Pavlinovic"},
			{"16:15", "17:00", "", "Coffee break", ""},
			{"17:00", "17:45", "", "GoTime podcast - Live recording", ""},
			{"17:45", "18:00", "", "Closing words", ""},
			{"18:00", "22:00", "Music and drinks", "(open bar, drinks included in the", "conference pass)"},
		},
	},
	{"Thursday June 20 (TZ GMT+2)",
		[]Talk{
			{"09:00", "10:00", "", "Breakfast", ""},
			{"10:00", "10:30", "", "Opening words", ""},
			{"10:30", "11:00", "", "HTTP routing enhancements", "Jonathan Amsterdam"},
			{"11:00", "11:15", "", "15 min break", ""},
			{"11:15", "11:45", "Accelerate your applications with KEDA:", "a beginner's guide to auto-scaling and", "event-driven arch. Ayesha Kaleem TRACK1"},
			{"11:15", "11:45", "Unraveling Go anti-patterns for", "clean and efficient code", "Rabieh Fashwall - TRACK 2"},
			{"11:45", "12:00", "", "15 min break", ""},
			{"12:00", "12:30", "", "How to win frames and influence pointers", "Felix Geisendorfer - TRACK 1"},
			{"12:00", "12:30", "Technical documentation: How can I write", "them better and why should I care?", "Hila Fish - TRACK 2"},
			{"12:30", "14:00", "", "Lunch break", ""},
			{"14:00", "14:45", "Securing containers against know", "Go vulnerabilities", "Zvonimir Pavlinovic"},
			{"14:45", "15:00", "", "15 min break", ""},
			{"15:30", "16:15", "", "Lightning talks", ""},
			{"16:15", "16:30", "", "Closing words", ""},
			{"16:30", "18:00", "", "Networking", ""},
		},
	},
}

func setCustomData() {
	if YourCompany == "" {
		YourCompany = DefaultCompany
	}

	if YourName == "" {
		YourName = DefaultName
	}

	if YourTitle == "" {
		YourTitle = DefaultTitle
	}

	if YourSocial == "" {
		YourSocial = DefaultSocial
	}

	profileImg, profileErr = base64.StdEncoding.DecodeString(ProfilePic)

}

type Scene struct {
	description, optionA, optionB, optionC string
	sceneA, sceneB, sceneC                 int
}

var sceneData = []Scene{
	{"123456789 123456789 123456789 123456789 123456789 ", "[A] Option A of dialogue", "B- Option B of dialogue", "(C) Option C of dialogue", 0, 0, 0},
}
