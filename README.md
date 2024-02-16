# go-with-tests
Learning how to test golang code with https://quii.gitbook.io/learn-go-with-tests/. 

### writing test
Writing a test is just like writing a function, with a few rules

* It needs to be in a file with a name like `xxx_test.go`
* The test function must start with the word `Test`
* The test function takes one argument only `t *testing.T`
* In order to use the `*testing.T` type, you need to `import "testing"`, like we did with `fmt` in the other file

For now, it's enough to know that your `t` of type `*testing.T` is your "hook" into the testing framework so you can do things like `t.Fail()` when you want to fail.

### Go doc

Another quality of life feature of Go is the documentation. You can launch the docs locally by running `godoc -http :8000`. If you go to [localhost:8000/pkg](http://localhost:8000/pkg) you will see all the packages installed on your system.

The vast majority of the standard library has excellent documentation with examples. Navigating to [http://localhost:8000/pkg/testing/](http://localhost:8000/pkg/testing/) would be worthwhile to see what's available to you.

If you don't have `godoc` command, then maybe you are using the newer version of Go (1.14 or later) which is [no longer including `godoc`](https://golang.org/doc/go1.14#godoc). You can manually install it with `go install golang.org/x/tools/cmd/godoc@latest`.

