# 🌀 Gortex

> **The Intelligent Neural Cortex for Pure Go Full-Stack Development.**  
> 100% Pure Go • Zero JavaScript • Native PyTorch/ONNX AI Inference • Powered by Gin & Vugu

[![Go Reference](https://go.dev)](https://go.dev)
[![License: MIT](https://shields.io)](https://opensource.org)

---

**Gortex** 是一個極致硬核的純 Go 語言全棧 Web 框架。

它將 **Vugu** (WebAssembly 前端) 與 **Gin** (高效能後端) 進行底層編譯融合，並原生內建 **PyTorch/ONNX** 矩陣推理引擎。

開發者無需配置 Node.js、零 JavaScript、更不用搭建複雜的 Python 伺服器，即可用 100% 純 Go 打造出具備深度學習能力的網頁應用。

## ✨ 核心特性

*   **⚡ 100% Pure Go**: 前端直接編譯為 WebAssembly (Wasm) 於瀏覽器運行，徹底告別 NPM 複雜配置。
*   **🧠 Native AI Inference**: 內建 ONNX Runtime / LibTorch 記憶體綁定，在 Go 進程內以 C++/CUDA 速度完成極速推理。
*   **🔄 Zero-API Typed RPC**: 前後端 100% 共享 Go Struct，編譯期型別安全，欄位變更時編譯立刻報錯。
*   **🔥 Hot-Reloading Dev Server**: 內建自研監聽器，自動執行 `vugu-gen` 與 `wasm` 編譯，並透過 WebSocket 即時刷新。
*   **📦 Single Binary**: 網頁 UI、Wasm 靜態資源與後端 API 引擎最終打包進單一執行檔，一鍵部署，極低記憶體佔用。

---

## 📂 專案目錄結構

Gortex 採用「約定大於配置」的現代單一倉庫 (Monorepo) 結構：

```text
my-gortex-app/
├── app/                  # 🎨 前端 Vugu WebAssembly 組件 (root.vugu)
├── server/               # ⚙️ 後端 Gin 業務邏輯與 API 路由 (main.go)
├── models/               # 🧠 PyTorch / ONNX 模型存放區 (.onnx / .pt)
├── shared/               # 🤝 前後端 100% 共享的 Go 資料型別 (types.go)
├── go.mod
└── gortex.toml           # 框架核心配置文件
```

---

## 💻 核心程式碼範例

### 1. 共享型別定義 (`shared/types.go`)
```go
package shared

type PredictRequest struct {
	RawImage []byte `json:"raw_image"`
}

type PredictResponse struct {
	ClassLabel string  `json:"class_label"`
	Confidence float64 `json:"confidence"`
}
```

### 2. 後端核心與 AI 推理 (`server/main.go`)
```go
package main

import (
	"context"
	"net/http"
	"gortex"
	"my-gortex-app/shared"
	"://github.com"
)

func main() {
	// 初始化 Gortex 核心組件
	app := gortex.New(gortex.Config{
		FrontendDir: "./app",
		WasmOutName: "main.wasm",
	})

	// 自動執行前端 Wasm 背景編譯管道
	if err := app.CompileWasm(); err != nil {
		panic(err)
	}

	// 載入 PyTorch 導出的 ONNX 模型至系統記憶體
	visionBrain := app.NewAIEngine("models/resnet50.onnx")

	// 註冊極速 AI 推理 API
	app.API.POST("/analyze", func(c *gin.Context) {
		var req shared.PredictRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
			return
		}

		// 呼叫底層 C++/CUDA 進行權重矩陣運算
		result := visionBrain.Predict(context.Background(), req.RawImage)
		c.JSON(http.StatusOK, result)
	})

	// 啟動全棧服務 (同時託管 WebAssembly 前端與 API)
	app.Run(":3000")
}
```

---

## 🛠️ 生產環境編譯

準備部署時，Gortex 會將所有前端代碼編譯為 Wasm 並利用 `//go:embed` 打包進單一執行檔：

```bash
gortex build --release
```
編譯完成後，你只會得到一個獨立的二進位檔案，直接丟上 **Ubuntu、Docker 或 AWS** 即可完美運行。

## 📄 開源許可證

本專案基於 **MIT License** 許可證開源 - 詳見 [LICENSE](LICENSE) 檔案。
