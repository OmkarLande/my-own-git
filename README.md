# My Own Git
Version Control clone in golang

## Tasks Completed: 
- **Initialize a Git repository**: Set up a new repository to start tracking your project. 
- **Writing Blob Object**: Store file data as a blob object. 
- **Reading Blob Object**: Retrieve and display the content of a blob object. 

## Tasks Remaining: 
- **Writing Tree Object**: Implement the functionality to write tree objects that represent directory structures. 
- **Reading Tree Object**: Implement the functionality to read and interpret tree objects. 
- **Create a Commit**: Develop the ability to create commit objects that record changes in the repository. 
- **Clone a Repository**: Add the feature to clone an existing repository.

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