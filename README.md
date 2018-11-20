# Test initialization of packages

## Description of problem

If you run this file `main.go` you get following output:

```console
$ go run main.go 
inside root init
inside noop
values:  [0 1 2 3 4 5 6 7 8 9]
```

The initialization code in packages `one` and `two` is not called, which is a mystery!

The code structure is laid out like following:

```console
$ exa -T
.
├── main.go
├── README.md
└── root
   ├── one
   │  └── one.go
   ├── root.go
   └── two
      └── two.go
```

Packages `one` and `two` call up `AddValue` function of `root` package to add values from their `init` functions. But those are not invoked for some reason. The initialization code of pkg `one` and `two` never runs. Hence an extra value of `-1` and `-2` is never added.


## Setup dev env

```bash
mkdir -p $GOPATH/src/github.com/surajssd/study
cd $GOPATH/src/github.com/surajssd/study
git clone https://github.com/surajssd/pkgsall
cd pkgsall
go run main.go
```
