# shard-concurrency

シャード単位で並列実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/shard-concurrency
BenchmarkMigration-8   	       6	2007694299 ns/op	 3661337 B/op	  130181 allocs/op
PASS
ok  	github.com/TakumaKurosawa/shard-concurrency	24.212s
```
