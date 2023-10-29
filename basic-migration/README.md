# basic-migration

シャードごとに逐次実行するマイグレーションのサンプルです。

**※DynamoDBへのインサート部分は `time.Sleep()` を使って擬似的に処理時間を稼いでいます。**

### 実行結果

```shell
goos: darwin
goarch: arm64
pkg: github.com/TakumaKurosawa/migration
BenchmarkMigration-8   	       6	12054442208 ns/op	 3649185 B/op	  130120 allocs/op
BenchmarkMigration-8   	       6	12052751236 ns/op	 3649241 B/op	  130121 allocs/op
BenchmarkMigration-8   	       6	12017864194 ns/op	 3649593 B/op	  130127 allocs/op
BenchmarkMigration-8   	       6	12063526986 ns/op	 3649342 B/op	  130123 allocs/op
BenchmarkMigration-8   	       6	12029776667 ns/op	 3648978 B/op	  130119 allocs/op
BenchmarkMigration-8   	       6	11978084424 ns/op	 3648970 B/op	  130120 allocs/op
PASS
ok  	github.com/TakumaKurosawa/migration	505.607s

```
