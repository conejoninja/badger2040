# Bluetooth Tutorial

## What you need

    - Pimoroni Badger2040W
    - Personal computer with Go 1.22 and TinyGo 0.32 installed, and a serial port.

## Code

### step1.go - Bluetooth scan

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor ./tutorial/bluetooth/step1
```

You will see output on your terminal.

### step2.go - Bluetooth scan on Badger2040-W display

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb ./tutorial/bluetooth/step2
```

You will see the bluetooth scan output on the Badger2040-W display.

### step3.go - Bluetooth heart rate

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=72:20:D2:1B:75:7E" ./tutorial/bluetooth/step3
```

You can connect from the Badger2040-W to your mobile phone or any other device/software that can produce the data from a standard Bluetooth heart rate device.

### step4.go - Bluetooth heart rate on Badger2040-W display

Run the code.

```shell
tinygo flash -target badger2040-w -stack-size=8kb -monitor -ldflags="-X main.DeviceAddress=72:20:D2:1B:75:7E" ./tutorial/bluetooth/step4
```

This is the same heart rate device as the previous example, but it also shows the data on the Badger2040-W display. You will still need to connect to your mobile phone or any other device/software that can produce the data for a standard Bluetooth heart rate device.
