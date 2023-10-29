# shard-concurrency

シャード単位で並列実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/migration
BenchmarkMigration-8   	       6	2070194986 ns/op	 3665782 B/op	  130165 allocs/op
BenchmarkMigration-8   	       6	2023206132 ns/op	 3673504 B/op	  130195 allocs/op
BenchmarkMigration-8   	       6	2035839083 ns/op	 3666450 B/op	  130179 allocs/op
BenchmarkMigration-8   	       6	2022268576 ns/op	 3666477 B/op	  130186 allocs/op
BenchmarkMigration-8   	       6	2007865854 ns/op	 3666353 B/op	  130187 allocs/op
BenchmarkMigration-8   	       6	2006155417 ns/op	 3680954 B/op	  130209 allocs/op
PASS
ok  	github.com/TakumaKurosawa/migration	146.245s
```
