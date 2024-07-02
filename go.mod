module github.com/conejoninja/badger2040

go 1.22

require (
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/soypat/cyw43439 v0.0.0-20240627234239-a62ee4027d66
	github.com/soypat/seqs v0.0.0-20240527012110-1201bab640ef
	golang.org/toolchain v0.0.1-go1.9rc2.windows-amd64
	tinygo.org/x/drivers v0.28.0
	tinygo.org/x/tinydraw v0.4.0
	tinygo.org/x/tinyfont v0.4.0
)

require (
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/tinygo-org/pio v0.0.0-20231216154340-cd888eb58899 // indirect
	golang.org/x/exp v0.0.0-20230728194245-b0cb94b80691 // indirect
)

replace tinygo.org/x/drivers => github.com/conejoninja/drivers v0.0.0-20240515082542-5f2645f5444d
