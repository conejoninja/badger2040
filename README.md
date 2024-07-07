# badger2040

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
