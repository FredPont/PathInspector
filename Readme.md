

#  Path Inspector
Path Inspector is a software to compute recursively file path sizes and list all the path that are above a limit.


# Quick start
- the software can be run in interactive mode or command line
- interactive mode is the default mode : just start the software and answer the questions
- to have absolute path, enter the absolute path of the directory to scan

- command line mode :
```
Usage 
  -d string
        path (default ".")
  -i    Interactive mode. Important syntax is -i=false (default true)
  -l int
        maximal path length (default 255)
```

example : 

```
# non interactive mode, list path >= 4, in the /home/fred/test directory
pathin-x86_64_linux.bin -i=false -l 4 -d /home/fred/test
```

# ScreenShots
![CLI](src/images/screenshot.png)
