# AnyChart GO demo example

## Overview

This example shows how to use Anychart library with the Go programming language and MySQL database.

## Running

To use this sample you must have Go installed (as described at https://golang.org/doc/install), and follow the Go code organization https://golang.org/doc/code.html

To start this example run commands listed below.

Navigate to your Go working folder:
```
cd $GOPATH/src/github.com/{user}
```

Clone the repository from github.com:
```
$ git clone git@github.com:AnyChart/golang-template.git
```

Navigate to the repository folder:
```
$ cd golang-template
```

Install dependencies:
```
$ go get ./...
```

Compile demo:
```
$ go install github.com/{user}/golang-template
```

Set up MySQL database, use -u -p flags to provide your user name and password:
```
$  mysql < database_backup.sql
```

Run example:
```
$ golang-template
```

Open browser at http://localhost:8080/

## Workspace
Your workspace should look like:
```
$GOPATH/
    bin/
    pkg/
    src/
        github.com/
            {user}/
                golang-template/
                    public/
                        static/
                            css/
                                style.css   # css file
                    templates/
                        index.html          # html template
                    database_backup.sql     # MySQL database dump
                    main.go                 # main Go code
                    README.md

```

## Technologies
Language - Go<br />
Database - MySQL<br />
Web framework - net/http, html/template<br />

## Further Learning
* [Documentation](https://docs.anychart.com)
* [JavaScript API Reference](https://api.anychart.com)
* [Code Playground](https://playground.anychart.com)
* [Technical Support](https://anychart.com/support)
