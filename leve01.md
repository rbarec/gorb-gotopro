## Learn golang - level01

Aprende del primer nivel de conocimientos de go

## Golang our URL GITHUP
https://github.com/rbarec/gorb-util
https://github.com/rbarec/gorb-gotopro


## Comandos usados en este nivel
go test
go mod init github/hello
go list -m all
go list -m rsc.io/quote/v3
go doc rsc.io/quote
go doc rsc.io/quote/v3
go mod tidy

##  my practices
https://blog.golang.org/using-go-modules

make  folder *hello*,  *hello.go* and *hello_test.go*
haga CMD en carpeta hello y ejecutar comando
go test

```
C:\0SYS\gitexternal\goModulesExamples\hello>go test
PASS
ok      _/C_/0SYS/gitexternal/goModulesExamples/hello   3.235s
```

vamos a crear un modulo sobre esta carpeta:
go mod init example.com/hello


```
go mod init github.com/rbarec/gorb-gotopro/leve01/hello
go: creating new go.mod: module github.com/rbarec/gorb-gotopro/leve01/hello

FILE: go.mod **************
module github.com/rbarec/gorb-gotopro/leve01/hello


***************************EOF

```


## Adding a dependency

## FAQ - Software Dependency Problem [[KNOW]]
Note that while the go command makes adding a new dependency quick and easy, it is not without cost. Your module now literally depends on the new dependency in critical areas such as correctness, security, and proper licensing, just to name a few. For more considerations, see Russ Cox's blog post, *�Our Software Dependency Problem.�*
URL> https://research.swtch.com/deps  {{CANDIDATE}}


## FAQ - GIT Golang  pseudo-version [[KNOW]]
De todas estas version la (1) es una pseudo version. URL para aprender mas URL> https://golang.org/cmd/go/#hdr-Pseudo_versions.  {{CANDIDATE}}
(1) golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
(2)rsc.io/quote v1.5.2
(3)rsc.io/sampler v1.3.0


## FAQ - 

