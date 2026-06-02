//go:build js && wasm
// +build js,wasm

package main

import (
	"flag"

	"github.com/vugu/vugu"
	"github.com/vugu/vugu/domrender"
)

func main() {
	mountPoint := flag.String("mount-point", "#vugu_mount_point", "query selector for the root mount point")
	flag.Parse()

	renderer, err := domrender.New(*mountPoint)
	if err != nil {
		panic(err)
	}
	defer renderer.Release()

	buildEnv, err := vugu.NewBuildEnv(renderer.EventEnv())
	if err != nil {
		panic(err)
	}

	rootBuilder := &Root{}

	for ok := true; ok; ok = renderer.EventWait() {
		buildResults := buildEnv.RunBuild(rootBuilder)

		if err := renderer.Render(buildResults); err != nil {
			panic(err)
		}
	}
}
