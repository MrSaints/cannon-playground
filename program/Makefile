.PHONY: elf32
elf32: $(patsubst %/go.mod,bin/%.elf,$(wildcard */go.mod))

.PHONY: elf
elf: elf32

bin:
	mkdir bin

bin/%.elf: bin
	cd $(@:bin/%.elf=%) && GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o ../$@ .
