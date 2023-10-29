# Migrating-80-BILLION-RECORDS-From-MySQL-to-Bigtable

GopherCon 2023

Source: https://www.gophercon.com/agenda/session/1158499


## 🔍 Quick find the results

- [Basic Migration](./basic-migration/README.md)
- [Shard Concurrency](./shard-concurrency/README.md)
- [Worker Pattern](./worker-pattern/README.md)

### Benchstat result

basic-migration vs. shard-concurrency vs. worker-pattern

```shell
            │ basic-migration.out │        shard-concurrency.out         │         worker-pattern.out         │
            │       sec/op        │    sec/op      vs base               │   sec/op     vs base               │
Migration-8        12041.26m ± 1%   2022.74m ± 2%  -83.20% (p=0.002 n=6)   26.50m ± 6%  -99.78% (p=0.002 n=6)

            │ basic-migration.out │       shard-concurrency.out        │         worker-pattern.out         │
            │        B/op         │     B/op      vs base              │     B/op      vs base              │
Migration-8          3.480Mi ± 0%   3.497Mi ± 0%  +0.47% (p=0.002 n=6)   3.528Mi ± 1%  +1.37% (p=0.002 n=6)

            │ basic-migration.out │       shard-concurrency.out       │        worker-pattern.out         │
            │      allocs/op      │  allocs/op   vs base              │  allocs/op   vs base              │
Migration-8           130.1k ± 0%   130.2k ± 0%  +0.05% (p=0.002 n=6)   130.7k ± 0%  +0.47% (p=0.002 n=6)
```

shard-concurrency vs. worker-pattern

```shell
            │ shard-concurrency.out │         worker-pattern.out         │
            │        sec/op         │   sec/op     vs base               │
Migration-8           2022.74m ± 2%   26.50m ± 6%  -98.69% (p=0.002 n=6)

            │ shard-concurrency.out │         worker-pattern.out         │
            │         B/op          │     B/op      vs base              │
Migration-8            3.497Mi ± 0%   3.528Mi ± 1%  +0.89% (p=0.002 n=6)

            │ shard-concurrency.out │        worker-pattern.out         │
            │       allocs/op       │  allocs/op   vs base              │
Migration-8             130.2k ± 0%   130.7k ± 0%  +0.42% (p=0.002 n=6)

```

## 🛫 Pre-requisites

- Docker installed
- Docker Compose installed
- Go runtime installed

## 💻 Setup

Just type this command!

```shell
make setup
```

## 🏃‍Run

```shell
# Run all benchmarks
make bench/all

# Run basic-migration benchmark
make bench/basic-migration

# Run shard-concurrency benchmark
make bench/shard-concurrency

# Run worker-pattern benchmark
make bench/worker-pattern
```

## 🌳 Environment

| MySQL                | Count  |
|----------------------|:------:|
| MySQL logical shards |   6    |
| Records per shard    | 10,000 |

## References

- [Migrating 80 Billion Records from mySQL to Bigtable](https://www.gophercon.com/agenda/session/1158499)
- [[Go] benchstat/go tool traceコマンドをつかったベンチマークの可視化](https://budougumi0617.github.io/2020/12/04/goroutine_tuning_with_benchmark_benchstat_trace/)
- [go tool traceでgoroutineの実行状況を可視化する](https://yuroyoro.hatenablog.com/entry/2017/12/11/192341)
