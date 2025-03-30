package main

import (
	"fmt"
	"sync"
)

func main() {
	// welcome を返す
	// スコープの中で使われている変数をさらに小さいスコープ（ゴールーチン）で読み出す場合はアドレスを参照
	salutation := "hello"
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
	fmt.Println()

	// forループがゴールーチンよりも先に終わる
	// 最後にsalutationに保持されている値の good day を返す
	for _, salutation = range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	fmt.Println()

	// コピーした値をクロージャーに渡せば大丈夫
	for _, salutation = range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
	fmt.Println()
}
