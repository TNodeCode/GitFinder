# GitFinder

A tool that finds all Git projects on your machine.

Executables can be found in `/dist`.

## How to build

Prerequisites:

- Go is installed
- You have a bash terminal (i.e. Git Bash for Windows, standard terminal on Mac and Linux)

Go to the project's root directory, open a bash terminal and run the following command:

```
$ bash ./scripts/build.sh
```

You will now find the executables in the `/dist` directory.

## How To run

Scenario: You're on your Mac / PC and you want to find Git projects in your home directory.

Open a console, navigate to gitfinder binary and execute the following command:

Mac / Linux:

```bash
$ ./gitfinder -dir "$HOME"
```

Windows Powershell:

```
$ ./gitfinder.exe -dir "HOME"
```

If GitFinder takes too much time you can limit the number of subdirectories that are checked:

```
$ ./gitfinder -dir "$HOME" -depth 3
```

At the end GitFinder prints all directories to the console where Git projects were found.

## Example

Console output:

```bash
$ ./gitfinder -dir "$HOME"
Start searching Git projects in /home/myuser ...
2021-02-13 11:07:32.5876333 +0100 CET m=+5.003815601 Found 1 Git Projects
2021-02-13 11:07:37.5873489 +0100 CET m=+10.003531201 Found 3 Git Projects
1 [/home/myuser/AndroidStudioProjects/MusicPlayer https://github.com/myuser/music-player]
2 [/home/myuser/GoProjects/MySimpleApp https://github.com/myuser/my-simple-app]
3 [/home/myuser/NodeProjects/SimpleRestAPI https://github.com/myuser/simple-rest-api]
```

CSV output:

```
$ $ ./gitfinder -dir "$HOME" -csv results.csv
```

```csv
Directory,Origin
/home/myuser/AndroidStudioProjects/MusicPlayer,https://github.com/myuser/music-player
/home/myuser/GoProjects/MySimpleApp,https://github.com/myuser/my-simple-app
/home/myuser/NodeProjects/SimpleRestAPI,https://github.com/myuser/simple-rest-api
```
