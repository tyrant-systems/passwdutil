# passwdutil

* * *

## about

simple password utility. supports passwords
of the following type:

  - `bcrypt`

## usage

compile as a binary, ie:

```
~$ go build -v -o password-manager
```

or compile using docker:

```
~$ docker run --rm                        \
	-v ${current_dir}:/build              \
	-w /build golang:${golang_version}    \
		go build -v -o ${output_binary}
```

run.
