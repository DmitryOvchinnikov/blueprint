# Blueprint: blueprint with Kubernetes

## Installation

To clone the project, create a folder and use the git clone command. Then please read the [Makefile](Makefile) file to learn how to install all the tooling and docker images.

```
$ cd $HOME
$ mkdir code
$ cd code
$ git clone https://github.com/dmitryovchinnikov/blueprint or git@github.com:dmitryovchinnikov/blueprint.git
$ cd blueprint
```

## Create Your Own Version

If you want to create a version of the project for your own use, use the new gonew command.

```
$ go install golang.org/x/tools/cmd/gonew@latest

$ cd $HOME
$ mkdir code
$ cd code
$ gonew github.com/dmitryovchinnikov/blueprint github.com/mydomain/myproject
$ cd myproject
$ go mod vendor
```

Now you have a copy with your own module name. Now all you need to do is initialize the project for git.
