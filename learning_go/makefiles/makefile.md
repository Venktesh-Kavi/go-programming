## MakeFile Notes

### Make file syntax

``` makefile
target: prerequisites
    command
    command
    command    
```

* commands are indented with tabs
* targets can be a file (executable or object file) or name of an action (eg. clean)
* command is interpreted by a shell to get executed. By default ths shell if /bin/sh shell.
* prerequisites are the the dependency files used to create the target.

### Macros

* macros help making modifications to make files easy by avoiding repeating text entries.
* Eg.., here is CC is used as a macro for g++.
* Internal Macros
    * make has a pre-defined set of macros. `make -p`. Lists all internal macros in make
    * `@`

### Sample MakeFile

``` makefile
CC=g++

all: prog

prog: main.o, factorial.o hello.o
    $(CC) main.o factorial.o hello.o
main.o: main.cpp
    $(CC) ($CFLAGS) main.cpp
factorial.o: factorial.cpp
    $(CC) ($CFLAGS) factorial.cpp
hello.o: hello.cpp
    $(CC) ($CFLAGS) hello.cpp
clean:
    rm -rf *.o    
```

* Running plain `make` runs the first target (which is the default) `prog`.
* Make runs `prog` only if prog does not exist or if `main.o/factorial.o/hello.o` if any of this is
  new.

### Phony Targets

* Phone targets do not have file associated with them. They are just a list of commands.


* make runs the `all` specified target if not target is provided while providing `make`. `all` can
  consist of all the targets we want to build.
* `@` after a command suppresses make from printing the command. Typically make prints the command
  along with the output. Eg.., `@echo`.