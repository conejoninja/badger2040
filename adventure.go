package main

import (
	"strconv"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
	"tinygo.org/x/tinyfont/proggy"
)

func adventure() {
	quit := false
	s := 0
	score := 0
	money := 0
	costume := 0
	pancake := false
	talk := false
	opinion := 0
	for {
		println(s)
		scene(s)
		score += s
		if s == 5 {
			costume = 1
		} else if s == 3 {
			talk = true
		} else if s == 6 {
			pancake = true
		} else if s == 8 {
			money += 689
			costume = 2
		} else if s == 17 {
			money += 5
		} else if s == 20 {
			opinion = 57
		} else if s == 21 {
			opinion = 37
		} else if s == 22 {
			opinion = 21
		} else if s == 25 {
			badgeProfile()
		} else if s == 27 {
			if costume == 0 {
				s = 30
			} else if costume == 1 {
				s = 28
			} else {
				s = 29
			}
			continue
		} else if s == 32 {
			s = 34
			if talk {
				s = 33
			}
			continue
		} else if s == 41 {
			showGameStats(score, money, costume, opinion, pancake, talk)
		} else if s == 42 {
			showGameQR()
		} else if s == 44 {
			return
		}
		for {
			if btnA.Get() {
				s = sceneData[s].sceneA
				if s == 1 {
					money += 385
				} else if s == 2 {
					money += 335
				}
				break
			}
			if btnB.Get() {
				s = sceneData[s].sceneB
				if s == 1 {
					money += 600
				} else if s == 2 {
					money += 550
				}
				break
			}
			if btnC.Get() {
				s = sceneData[s].sceneC
				if s == 1 {
					money += 1050
				} else if s == 2 {
					money += 970
				}
				break
			}
			if btnUp.Get() || btnDown.Get() {
				return
			}
			time.Sleep(200 * time.Millisecond)

		}
		if quit {
			break
		}
	}
}

func scene(s int) {
	display.ClearBuffer()

	ss := splitBefore(sceneData[s].description)
	for i := int16(0); i < int16(len(ss)); i++ {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 4, 10+i*10, ss[i], black)
	}

	showRect(0, 95, WIDTH, 33, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 2, 104, "[A] "+sceneData[s].optionA, white)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 2, 114, "[B] "+sceneData[s].optionB, white)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 2, 124, "[C] "+sceneData[s].optionC, white)

	display.Display()
	display.WaitUntilIdle()
}

func splitBefore(str string) []string {
	l := len(str)
	a := 0
	s := make([]string, 1)
	for {
		if l <= 48 {
			s = append(s, str[a:a+l])
			break
		} else {
			for i := 48; i > 0; i-- {
				if string(str[a+i]) == " " {
					s = append(s, str[a:a+i])
					a = a + i + 1
					l = l - i - 1
					break
				}
			}
		}
	}
	return s
}

func showGameStats(score, money, costume, opinion int, pancake, talk bool) {

	c := "SCORE: " + strconv.Itoa(score)
	w32, _ := tinyfont.LineWidth(&freesans.Bold9pt7b, c)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, (WIDTH-int16(w32))/2, 26, c, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 50, "Money spent: "+strconv.Itoa(money)+" USD", black)

	c = "YOURSELF"
	if costume == 1 {
		c = "UNICORN PIJAMA"
	} else if costume == 2 {
		c = "DELUXE TUXEDO"
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 60, "You were dressed as: "+c, black)

	c = "NO pancakes"
	if pancake {
		c = "37 pancakes"
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 70, "Your breakfast: "+c, black)

	if talk {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 80, "You DID give a talk (and it was awesome)", black)
	} else {
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 80, "You did NOT give a talk (sad noises)", black)
	}
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 10, 90, "Your opinion is "+strconv.Itoa(opinion)+"% popular", black)

	display.Display()
	display.WaitUntilIdle()
}

func showGameQR() {
	// https://tinyurl.com/gceuADV
	QR("https://tinyurl.com/gceuADV")

	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 30, "SCAN &", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 50, "SHARE", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 70, "YOUR", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 90, "SCORE", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 10, 110, "ONLINE", black)

	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 210, 30, "SCAN &", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 210, 50, "SHARE", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 210, 70, "YOUR", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 210, 90, "SCORE", black)
	tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 210, 110, "ONLINE", black)

	display.Display()
	display.WaitUntilIdle()
}
