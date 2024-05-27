package main

import (
	"time"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

func schedule(day int, hour int) {
	quit := false
	for {
		showDay(day, hour)
		for {
			time.Sleep(200 * time.Millisecond)
			if btnDown.Get() {
				hour++
				if hour > len(scheduleData[day].talks)-2 {
					hour = len(scheduleData[day].talks) - 2
				} else {
					break
				}
			}
			if btnUp.Get() {
				hour--
				if hour < 0 {
					hour = 0
				} else {
					break
				}
			}
			if btnA.Get() {
				day--
				if day < 0 {
					day = 0
				} else {
					hour = 0
					break
				}
			}
			if btnC.Get() {
				day++
				if day > len(scheduleData)-1 {
					day = len(scheduleData) - 1
				} else {
					hour = 0
					break
				}
			}
			if btnB.Get() {
				quit = true
				break
			}
		}
		if quit {
			break
		}
	}
}

func showDay(day int, hour int) {
	display.ClearBuffer()
	showRect(0, 0, WIDTH, 26, black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 8, 18, scheduleData[day].title, white)

	showRect(0, 116, WIDTH, 12, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 0, 125, " [a] PREV DAY        [b] EXIT       [c] NEXT DAY", white)

	tinydraw.Line(&display, 0, 71, WIDTH, 71, black)
	tinydraw.Line(&display, 50, 26, 50, 116, black)

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 44, scheduleData[day].talks[hour].startHour, black)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 66, scheduleData[day].talks[hour].endHour, black)

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 38, scheduleData[day].talks[hour].line1, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 50, scheduleData[day].talks[hour].line2, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 62, scheduleData[day].talks[hour].line3, black)

	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 89, scheduleData[day].talks[hour+1].startHour, black)
	tinyfont.WriteLine(&display, &freesans.Regular9pt7b, 2, 111, scheduleData[day].talks[hour+1].endHour, black)

	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 83, scheduleData[day].talks[hour+1].line1, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 95, scheduleData[day].talks[hour+1].line2, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 56, 107, scheduleData[day].talks[hour+1].line3, black)

	display.Display()
	display.WaitUntilIdle()

}
