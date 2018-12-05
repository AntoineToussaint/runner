# runner

Run with:

```go run main.go run /bin/sh```

## MacOS

No permission, no chroot, no cgroup.

## Linux

When ran inside Docker ubuntu:

"""
panic: fork/exec /proc/self/exe: operation not permitted

goroutine 1 [running]:
github.com/hygge.io/hygge.io-core/runner.must(...)
	/Users/antoinetoussaint/go/src/github.com/hygge.io/hygge.io-core/runner/linux.go:63
github.com/hygge.io/hygge.io-core/runner.Run()
	/Users/antoinetoussaint/go/src/github.com/hygge.io/hygge.io-core/runner/linux.go:26 +0x2ed
main.main()
	/Users/antoinetoussaint/go/src/github.com/hygge.io/hygge.io-core/main.go:12 +0x69
  """
