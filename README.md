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
* no mock framework (yet)

## How to Build

within your `$GOPATH/src`

```
git clone https://github.com/ckolumbus/golangRestApiExampleWithDependencyInjection.git
cd golangRestApiExampleWithDependencyInjection
go get -t ./...
sqlite3 db.sqlite < db/schema.sql
go run main.go
```

# Todo

- [ ] improve documentation
- [ ] integrate mock framework
- [ ] integrate initial db/schema creation 
- [ ] investigate possible use of an ORM


## References
### Structure

c.f. [code guidlines](https://golang.org/doc/code.html)

### Naming conventions

* https://golang.org/doc/effective_go.html#names


### Configure Dev Env

* https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension#gopath-from-goinfergopath-setting


### Dependency handling

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

#### Mocks

* https://github.com/golang/mock
* https://github.com/DATA-DOG/go-sqlmock

### Debug 

* https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code

### other REST examples

 * https://github.com/kyawmyintthein/golangRestfulAPISample
 * https://github.com/emicklei/go-restful/blob/master/examples

