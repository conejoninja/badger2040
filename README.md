# badger2040

## Getting Started

Go to the [tutorial](./tutorial/basics) and follow instructions.

## Troubleshooting

### Unable to locate any volume

If you are seeing this error, it means that the device needs to be reset:
```
unable to locate any volume: [RPI-RP2]
```
When the device is connected to your computer, press both the `BOOTSEL` button
on the board, and the `reset` button (both are located in the back of the badge)
for 5 seconds, then release the `reset` button first while still holding the
boot for another 5 seconds to reset the device and allow tinygo to flash the
device.

The device might disconnect after being flashed, you need to reset it each time
before flashing it (e.g., for each step of the activity).

If the problem persists, try holding `BOOTSEL` while connecting the USB cable
to your computer. This puts the device in [USB mass storage device mode](https://projects.raspberrypi.org/en/projects/getting-started-with-the-pico/3).

```console
$ fdisk -l
...
Device     Boot Start    End Sectors  Size Id Type
/dev/sdb1           1 262143  262143  128M  e W95 FAT16 (LBA)
```

Make note of the device path, which may be different on your system (e.g. `/dev/sda1`).

On Linux, ensure that the device is mounted with a volume name of `RPI-RP2`:

```console
$ mkdir /media/RPI-RP2
$ mount /dev/sdb1 /media/RPI-RP2
```
