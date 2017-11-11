# Shuttle Command Line Client (GoLang version)

shuttle-cli is a simple cli SSH shortcut menu for macOS inspired by https://github.com/hendricius/shuttle-cli

Shuttle GUI version is not required, all you need is .shuttle.json file in your home directory.

Tested with macOS, Linux

```
# Show the configured hosts
$ shuttle ls
+-----+----------------------------------------------------------------------+
| #   | NAME                 | COMMAND                                       |
+-----+----------------------+-----------------------------------------------+
| 0   | Main Item            | ssh username@dev.example.com                  |
| 1   | Submenu Item #1.1    | ssh username@blog.example.com                 |
+-----+----------------------+-----------------------------------------------+
 
# Connect to the host by index
$ shuttle 0
Last login: Thu Dec 15 16:26:45 2016 from 127.0.0.1
[username@dev.example.com ~]#  
 
# Connect to the host by name
$ shuttle --name Submenu
Last login: Thu Dec 15 16:27:54 2016 from 127.0.0.1
[username@blog.example.com ~]# 

```


## Installation

Required Prerequisites:

- `Go`
- `dep` Go dependency management tool

Make sure you don't forget to setup `GOPATH` and add the workspace's bin subdirectory to your `PATH`.
```
$ export PATH=$PATH:$GOPATH/bin
```

Download repository and install dependencies:
```
$ go get -v github.com/slyzerwar/shuttle-cli
 
# Install dependencies
$ cd $GOPATH/src/github.com/slyzerwar/shuttle-cli/
$ dep ensure
```

To install shuttle-cli to be accessible from anywhere on your system:
```
$ go build -i -o $GOPATH/bin/shuttle -ldflags "-s -w" github.com/slyzerwar/shuttle-cli
```
That's it. Now you can use shuttle-cli:
```
$ shuttle

shuttle-cli is a simple cli SSH shortcut menu for macOS
 
Usage:	 shuttle <index>
	 shuttle --name <name>
	 shuttle <command>
 
<name>	 name of the configured host
<index>	 index of the configured host
 
Commands:
 
ls	 List hosts
e	 Edit shuttle configuration
 
```

If you already have shuttle installed, make sure you have `~/.shuttle.json`. However if you're running on Linux or macOS but don't want to install shuttle (GUI version), you can create shuttle config by using this example: https://github.com/fitztrev/shuttle/blob/master/tests/.shuttle.json

## Usage

Start shuttle by typing:

```
$ shuttle

shuttle-cli is a simple cli SSH shortcut menu for macOS
 
Usage:	 shuttle <index>
	 shuttle --name <name>
	 shuttle <command>
 
<name>	 name of the configured host
<index>	 index of the configured host
 
Commands:
 
ls	 List hosts
e	 Edit shuttle configuration
 
```

