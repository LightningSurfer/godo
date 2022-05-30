# GODO
Welcome to GODO, a CLI TODO app to track things done and things left to go do.

Functional, but not necessarily optimal. Use at your own risk.

"Do. Or do not. There is no try." - Yoda

### Tech
Written in Go. Uses the CLI app library, [Cobra](https://github.com/spf13/cobra).

### Installation
To install it, run:
```
go get github.com/LightningSurfer/godo
```

This program saves a copy of your TODO list in a .todo file in your home directory by default.

If you'd like to store this in some other directory within your home directory, you can `.env` file, and set TODOS_LOCATION to where ever you'd like.

Ex:
```
TODOS_LOCATION="SomeOtherDir/.some_other_filename"
```
