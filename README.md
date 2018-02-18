GO REST API service
===================

## Design

* two layered micro service 
  * web api controller logic 
  * persistence layer
* Data transfer object definition for domain entity data representation
* using dependency injection for each layer for testability
* using [echo](https://echo.labstack.com/) as web api framework
* no DI framwork, manual injection
* [gomock][gomock] with mockgen as mocking framework
* test/coverage report (and much more) integrated via build script
* use DATA-DOG/go-sqlmock for persistence tests
* code structure according to [thockin/go-build-template](https://github.com/thockin/go-build-template)

## How to Build

set your `$GOPATH` (applies to unixoid OS', for windows replace all 
`$ENVVARNAME` whith `%ENVVARNAME%` )

get and build project plus build & test dependencies

```
go get -t -v github.com/ckolumbus/golangRestApiExampleWithDependencyInjection
```

## How to run Tests

```
go test -v github.com/ckolumbus/golangRestApiExampleWithDependencyInjection/...
```

## How to Run 

No special setup needed, database and schema are created automatcially in `./db.sqlite`

```
cd $GOPATH/bin
./golangRestApiExampleWithDependencyInjection
```

## Using MAGE

[`mage`][mage] is a go based task runner which implements
an internal DSL define tasks and dependencies.

Install mage by running

```
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

After this you can get a list of possbile targets with

```
cd $GOPATH/src/github.com/ckolumbus/golangRestApiExampleWithDependencyInjection
$GOPATH/bin/mage
$GOPATH/bin/mage -v check
$GOPATH/bin/mage -v install
```

Output
```
Targets:
  check               Run tests and linters
  checkVendor         verifies that vendored packages match git HEAD
  fmt                 Run gofmt linter
  install             binary
  lint                Run golint linter
  service             Build binary
  serviceNoGitInfo    Build Service without git info
  serviceRace         Build binary with race detector enabled
  test                Run tests
  test386             Run tests in 32-bit mode
  testCoverHTML       Generate test coverage report
  testRace            Run tests with race detector
  vendor              Install Go Dep and sync vendored dependencies
  vet                 Run go vet linter
```

# Todos

- [x] create build script: use [`mage`][mage]
- [ ] improve documentation
- [x] integrate mock framework : use [gmock][gomock]
- [x] integrate initial db/schema creation 
- [ ] investigate possible use of an ORM
- [x] ~~investigate seperation of tests (e.g. controller) into own package and/or directory~~ 
      according to all best practices on the net this seperation should not be done
- [ ] add mock generation to build script (example: `$GOPATH/bin/mockgen -source pkg/employee/persistence/IPersistEmployee.go -destination controllers/PersistEmployeeMock_test.go -package controllers`)
- [ ] logging concept
- [ ] handling arrays: use elments with or without "*"?


## References
### Structure

c.f. [code guidlines](https://golang.org/doc/code.html)

### Naming conventions

* https://golang.org/doc/effective_go.html#names


### Configure Dev Env

* https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension#gopath-from-goinfergopath-setting


### Handling of external Dependency 

Install all dependencies,  `-t` includes test dependencies

```go get -t ./...```

* https://coderwall.com/p/arxtja/install-all-go-project-dependencies-in-one-command



### Build


### Tests

according to the (test structure][https://golang.org/doc/code.html#Testing] definition the
test file are located next to the production code files. 

TODO: search for best practices, maybe seperate test from production code


* https://golang.org/pkg/testing/
* https://golang.org/pkg/testing/#hdr-Subtests_and_Sub_benchmarks

### Dependency Injection / Inversion of Control

* [karlkfi/inject][karlkfi/inject]
  **NO** change whatsover on existing code when doing ctor injection already
* [facebookgo/inject][facebookgo/inject]
  most references one, but IMHO quite some ceremony needed
* [magic003/alice][magic003/alice]
  somewhat in between the two above


#### Mocks

* https://github.com/golang/mock
* https://github.com/DATA-DOG/go-sqlmock

### Debug 

* https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code

### other REST examples

 * https://github.com/kyawmyintthein/golangRestfulAPISample
 * https://github.com/emicklei/go-restful/blob/master/examples

## JSON handling

* [how to de-/serialze json with correct quoting](http://goinbigdata.com/how-to-correctly-serialize-json-string-in-golang/)

## Task runner

The goal was to have a `go` based task runner to stay whitin one technology, my personal
preference goes towards internal DSL, i.e. the task description is a go file itself.

There is quite a number of task runners fo `go`, most are [yaml](http://yaml.org/) based, 
some are `go` based.

### YAML

  - [realize](https://github.com/tockins/realize) backed by commercial company
  - [go-task](https://github.com/go-task/task)  with a nice shell abstraction for cross platform scripting
  - [godo](https://github.com/go-godo/godo) since 3yrs no update
  - [myke](https://github.com/goeuro/myke)
  - [zeus](https://github.com/dreadl0ck/zeus)
  - [tusk](https://github.com/rliebz/tusk)

### GO

  - [grift](https://github.com/markbates/grift)
  - [mage][mage]


[gomock]: https://github.com/golang/mock
[mage]: https://github.com/magefile/mage
[karlkfi/inject]: https://github.com/karlkfi/inject
[facebookgo/inject]: https://github.com/facebookgo/inject
[magic003/alice]: https://github.com/magic003/alice
 