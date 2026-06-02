package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/vugu/vugu/simplehttp"
)

func main() {
	r := gin.Default()

	appDir, err := filepath.Abs("./app")
	if err != nil {
		panic(err)
	}

	vuguHandler := simplehttp.New(appDir, true)

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Goflare backend",
			"status":  "ok",
		})
	})

	r.GET("/wasm_exec.js", func(c *gin.Context) {
		wasmExecPath, err := findWasmExecPath(runtime.GOROOT())
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.File(wasmExecPath)
	})

	r.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusOK)
		vuguHandler.ServeHTTP(c.Writer, c.Request)
	})

	if err := r.Run(":3000"); err != nil {
		panic(err)
	}
}

func findWasmExecPath(goRoot string) (string, error) {
	candidates := []string{
		filepath.Join(goRoot, "lib", "wasm", "wasm_exec.js"),
		filepath.Join(goRoot, "misc", "wasm", "wasm_exec.js"),
	}

	for _, candidate := range candidates {
		_, err := os.Stat(candidate)
		if err == nil {
			return candidate, nil
		}
		if !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("stat %q: %w", candidate, err)
		}
	}

	return "", fmt.Errorf("wasm_exec.js not found under %q", goRoot)
}
