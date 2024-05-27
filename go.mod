module github.com/conejoninja/badger2040

go 1.22.1

require (
	golang.org/toolchain v0.0.1-go1.9rc2.windows-amd64
	golang.org/x/image v0.16.0
	tinygo.org/x/drivers v0.27.1-0.20240509133757-7dbca2a54349
	tinygo.org/x/tinydraw v0.4.0
	tinygo.org/x/tinyfont v0.4.0
)

require (
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/makeworld-the-better-one/dither/v2 v2.4.0 // indirect
)

replace github.com/aykevl/tinygl => github.com/hybridgroup/tinygl v0.0.0-20240510175839-ce978081c4d0

replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240515082542-5f2645f5444d
