package main

import (
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/proggy"
)

func adventure() {
	quit := false
	for {
		scene(0)
		for {
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
		tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 4, 20+i*10, ss[i], black)
	}

	showRect(0, 92, WIDTH, 36, black)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 6, 104, sceneData[s].optionA, white)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 6, 114, sceneData[s].optionB, white)
	tinyfont.WriteLine(&display, &proggy.TinySZ8pt7b, 6, 124, sceneData[s].optionC, white)

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
