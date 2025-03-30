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

	// ランダムに値を返す
	for _, salutation = range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}
