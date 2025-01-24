package main

import (
	"bytes"
	"debug/elf"
	"io"
	"os"

	"github.com/ethereum-optimism/optimism/cannon/mipsevm/multithreaded"
	"github.com/ethereum-optimism/optimism/cannon/mipsevm/program"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	elfProgram, err := elf.Open("program/bin/mt-hello.elf")
	if err != nil {
		panic(err)
	}

	meta, err := program.MakeMetadata(elfProgram)
	if err != nil {
		panic(err)
	}

	state, err := program.LoadELF(elfProgram, multithreaded.CreateInitialState)
	if err != nil {
		panic(err)
	}

	// err = program.PatchGoGC(elfProgram, state)
	// if err != nil {
	// 	panic(err)
	// }

	err = program.PatchStack(state)
	if err != nil {
		panic(err)
	}

	var stdOutBuf, stdErrBuf bytes.Buffer
	logger := log.NewLogger(log.LogfmtHandlerWithLevel(os.Stdout, log.LevelInfo))
	us := multithreaded.NewInstrumentedState(state, nil, io.MultiWriter(&stdOutBuf, os.Stdout), io.MultiWriter(&stdErrBuf, os.Stderr), logger, meta)

	maxSteps := 2_000_000
	for i := 0; i < maxSteps; i++ {
		if us.GetState().GetExited() {
			break
		}
		_, err := us.Step(true)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("%d: %+v\n", i, witness.StateHash.Hex())
	}
}
