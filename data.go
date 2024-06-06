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
	{
		"As usually, you are reading Golang Weekly, among very useful information you notice the CFP for GopherCon.EU is open! It will be in Berlin, 17-20th June. What do you want to do?",
		"Oh no! you don't have any idea for a talk.",
		"Let's talk about AI, it's the new pink!",
		"You already have the slides for a TinyGo talk.",
		1, 2, 3,
	},
	{
		"It's ok, you still take advantage of the date and got an early-bird ticket at a discounted price:",
		"Self-payin gopher (385€)",
		"Corporate gopher (600€)",
		"Premium gopher + workshops (1050€)",
		4, 4, 4,
	},
	{
		"You worked hard on your slides and send your abstract in time. A few days passed and received the bad news that your talk wasn't accepted, " +
			"nobody is interested anymore in AI. But they offer you a discount price for the tickets:",
		"Self-payin gopher (335€)",
		"Corporate gopher (550€)",
		"Premium gopher + workshops (970€)",
		4, 4, 4,
	},
	{
		"Flying drones, hens, LEDs, lasers, music,... who will not like such a talk about TinyGo? Your talk of course was accepted and your trip is fully paid.",
		"YAY!",
		"Super yay!",
		"OMG!!1! a dream come true",
		4, 4, 4,
	},
	{
		"BEEP BEEP BEEP. It's your alarm clock. Today is the day, your need to prepare and go to the airport, Berlin and a bunch of gophers are waiting for you.",
		"Take a taxi to the airport. No time to waste!",
		"Have Liam's signature pancakes.",
		"Sleep a bit more",
		5, 6, 4,
	},
	{ // 5
		"You took a taxi, arrived at the airport with enough time, there was no need to rush that much. When you go to the check-in baggage desktop you realize you forgot your suitcase with everything. You are only wearing your unicorn-pijama.",
		"No time to go back, you board the plane.",
		"You shop something at the airport's shop.",
		"You don't care. YOLO! Go to your gate.",
		7, 8, 7,
	},
	{
		"You have the best breakfast ever and feel full of energy, so much you pick up your suitcase and RUN to the airport. You pass the security check without issues.",
		"Better not waste time, board the plane.",
		"You shop something at the airport's shop.",
		"A quick visit to the toilet before boarding.",
		7, 8, 7,
	},
	{
		"You are finally at the plane, go up to your seat: 27B. Your seat neighborg looks like a gopher too, who is holding a shark plushie and a network switch.",
		"Ask about the shark.",
		"Ask about the network switch.",
		"Goeiemorgen.",
		9, 10, 11,
	},
	{
		"You've bought the most expensive suit and bowtie of your life, but holy guacomole, what a nice and superb suit. You look awesome.",
		"Go to the gate and board the plane",
		"Go to the gate and board the plane",
		"Go to the gate and board the plane",
		7, 7, 7,
	},
	{
		"-\"This shark? It's name is Blahaj and is the administrator of our mastodon instance. This is probably the weirdest train I've ever been in, I've never seen a train with wings before!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Me neither",
		12, 12, 12,
	},
	{ // 10
		"-\"This switch? I need it for the uplink of our mastodon instance. This is probably the weirdest train I've ever been in, I've never seen a train with wings before!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Me neither",
		12, 12, 12,
	},
	{
		"-\"Dit is waarschijnlijk de vreemdste trein waar ik ooit in heb gezeten, ik heb nog nooit een trein met vleugels gezien!\"",
		"Yeah... sure... ",
		"Oh cool",
		"Ik ook niet",
		12, 12, 12,
	},
	{
		"The plane finally landed. You check in the hotel and have a few hours left to visit the city. Berlin is famous for:",
		"Currywurst is a matter of pride for Berliners",
		"You know the lyrics of all Rammstein's songs",
		"The Brandenburg Gate is a must",
		13, 14, 15,
	},
	{
		"You ask the concierge at the hotel about the best Currywurst in the city and go there. You make a mess of yourself, but it was delicious. When you start leaving the place, a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5EUR & proceed to listen to him",
		16, 16, 17,
	},
	{
		"Unfortunately there is no Rammstein concert today so you decide to go for a walk around the city to make time. When walking through a dark alley a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5EUR & proceed to listen to him",

		16, 16, 17,
	},
	{ // 15
		"You visit the Brandenburg Gate, it's amazong. While walking back to the hotel, you go through a dark alley and a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5EUR & proceed to listen to him",
		16, 16, 17,
	},
	{
		"-\"Hey you... yeah " + YourName + "! Listen to me, I'm Ron Evans!, but the Ron Evans from the future!!, the year 2053, the last gopher. You need to listen carefully I have very little time, remember the code : <<IDKFA>> You are the only hope to save the planet and the humanity! (truust me)\" ... zooosh... and he disappeared",
		"Wait! What just happened?",
		"IDKFA... IDKFA... IDKFA... you try to remember",
		"....",
		18, 18, 18,
	},
	{
		"-\"Oh hello " + YourName + "! Thank you for listening to me, I'm Ron Evans, but the Ron Evans from the future!!, the year 2053, the last gopher. You need to listen carefully I have very little time, remember the code : <<IDKFA>> You are the only hope to save the planet and the humanity! (truust me)\" ... zooosh... and he disappeared",
		"Wait! What just happened?",
		"IDKFA... IDKFA... IDKFA... you try to remember",
		"....",
		18, 18, 18,
	},
	{
		"After this misterious encounter, you want to go back to the hotel as fast as you can. In front of you there is a group of friendly gophers. Thankfully they are going to the same hotel and can carry you. But each one has it's own method of transportation.",
		"You truust Francesc with his bicycle",
		"Sago35 can carry you in his motocross bike",
		"Share an e-scooter with Manolo",
		19, 19, 19,
	},
	{
		"You all arrive at the hotel and decided to take a drink at the bar. There you encounter Mat and Nathalie recording a GoTime podcast, they ask you about your unpopular opinion:",
		"sock-sock-shoe-shoe is the correct way",
		"rabbits should be the internet animal",
		"",
		20, 21, 22,
	},
}
