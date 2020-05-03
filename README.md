# trash
Replacement for rm command.  
The trash command reproduces the directory structure into the trash can.
# Setup
Use $TRASH_CAN_PATH to specify path.  
It is ~/.Trash in the default.
# Available Commands
## move
Move the files to the trash can.
## restore
Move the files in the trash can to the current directory.
### --all(-a)
Target all the files.
## list
Show the file names in the trash can.
### --path(-p)
Show the file paths in the trash can.
### --reverse(-r)
Show the file names in reverse order.
### --days(-d) [days]
Show the file names moved to the trash can within [days] days.
```bash
$ trash --days 1
$ trash -d 10
```
### --size(-s) [size]
Show the files with size greater than [size] size.
```bash
$ trash --size 500MB
$ trash -s 1GB
```
## size
The size of the trash can directory.
## remove
Remove a file in the trash can.
## auto-remove
Remove the files if the date and time that the file moved in the trash can exceed the specified period.
### --period(-p) [days]
The option can specify the period. It is 30 days in the default.
### .bashrc
```bash
which trash &> /dev/null && trash auto-remove
```
### .zshrc
```zsh
[[ -n ${commands[trash]} ]] && trash auto-remove
```
