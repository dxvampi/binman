package main

import (
	"fmt"

	"github.com/dxvampi/binman/internal/store"
)

func main() {
	binaries, err := store.Load()
	if err != nil {
		fmt.Println("load error:", err)
		return
	}

	binaries = append(binaries, store.Binary{Alias: "java8", Path: "/usr/lib/jvm/java-8/bin/java"})

	err = store.Save(binaries)
	if err != nil {
		fmt.Println("save error:", err)
		return
	}

	fmt.Println(binaries)
}
