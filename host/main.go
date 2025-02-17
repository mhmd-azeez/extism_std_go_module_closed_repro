package main

import (
	"context"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
)

func main() {
	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{Path: "../plugin/main.wasm"},
		},

		// this works (when calling count_vowels)
		// Wasm: []extism.Wasm{
		// 	extism.WasmUrl{
		// 		Url: "https://github.com/extism/plugins/releases/latest/download/count_vowels.wasm",
		// 	},
		// },
		AllowedHosts: []string{"github.com", "api.github.com"},
		AllowedPaths: map[string]string{"/tmp": "/tmp"},
	}

	ctx := context.Background()
	config := extism.PluginConfig{
		EnableWasi: true,
	}
	plugin, err := extism.NewPlugin(ctx, manifest, config, []extism.HostFunction{})

	if err != nil {
		fmt.Printf("Failed to initialize plugin: %v\n", err)
		os.Exit(1)
	}

	data := []byte("Hello, World!")
	exit, out, err := plugin.Call("count_vowels", data)
	plugin.SetLogger(func(level extism.LogLevel, s string) {
		fmt.Println(level, s)
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(int(exit))
	}

	response := string(out)
	fmt.Println(response)
}
