## Learn golang - level01

Aprende del primer nivel de conocimientos de go

## Golang our URL GITHUP
https://github.com/rbarec/gorb-util
https://github.com/rbarec/gorb-gotopro


## Comandos
go test
go mod init github/hello


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

```
C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello2>go test
go test
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c

 github.com/rbarec/leve01/hello2 [github.com/rbarec/leve01/hello2.test]
.\hello2.go:3:8: imported and not used: "rsc.io/quote"
FAIL    github.com/rbarec/leve01/hello2 [build failed]
```

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello2>go test
> github.com/rbarec/leve01/hello2 [github.com/rbarec/leve01/hello2.test]
.\hello2.go:3:8: imported and not used: "rsc.io/quote"
FAIL    github.com/rbarec/leve01/hello2 [build failed]

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello2>go test
--- FAIL: TestHello2 (0.00s)
    hello2_test.go:8: Hello2() = "Hello, world.", want "Hello2, world."
FAIL
exit status 1
FAIL    github.com/rbarec/leve01/hello2 3.257s


### Latest tagged 

git tag -a v1.4
git tag -a v1.4 -m "my version 1.4"
git tag   //LISTAR LOS TAGS
git tag -l *-rc*  //LISTAR SOLO LOS RC

git log --pretty=oneline   //VER LOS COMMITS
15027957951b64cf874c3557a0f3547bd83b3ff6 Merge branch 'feature'
a6b4c97498bd301d84096da251c98a07c7723e65 add update method for thing
0d52aaab4479697da7686c15f77a3d64d9165190 one more thing

git tag -a v1.2 15027957951b64cf874c3557a0f3547bd83b3ff6

In the event that you must update an existing tag, the -f FORCE option must be used.
git tag -a -f v1.4 15027957951b64cf874c3557a0f3547bd83b3ff6



## GIT Release naming conventions [[KNOW]]

Important note about release tags:

unstable(deprecated)
alpha
beta
rc  (release candidate)

OK  8.x-1.0-rc1 will appear as a valid tag, but the following will NOT:
NOT 8.x-1.0-release1 (wrong word)
NOT 8.x-1.0-rc (doesn't end in a digit)
NOT 8.x-1.0-ALPHA1 (uppercase)

More samples:  https://www.drupal.org/node/467020
Versiones sample product:  https://rubygems.org/gems/bundler/versions

alpha: Most reported errors are resolved, but there may still be serious outstanding known issues, including security issues. Project is not thoroughly tested, so there may also be many unknown bugs. There is a README.txt/README.md that documents the project and its API (if any). The API and DB schema may be unstable, but all changes to these are reported in the release notes, and hook_update_N is implemented to preserve data through schema changes, but no other upgrade/update path. Not suitable for production sites. Target audience is developers who wants to participate in testing, debugging and development of the project.

beta: All critical data loss and security bugs are resolved. If the module offers an API, it should be considered frozen, so that those using the API can start upgrading their projects. If it is an upgrade or update of a project, an upgrade/update path should be offered, and it should be possible for existing users to upgrade/update to the new version without loss of data. All documentation should be up to date. Target audience is developers who wants to participate in testing, debugging and development of the project, and developers of other projects that interfaces the project. Not generally suitable for production sites, but may be used on some production sites if the site administrator knows the project well, and knows how to handle any outstanding issues.

rc: A release candidate should only be created when all reported critical bug type issues are fixed in the project's issue queue. This tag should only be used when the developer believes that the project is ready for use on a production site. There is no official best practice for how long a project should be a release candidate before creating a official .0 release, but it is suggested that it should be out for at least a month with status set to "needs review". If something (e.g. a new critical bug is reported) makes it necessary to create a new release during this period, a new release candidate should be created and this should remain for at least a month with status set to "needs review".

```
C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello2>go list -m all
github.com/rbarec/leve01
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0	
```

# level01- Hello03 
 _gorb_gotopro[bbkt]\level01\hello3>
 
 ````
 import (
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
    return quote.Hello()
}

func Proverb() string {
    return quoteV3.Concurrency()
}

/////////// test
func TestHello3(t *testing.T) {
    want := "Concurrency is not parallelism."
    if got := Proverb(); got != want {
        t.Errorf("Proverb() = %q, want %q", got, want)
    }
}

````


C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>go list -m
github.com/rbarec/leve01

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>go list -m rsc.io/quote
rsc.io/quote v1.5.2

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>go list -m rsc.io/quote/v3
rsc.io/quote/v3 v3.1.0

#### usando comando  *go doc pkg*

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>go doc rsc.io/quote/v3
package quote // import "rsc.io/quote/v3"
Package quote collects pithy sayings.
func Concurrency() string
func GlassV3() string
func GoV3() string
func HelloV3() string
func OptV3() string

C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>go doc rsc.io/quote
package quote // import "rsc.io/quote"
Package quote collects pithy sayings.
func Glass() string
func Go() string
func Hello() string
func Opt() string
C:\0SYS\gitrb\_gorb_gotopro[bbkt]\level01\hello3>	


####  Borrando los paquetes no usados
.
