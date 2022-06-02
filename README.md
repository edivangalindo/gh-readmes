# gh-readmes
A tool to download readmes 📃

Pre-requisites:

* You need to config an env called GH_AUTH_TOKEN with your personal access token, to do the requests

How to use:

```bash
cat repos.txt | gh-readmes
```
If you need use this inside a toolchain, you can use the flag -v to print the content of readme.

Eg:

```bash
cat repos.txt | gh-readmes -v | ag -i "[0-9a-zA-Z]{40}" -o --nofilename
```

After that, readmes from each of the projects will be downloaded in the **readmes** folder, which will be created in the same location as the tool's execution.

## Installation

First, you'll need to [install go](https://golang.org/doc/install).

Then run this command to download + compile gh-readmes:
```
go install github.com/edivangalindo/gh-readmes@latest
```

You can now run `~/go/bin/gh-readmes`. If you'd like to just run `gh-readmes` without the full path, you'll need to `export PATH="/go/bin/:$PATH"`. You can also add this line to your `~/.bashrc` file if you'd like this to persist.
