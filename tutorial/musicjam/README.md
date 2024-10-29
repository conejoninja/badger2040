# TinyGo Music Jam

Make music using your own TinyGo-based customized MIDI controller using audio software running on your notebook computer.

```
┌────────────────────────────┐      ┌────────────────────────────────────────────────┐
│                            │      │                                                │
│ ┌────────────────────────┐ │      │ ┌──────────────────────┐                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │     MIDI Controller    │ │      │ │       USB-MIDI       │                       │
│ │                        ├─┼──────┼─►                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ │                        │ │      │ │                      │                       │
│ └────────────────────────┘ │      │ └──────────┬───────────┘                       │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │            │                                   │
│                            │      │ ┌──────────▼───────────┐                       │
│                            │      │ │                      ├─────────────────────┐ │
│                            │      │ │                      │                     │ │
│                            │      │ │     Web MIDI API     │ Web Software Synth  │ │
│                            │      │ │                      │                     │ │
│                            │      │ │                      ├─────────────────────┘ │
│                            │      │ │                      │                       │
│                            │      │ │                      │                       │
│                            │      │ │                      │                       │
│                            │      │ └──────────────────────┘                       │
│                            │      │                                                │
└────────────────────────────┘      └────────────────────────────────────────────────┘

  Badger2040                        Computer

```

Thanks to USB-MIDI standard, the Badger2040 will appear as a standard MIDI controller. You can use it to connect to online instruments that use the Web MIDI API.


## Online Synths and Instruments

This is just a list of a few of the available online synths and other virtual instruments.

https://midi.city/

https://www.websynths.com/microtonal/

https://www.gsn-lib.org/apps/cardboardsynth/index.html

https://g200kg.github.io/webaudio-tinysynth/soundedit.html

https://signal.vercel.app/

https://juno-106.js.org/

https://virtualpiano.eu/

https://experiments.withgoogle.com/ai/sound-maker/view/

## Hardware Controllers

A hardware MIDI controller lets you control the virtual instruments running on your computer.

Each of the TinyGo programs for the MIDI Controller is intended to run directly on the external hardware to send MIDI commands via the USB interface to your computer.

All of the musical activities can be done with the Badger2040.


### onenote

This introductory MIDI controller sends only a single note.

To build/flash the `onenote` program:

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/musicjam/onenote
```

Press the A button.

This should send MIDI messages that can trigger sounds on your computer by using your TinyGo MIDI controller.

Open a web page with one of the online synths listed above. You should be able to use your new custom MIDI controller to make music.

Have fun!


### chorder

This MIDI controller sends entire chords with a single touch.

Each time you press the controller, it will play the next chord in the programmed chord progression.

To build/flash the `chorder` program:

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/musicjam/chorder
```

Launch one of the online synths. You should be able to use your new custom MIDI controller to make music.

Have fun!


### therestick

This MIDI controller lets you use a joystick as a digital theremin.

You must have the PCF8591 Analog to Digital Converter board and an analog joystick.

To build/flash the `therestick` program:

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/musicjam/therestick
```

Launch one of the online synths. You should be able to use your new custom MIDI controller to make music.

Have fun!


### arpeggio

This MIDI controller plays arpeggios when you press one of the buttons.

To build/flash the `arpeggio` program:

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/musicjam/arpeggio
```

Launch one of the online synths. You should be able to use your new custom MIDI controller to make music.

Have fun!
