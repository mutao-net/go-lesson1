# go-lesson1
- Go入門その1
- 簡単な計算機を作ってみる。

## 進め方
以下のように実装を進めること
-  1. まずは入力値を受け取るformを作成
-  2. 引数を2個取って足し算ができるようになること
-  3. 引数を2個取って引き算ができるようになること
-  4. 掛け算、割り算の実装

## 注意
- 3rdパーティ製のFWを使わずにGoのデフォルトで入っているパッケージで実装を進めること
- CSSは単純なもの。もしくは不要。

## ゴール
- 拘らずに「進め方」に書いたものができていればいい
  - まずは書き方に慣れたい。
  -  まずはガシガシ書いてそのあとで品質を上げること。

- http://localhost:8081/index/
  
## 感想
- 簡単な文法のおさらいをできたけど、特に拡張性のあるものでもないので品質後回しでざっくり実装した。
- funcMapとかその辺を無理やり実装したかった
- 次の課題を作るときはもっとちゃんとしたお題にする。


```
$ go test -v ./...
?       go-lesson1      [no test files]
=== RUN   TestAverage
--- PASS: TestAverage (0.00s)
PASS
ok      go-lesson1/mylib        0.152s
?       go-lesson1/mylib/hello  [no test files]
```]