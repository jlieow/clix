# clix
clix stands for cli extender.

Designed to enhance and extend the functionality of command-line interfaces (CLI). CliX empowers developers, system administrators, and power users by allowing them to define, 
customize, and chain commands with ease.

Used `cobra-cli init` to initialise a CLI structure. If you have not created a path to the go executables, run `~/go/bin/cobra-cli init`.

Add GOPATH/bin directory to your PATH environment variable via `.bash_profile`, `.bashrc` or `.zshr` so you can run Go programs anywhere with `export PATH=$PATH:$(go env GOPATH)/bin`.

To install and use, perform the following commands:
```
go build
go install
clix
```

~/go/bin


# symlinks

sudo ln -s ~/go/bin/clix /usr/local/bin/x
sudo ln -s ~/go/bin/clix /usr/local/bin/t
ln -s ~/go/bin/clix /usr/bin/t
sudo alias x='~/go/bin/clix'