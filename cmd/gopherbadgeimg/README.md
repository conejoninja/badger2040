# gopherbadgeimg

Simple, educational tool to transform images to the bitmap format supported by
the 2024 gophercon badger-w (and therefore many TinyGo devices!).

All of the most interesting bits in the code are heavily documented in-line.

Build the tool using go build:

`go build .`

Run it as follows to generate a profile image:

`./gopherbadgeimg -outmode base64 -ratio profile gopher-base.png`

Or, for the splash screen:

`./gopherbadgeimg -outmode rice -ratio splash -show --disable-dithering tainigo_128.png`

You can include the image in 3 different formats:

1. In the [Makefile](https://github.com/hybridgroup/badger2040/blob/main/Makefile)
from the badge repo, you're allowed to replace the profile image by dropping in
base64-encoded pixel data.
This can be printed to STDOUT using `--outmode base64`, and can be pasted directly into the Makefile.
1. The generator will also create a .bin file, which can be embedded into your
code using `go:embed`. An example of this can be found in `main_test.go`.
Use `--outmode bin` to create the bin files.
1. Finally, the generator can create a `*-generated.go` file similar to the
[tainigo.go](https://github.com/hybridgroup/badger2040/blob/main/tainigo.go) file,
to demonstrate an alternative to go embed. Use mode `--outmode rice` to create this file.
The option name is a reference to [an elegant package from a more civilized age.](https://github.com/GeertJohan/go.rice)
