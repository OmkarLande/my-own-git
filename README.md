# My Own Git
Version Control clone in golang
## Task Done:
- Initialize a git repository
- Writing Blob Object
- Reading Blob Object

## Task Remaining:
- Writing Tree Object
- Reading Tree Object
- Create a commit 
- Clone a repository

### How to run:
```shell
    ./mygit.sh <command> <arg1> <arg2> ...
```
### Commands:
- init: Initialize a git repository
```shell
    ./mygit.sh init
```
- write: Write a Blob object
```shell
    ./mygit.sh hash-object -w <file(name)>
```

- read: Read a Blob object
```shell
    ./mygit.sh cat-file -p <blob-id>
```