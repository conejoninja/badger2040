package main

import "encoding/base64"

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
	{"Wednesday 6th November",
		[]Talk{
			{"09:00", "12:00", "Guided Hack Session", "Workshop - morning session", "Ron Evans & TinyGo"},
			{"09:00", "12:00", "Golang Concurrency", "Workshop - morning session", "Roman Gurevitch"},
			{"09:00", "12:00", "Intro & Fast Track to Go", "Workshop - morning session", "Leah Garrett"},
			{"12:00", "13:30", "", "LUNCH BREAK", ""},
			{"13:30", "16:30", "Guided Hack Session", "Workshop - afternoon session", "Ron Evans & TinyGo"},
			{"13:30", "16:30", "Golang Concurrency", "Workshop - afternoon session", "Roman Gurevitch"},
			{"13:30", "16:30", "Intro & Fast Track to Go", "Workshop - afternoon session", "Leah Garrett"},
		},
	},
	{"Thursday 7th November",
		[]Talk{
			{"08:30", "09:30", "Registration", "Check-in and collect your lanyards", ""},
			{"09:30", "9:45", "", "Welcome & Acknowledgements", "Opening remarks"},
			{"09:45", "10:25", "", "Things I Find Myself Repeating About Go", "Dave Cheney"},
			{"10:25", "10:55", "", "Wrong Comparison is the Thief of Joy", "Elisa Xu"},
			{"10:55", "11:15", "BREAK", "Morning tea and Networking", "BREAK"},
			{"11:15", "11:45", "", "The Why of Iterator Design", "Axel Wagner"},
			{"11:45", "12:15", "The Gopher Orchestra: Distributed ", "Machine Learning with Go", "Rashmi Nagpal"},
			{"12:15", "13:45", "LUNCH", "At Everest Theatre, Seymour Centre", "LUNCH "},
			{"13:45", "14:15", "", "80% faster, 70% less memory", "Charles Korn and Joshua Hesketh"},
			{"14:15", "14:45", "", "Kubernetes Operators with Go", "Miguel Elias dos Santos"},
			{"14:45", "15:15", "", "Mastering Distributed Tracing with Go", "Renaldi Gondosubroto"},
			{"15:15", "15:35", "BREAK", "Coffee and Networking", "BREAK"},
			{"15:35", "16:05", "", "Overview of Gomobile and Its Use Cases", "Miho Sasaki"},
			{"16:05", "16:40", "Is Go A Good Language to Build", "Compilers?", "Sungmin Han"},
			{"16:40", "17:00", "", "Wrap-up and closing remarks", "CLOSE"},
		},
	},
	{"Friday 8th November",
		[]Talk{
			{"08:30", "09:00", "Registration", "Check-in and collect your lanyards", ""},
			{"09:00", "09:10", "", "Welcome & Acknowledgements", "Opening remarks"},
			{"09:10", "09:40", "", "Flight, FORTRAN, Control and Go", "Patricio Whittingslow"},
			{"09:40", "10:10", "Where the bloody hell are you?", "GIS in Go", "Darrell Chua"},
			{"10:10", "10:30", "BREAK", "Morning tea and Networking", "BREAK"},
			{"10:30", "11:00", "Turn Your Home Into a Discotheque", "with Go", "Esme Lamb"},
			{"11:00", "11:30", "", "Doing Something for Fun", "Dan Kortschak"},
			{"11:30", "12:00", "Lessons Learned: A Developer's", "Perspective on Teaching and Learning", "Leah Garrett"},
			{"12:00", "13:30", "LUNCH", "At Everest Theatre, Seymour Centre", "LUNCH"},
			{"13:30", "14:00", "When the 'Go'ing Gets Tough: Relieving", "Pressure on a Budget", "Ellen Gao"},
			{"14:00", "14:30", "WASM on Edge Serverless; Make Your Life", "Easier with Go and TinyGo", "Kay Sawada"},
			{"14:30", "15:00", "ReverseProxy is the Amazing Gem of", "the Go Std Lib", "Alexander Else"},
			{"15:00", "15:20", "BREAK", "Coffee and Networking", "BREAK"},
			{"15:20", "15:50", "Quick, insightful talks by", "various speakers", "Lightning Talks"},
			{"15:50", "16:35", "", "Go Back To The Future", "Ron Evans"},
			{"16:35", "17:00", "", "Wrap-up and closing remarks", "CLOSE"},
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
		"As usually, you are reading Golang Weekly, among very useful information you notice the CFP for GopherCon AU is open! It will be in Sydney, 6-8th November. What do you want to do?",
		"Oh no! you don't have any idea for a talk.",
		"Let's talk about AI, it's the new pink!",
		"You already have the slides for a TinyGo talk.",
		1, 2, 3,
	},
	{
		"It's ok, you still take advantage of the date and got an early-bird ticket at a discounted price:",
		"Self-payin gopher (A$385)",
		"Corporate gopher (A$500)",
		"Premium gopher + workshops (A$850)",
		4, 4, 4,
	},
	{
		"You worked hard on your slides and send your abstract in time. A few days passed and received the bad news that your talk wasn't accepted, " +
			"nobody is interested anymore in AI. But they offer you a discount price for the tickets:",
		"Self-payin gopher (A$335)",
		"Corporate gopher (A$450)",
		"Premium gopher + workshops (A$770)",
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
		"BEEP BEEP BEEP. It's your alarm clock. Today is the day, your need to prepare and go to the airport, Sydney and a bunch of gophers are waiting for you.",
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
		"The plane finally landed. You check in the hotel and have a few hours left to visit the city. Sydney is famous for:",
		"Those delicious Sydney rock oysters",
		"You know the lyrics of all AC/DC' songs",
		"The Sydney opera house",
		13, 14, 15,
	},
	{
		"You ask the concierge at the hotel about the best Sydney rock oysters in the city and go there. You make a mess of yourself, but it was delicious. When you start leaving the place, a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5AUD & proceed to listen to him",
		16, 16, 17,
	},
	{
		"Unfortunately there is no AC/CD concert today so you decide to go for a walk around the city to make time. When walking through a dark alley a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5AUD & proceed to listen to him",

		16, 16, 17,
	},
	{ // 15
		"You visit the opera house, it's amazing. While walking back to the hotel, you go through a dark alley and a weird looking person, too clean to be a bum, but with crazy person vibes and a tinfoil hat approaches. He tries to stop you.",
		"You run as fast as you can",
		"Try to ignore him",
		"Give him 5AUD & proceed to listen to him",
		16, 16, 17,
	},
	{
		"-\"Hey you... yeah " + YourName + "! " + YourName + " who works at " + YourCompany + " listen to me, I'm Ron Evans!, but the Ron Evans from the future!!, the year 2053, the last gopher. You need to listen carefully I have very little time, remember the code : <<IDKFA>> You are the only hope to save the planet and the humanity! (truust me)\" ... zooosh... and he disappeared",
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
		"Takasago can carry you in his motocross bike",
		"Share an e-scooter with Manolo",
		19, 19, 19,
	},
	{
		"You all arrive at the hotel and decided to take a drink at the bar. There, you encounter Angelica and Johnny recording a GoTime podcast, they ask you about your unpopular opinion:",
		"It's pronounced data, not data",
		"Rabbits are not rodent, they are lagomorphs",
		"Cheese cake is the best dessert",
		20, 21, 22,
	},
	{ // 20
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that data should be pronounced data instead of data. What do you think Angelica? I'm a bit socked myself. And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that rabbits are ... lagomorphs? Isn't that the Alien movie? That doesn't sound like an opinion but more like a fact. What do you think Johnny? I'm a bit socked myself. And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"-\"Hey wonderful listeners, we welcome " + YourName + " to our program, the gopher who states that the best dessert is the cheesecake. I've to say I totally agree with that. What do you think Angelica? And with this we end our episode for today\"",
		"Time to go to bed",
		"Time to go to bed",
		"Time to go to bed",
		23, 23, 23,
	},
	{
		"After a long day, you arrive to your room, your bed for the next days is waiting for you. You close your eyes. Looks like it's just a second, but wake up fully rested and ready for the first day of GopherCon. You skip breakfast because I'm too tired of adding options.",
		"-----",
		"Go to Seymour Centre",
		"(event location)",
		24, 24, 24,
	},
	{
		"You arrive at the door. A familiar face greets you, it's the wonderful Dave Cheney. You pick up your new TinyGo powered e-ink badge and look at it.",
		"-----",
		"Look at your badge",
		"-----",
		25, 25, 25,
	},
	{ // 25
		"",
		"-----",
		"-----",
		"-----",
		26, 26, 26,
	},
	{
		"You notice there's a SCHEDULE function in your badge, you could navigate today's schedule and pick your favourite talks you want to attend.",
		"Flight, FORTRAN, Control and Go - Patricio Whittingslow",
		"Overview of Gomobile and Its Use Cases - Marvin Hosea",
		"WASM on Edge Serverless - Kay Sawada",
		27, 27, 27,
	},
	{ // left empty, redirect to proper scene according the clothes.
		"",
		"",
		"",
		"",
		28, 29, 30,
	},
	{
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but some people stops you. They hand you a book and ask you to sign it, they have confused you with someone else. Since you are sporting your unicorn pijama, they think you are Aurelie, the famous Go book author.",
		"Sign it!",
		"Try to explain you are not her",
		"Je ne parle pas français",
		31, 31, 31,
	},
	{
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but some people stops you. They hand you a martini and ask you for crypto investment advice, they have confused you with someone else. Since you are sporting the faboulous suit and bow tie, they think you are Tanguy Herman, the best dressed gopher in the world.",
		"Shaken, not stirred",
		"Buy high, sell low... or the other way around",
		"Try to explain you are not her",
		31, 31, 31,
	},
	{ // 30
		"The talk was super fun and you learned a lot. You are preparing to go to the next talk but you find some old friends and catch up with your lives. You talked so much you missed the next talk and it's lunch time already.",
		"Yay! Veggie sandwiches.",
		"This quinoa salad is top.",
		"Not sure what I'm eating but it's superb.",
		32, 32, 32,
	},
	{
		"You try to explain they are mistaken to no avail. Anyway you talked so much you missed the next talk and it's lunch time already.",
		"Yay! Veggie sandwiches.",
		"This quinoa salad is top.",
		"Not sure what I'm eating but it's superb.",
		32, 32, 32,
	},
	{ // REDIRECT IN CASE  YOU GAVE A TALK OR NOT
		"",
		"",
		"",
		"",
		33, 34, 34,
	},
	{
		"After lunch it's time for your talk: flying drones, hens, LEDs, lasers, music,... you left the auditorium in awe, asking for more. Time for the grand finale, you only have one shot, one opportunity:",
		"FIREWORKS never fail.",
		"Launch a weather ballon.",
		"One word: Robot-laser-tag.",
		35, 36, 37,
	},
	{
		"After lunch it's time for another talk. The Ron Evans from your timeline is flying some drones inside the auditorium, one is going crazy and is going to attack some gophers. You could help him by rebooting the drone remotely, but need to introduce the reboot CODE:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{ // 35
		"Fireforks inside an auditorium? Yeah why not? Not sure who approved this but you need to introduce the launch CODE:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"You all go outside, a big weather balloon is waiting on the ground. Introduce the launch CODE to initiate the countdown:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"Autonomous robots armed with lasers? Yeah why not? Not sure who approved this but you need to introduce the login CODE for them to boot up:",
		"IDGAF",
		"IDKFA",
		"IDDQD",
		38, 39, 38,
	},
	{
		"Wrong code, rememmber what that crazy person yesterday was telling you. You have another opportunity:",
		"IDDQD",
		"IDKFA",
		"IDGAF",
		38, 39, 38,
	},
	{
		"Correct code, sequence initiated. YAY! The talk is a big success, #TinyGo is trending topic. GopherCon is coming to an end and only the social mixer is left to attend. You talked with many gopher there, laugh at bad tech jokes and make some friends.",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		40, 40, 40,
	},
	{ // 40
		"This is the end of this adventure. Continue to the next screen to see some stats of your adventure.",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		41, 41, 41,
	},
	{ // STATS PAGE
		"",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		42, 42, 42,
	},
	{ // QR PAGE
		"", "", "", "",
		43, 43, 43,
	},
	{ // THE END
		"Thank you for playing this adventure. I hope you enjoyed playing it as much as I (CONEJO) did making it. I hid some easter eggs (not very well hidden) and references. Feel free to share your opinion about it in person or online at @conejo@social.tinygo.org",
		"Click A to continue.",
		"Click B to continue.",
		"Click C to continue.",
		44, 44, 44,
	},
	{ // THE REAL END
		"", "", "", "",
		45, 45, 45,
	},
}
