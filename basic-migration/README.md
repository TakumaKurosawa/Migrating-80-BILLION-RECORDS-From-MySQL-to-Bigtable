# basic-migration

シャードごとに逐次実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/basic-migration
BenchmarkMigration-8   	       6	12258417403 ns/op	 3650237 B/op	  130123 allocs/op
PASS
ok  	github.com/TakumaKurosawa/basic-migration	86.400s
```
