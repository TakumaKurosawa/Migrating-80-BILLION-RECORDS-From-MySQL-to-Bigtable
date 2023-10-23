# worker-pattern

ワーカーパターンで実行するマイグレーションのサンプルです。

### 実行結果

※DynamoDBへ実際にInsertするのではなく、1msかかるものとして、 `time.Sleep()` を導入した結果になります。

```shell
-------- Start migration --------
hashdb-1 is started!!
hashdb-2 is started!!
hashdb-3 is started!!
hashdb-4 is started!!
hashdb-5 is started!!
hashdb-6 is started!!
Worker-0 is started!!
Worker-1 is started!!
Worker-2 is started!!
Worker-3 is started!!
Worker-4 is started!!
Worker-5 is started!!
... 省略
Worker-96 is started!!
Worker-97 is started!!
Worker-98 is started!!
Worker-99 is started!!
hashdb-6 is done!
hashdb-4 is done!
hashdb-3 is done!
hashdb-5 is done!
hashdb-2 is done!
hashdb-1 is done!
-------- Finish migration --------
Migration took 104.160209ms
```
