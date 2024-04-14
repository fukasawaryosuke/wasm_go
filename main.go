package main

import (
    "math/rand"
    "syscall/js"
)

func main() {
    c := make(chan struct{}, 0)

    // WebAssembly用の関数を登録
    js.Global().Set("calculateAverage", js.FuncOf(func(this js.Value, p []js.Value) interface{} {
        // 乱数を生成し、その平均を計算する
        var sum float64
        for i := 0; i < 1000; i++ {
            sum += rand.Float64()
        }
        return sum / 1000
    }))

    <-c // 無限ループでプログラムが終了しないようにする
}
