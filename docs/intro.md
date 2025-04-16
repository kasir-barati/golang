# Install Go

```cmd
# Manjaro linux
sudo pacman -Syu go
```

# Structure of Go Code

- Packages:
  - A bunch of go files :smile:.
- Modules:
  - A bunch of packages :smile:.
  - So basically when we initiate a new project we are creating a new module.

```cmd
go mod init github.com/kasir-barati/golang
```

> [!TIP]
>
> The name of module usually is the URL of a GitHub repository. But we can use any name we want, i.e. `git mod init test`.

![go mod init result is a go.mod file](./assets/go-mod-res.png)

Now you can start creating new packages:

```cmd
code src/utils/main.go
```

- This is a special package, it is the entry point of the program.
- It is where the compiler starts executing the code from.
- You must define a function called `main` in this package. **This is mandatory**.

Now it's time to build and run the program:

```cmd
go build src/utils/main.go
```

- This will create a binary file called `main` at the root of the project.
- I gitignore the output of this command.

Now let's run the program:

```cmd
./main
```

> [!TIP]
>
> You can also run the program with the `go run src/utils/main.go` command.
