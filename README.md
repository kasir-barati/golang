# 1. Install go

```cmd
# Manjaro linux
sudo pacman -Syu go
# mac
brew install go
```

# 2. What the heck is go.mod?

- `go.mod` in Go:
  - Serves as a module file to manage dependencies.
  - Tracks the external modules your project depends on.
  - Specifies the name of your module and the modules it relies on.
- `package.json` in Node.js:
  - Manages project dependencies for Node.js projects.
  - Lists the packages/modules your project depends on.
  - Specifies metadata about your project.
- `requirements.txt` in Python:
  - Lists the Python packages required for a project.
  - Specifies the versions of the packages needed.

```cmd
go mod init github.com/kasir-barati/my-golang-journey
```

# 3. Execute it

```cmd
go run .
```

# Wanna use new module or 3rd party package?

Use it and then just do a `go mod tidy` :grin:

# [Should I commit my 'go.sum' file as well as my 'go.mod' file?](https://stackoverflow.com/q/53837919/8784518)
