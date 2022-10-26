# Tennis Scoring System

## Prerequisites

- Go 1.18.5 or later

### To run in Go

run

```sh
go run main.go
```

### To run via built binary file

run

```sh
./tennis_scoring_system
```

### Example input

**Please input tennis players such as 'Lisa VS Jennie'. The 1st player will be the one who serve first.**

```string
Lisa VS Jennie
```

**Please input the tennis match data.**

```string
[
    Lisa, Lisa, Jennie, Lisa, Lisa,
    Lisa, Jennie, Lisa, Lisa, Jennie, Jennie, Lisa,
    Lisa, Jennie, Lisa, Jennie, Jennie, Lisa, Jennie, Jennie,
    Lisa, Lisa, Lisa, Lisa,
    Lisa, Lisa, Lisa, Lisa,
    Lisa, Lisa, Lisa, Lisa,
    Lisa, Lisa, Lisa, Lisa,
]
```

**Result**

```string
Lisa Serve 1-0
15:0, 30:0, 30:15, 40:15BP

Lisa Serve 2-0
15:0, 15:15, 30:15, 40:15BP, 40:30BP, 40:40

Lisa Serve 2-1
15:0, 15:15, 30:15, 30:30, 30:40BP, 40:40, 40:ABP

Lisa Serve 3-1
15:0, 30:0, 40:0BP

Lisa Serve 4-1
15:0, 30:0, 40:0BP

Lisa Serve 5-1
15:0, 30:0, 40:0BP

Lisa Serve 6-1
15:0SP, 30:0SP, 40:0BPSP
```
