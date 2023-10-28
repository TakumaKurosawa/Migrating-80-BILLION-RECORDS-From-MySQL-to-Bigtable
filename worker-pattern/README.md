# worker-pattern

ワーカーパターンで実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/worker-pattern
BenchmarkMigration-8   	       6	  26384007 ns/op	 3752388 B/op	  130793 allocs/op
PASS
ok  	github.com/TakumaKurosawa/worker-pattern	0.349s
```
