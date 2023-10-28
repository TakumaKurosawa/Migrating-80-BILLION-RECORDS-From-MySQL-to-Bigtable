# Migrating-80-BILLION-RECORDS-From-MySQL-to-Bigtable

GopherCon 2023

Source: https://www.gophercon.com/agenda/session/1158499


## 🔍 Quick find the results

- [Basic Migration](./basic-migration/README.md)
- [Shard Concurrency](./shard-concurrency/README.md)
- [Worker Pattern](./worker-pattern/README.md)

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
