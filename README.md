# GoFuzzing

## スライド

https://docs.google.com/presentation/d/1pGA4iiyg74UsIyhM05aUovOuj0HEtqOtPW3QDXesIB8/edit#slide=id.g22ec37ca2c0_0_74

# 参考教材

公式チュートリアル(https://go.dev/doc/tutorial/fuzz) を参考に記録している。

# 環境
- go 1.20

# 学べること
- Fuzzingの書き方、実行方法など

# 手順

```
go test -fuzz FuzzReverse -fuzztime 30s -v

失敗するFuzzTestケースがあったら、fuzz/testdata/fuzz/FuzzReverse配下に記録される
```
