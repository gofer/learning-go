package main

func exercise003() {
	persons := make([]Person, 10_000_000, 10_000_000)
	for i := range persons {
		persons[i] = Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		}
	}
}

/*
# 3. について

## ii. 実行にかかる時間を測定する

```
$ time ./main
./main  0.22s user 0.03s system 128% cpu 0.196 total
```

## iii. GOGC の値を変更して，影響を調べる

```
$ time GOGC=10 ./main
GOGC=10 ./main  0.21s user 0.03s system 163% cpu 0.152 total

$ time GOGC=25 ./main
GOGC=25 ./main  0.29s user 0.04s system 223% cpu 0.149 total

$ time GOGC=50 ./main
GOGC=50 ./main  0.20s user 0.04s system 158% cpu 0.150 total

$ time GOGC=200 ./main
GOGC=200 ./main  0.21s user 0.03s system 162% cpu 0.148 total

$ time GOGC=400 ./main
GOGC=400 ./main  0.21s user 0.03s system 164% cpu 0.149 total

$ time GOGC=1000 ./main
GOGC=1000 ./main  0.10s user 0.12s system 398% cpu 0.055 total

$ time GOGC=off ./main
GOGC=off ./main  0.02s user 0.02s system 89% cpu 0.042 total
```

### iv. 環境変数 GODEBUG=gctrace=1 を設定して，GC の挙動を観察する

```
$ GODEBUG=gctrace=1 GOGC=10 ./main
gc 1 @0.000s 7%: 0.026+0.18+0.026 ms clock, 0.26+0.041/0.053/0+0.26 ms cpu, 0->0->0 MB, 0 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=25 ./main
gc 1 @0.000s 13%: 0.010+61+0.52 ms clock, 0.10+0/80/89+5.2 ms cpu, 381->382->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=50 ./main
gc 1 @0.000s 15%: 0.015+126+7.6 ms clock, 0.15+0/145/71+76 ms cpu, 381->382->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=100 ./main
gc 1 @0.000s 21%: 0.004+76+9.5 ms clock, 0.042+0/97/64+95 ms cpu, 381->381->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=200 ./main
gc 1 @0.000s 21%: 0.005+119+18 ms clock, 0.055+0/134/72+181 ms cpu, 381->382->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=400 ./main
gc 1 @0.000s 20%: 0.004+83+8.6 ms clock, 0.048+0/108/77+86 ms cpu, 381->382->381 MB, 381 MB goal, 0 MB stacks, 0 MB globals, 10 P

$ GODEBUG=gctrace=1 GOGC=1000 ./main
gc 1 @0.000s 15%: 0.005+83+2.8 ms clock, 0.052+0/107/79+28 ms cpu, 381->382->381 MB, 382 MB goal, 0 MB stacks, 0 MB globals, 10 P
```
*/
