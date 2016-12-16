# Shuttle Command Line Client (GoLang version)

shuttle-cli is a simple cli SSH shortcut menu for macOS inspired by https://github.com/hendricius/shuttle-cli

Shuttle GUI version is not required, all you need is .shuttle.json file in your home directory.

Tested with macOS, Linux
```
$ shuttle ls
+--------+----------------------------------------------------------------------+
| NUMBER | NAME                 | COMMAND                                       |
+--------+----------------------+-----------------------------------------------+
| 0      | Main Item            | ssh username@dev.example.com                  |
| 1      | Submenu Item #1.1    | ssh username@blog.example.com                 |
+--------+----------------------+-----------------------------------------------+

$ shuttle c 1
Last login: Thu Dec 15 16:26:45 2016 from 127.0.0.1
[username@dev.example.com ~]# 

```


## Installation

`GoLang` is required to build the source code. If you don't have it, have a look at https://golang.org/doc/install

Make sure you don't forget to setup `GOPATH` and add the workspace's bin subdirectory to your `PATH`.
```
$ export PATH=$PATH:$GOPATH/bin
```

Download repository to the proper place:
```
$ go get -v github.com/slyzerwar/shuttle-cli
$ cd $GOPATH/src/github.com/slyzerwar/shuttle-cli
```

To install shuttle-cli to be accessible from anywhere on your system:
```
$ go install
```
That's it. Now you can use shuttle-cli:
```
$ shuttle
Usage:
        shuttle ls (Show hosts)
        shuttle c <number> (Connect to host)
        shuttle e (Edit shuttle configuration)
```

If you already have shuttle installed, make sure you have `~/.shuttle.json`. However if you're running on Linux or macOS but don't want to install shuttle (GUI version), you can create shuttle config by using this example: https://github.com/fitztrev/shuttle/blob/master/tests/.shuttle.json

## Usage

Start shuttle by typing:

    $ shuttle
    Usage:
        shuttle ls (Show hosts)
        shuttle c <number> (Connect to host)
        shuttle e (Edit shuttle configuration)
