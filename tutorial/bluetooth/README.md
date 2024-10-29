# Bluetooth Tutorial

This series of tutorials is intended to help you learn about programming for Bluetooth using TinyGo and the TinyGo [Bluetooth package](https://github.com/tinygo-org/bluetooth)


## What you need

    - Pimoroni Badger2040W
    - Personal computer with Go 1.23 and TinyGo 0.34 installed, and a serial port.


## Code

### step1.go - Bluetooth scan

First run some code to scan your local area for Bluetooth devices. This program will use the Bluetooth interface in your Badger2040-W and show the output from it on your computer terminal using the USB interface.

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/bluetooth/step1
```

You will see output on your terminal. Each of the devices list is a Bluetooth device nearby that is advertising itself.


### step2.go - Bluetooth scan on Badger2040-W display

Next step is to scan for Bluetooth device as in the previous step, but displaying the output on the Badger2040-W e-ink display.

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/bluetooth/step2
```

You will see the bluetooth scan output on both your monitor and on the Badger2040-W display.


### step3.go - Bluetooth discover

Now that you know how to find Bluetooth devices that are nearby you, you can proceed to try to connect to one of them and find out what services it can offer.

You will need to use the MAC address (Linux or Windows) or the Bluetooth ID (macOS) to connect to a specific device.

Try one of of the devices you found when you were scanning in step1/step2.

Run the code.

Note that not all devices will allow you to connect to them, and that some that allow you to connect will not allow you to view the details of every service/characteristic.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./tutorial/bluetooth/step3
```

### step4.go - Bluetooth discover on Badger2040-W display

This is the same Bluetooth service discovery as the previous example, but it also shows the data on the Badger2040-W display.

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./tutorial/bluetooth/step4
```

You should see the output on both your terminal, and also on the Badger2040-W display.


### step5.go - Bluetooth heart rate

Now that you know how to find Bluetooth devices that are nearby you and how to connect to them, you can proceed to try to do something useful.

Let's connect the Badger2040-W to a Bluetooth heart rate sensor.

If you already have smart watch or app on your phone, you can connect to it. Otherwise you can obtain one for your mobile device such as those listed here:

https://www.cnet.com/health/how-to-track-your-heart-rate-with-a-smartphone/

You can also run a simulator on your laptop computer:

```shell

go run ./tutorial/bluetooth/heartsim
```

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./tutorial/bluetooth/step5
```

You can connect from the Badger2040-W to your mobile phone or any other device/software that can produce the data from a standard Bluetooth heart rate device.


### step6.go - Bluetooth heart rate on Badger2040-W display

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=[MAC address or Bluetooth ID goes here]" ./tutorial/bluetooth/step6
```

This is the same heart rate device as the previous example, but it also shows the data on the Badger2040-W display. You will still need to connect to your mobile phone or any other device/software that can produce the data for a standard Bluetooth heart rate device.
