# worker-pattern

ワーカーパターンで実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/migration
BenchmarkMigration-8   	       6	  26054167 ns/op	 3746630 B/op	  130882 allocs/op
BenchmarkMigration-8   	       6	  26562875 ns/op	 3699801 B/op	  130716 allocs/op
BenchmarkMigration-8   	       6	  26428771 ns/op	 3685982 B/op	  130684 allocs/op
BenchmarkMigration-8   	       6	  26312868 ns/op	 3696168 B/op	  130706 allocs/op
BenchmarkMigration-8   	       6	  26898944 ns/op	 3698601 B/op	  130749 allocs/op
BenchmarkMigration-8   	       6	  27957222 ns/op	 3709990 B/op	  130793 allocs/op
PASS
ok  	github.com/TakumaKurosawa/migration	1.873s
```
