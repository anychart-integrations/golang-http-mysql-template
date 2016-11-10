[<img src="https://cdn.anychart.com/images/logo-transparent-segoe.png?2" width="234px" alt="AnyChart - Robust JavaScript/HTML5 Chart library for any project">](https://anychart.com)
Go basic template
=========================

This example shows how to use Anychart library with the Go programming language and MySQL database.

## Running

To use this sample you must have Go installed as described [here](https://golang.org/doc/install), and follow the [official Go code organization](https://golang.org/doc/code.html);
MySQL installed and running, if not please check out [MySQL download page](https://dev.mysql.com/downloads/installer/) and follow [these instructions](http://dev.mysql.com/doc/refman/5.7/en/installing.html).

To check your installations, run the following command in the command line:
```
$ go version
go version go1.7.1 linux/amd64 # sample output
$ mysql --version
mysql  Ver 14.14 Distrib 5.5.52, for debian-linux-gnu (x86_64) using readline 6. # sample output
```

To start this example run commands listed below.

Clone the repository from github.com:
```
$ go get github.com/anychart-integrations/golang-http-mysql-template
```

Navigate to the repository folder:
```
cd $GOPATH/src/github.com/anychart-integrations/golang-http-mysql-template
```

Install dependencies:
```
$ go get ./...
```

Compile demo:
```
$ go install github.com/anychart-integrations/golang-http-mysql-template
```

Set up MySQL database, use -u -p flags to provide username and password:
```
$  mysql < database_backup.sql
```

Run example:
```
$ golang-http-mysql-template
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
                golang-http-mysql-template/
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
Language - [Go](https://golang.org/)<br />
Database - [MySQL](https://www.mysql.com/)<br />
Web framework - [net/http](https://golang.org/pkg/net/http), [html/template](https://golang.org/pkg/html/template/)<br />

## Further Learning
* [Documentation](https://docs.anychart.com)
* [JavaScript API Reference](https://api.anychart.com)
* [Code Playground](https://playground.anychart.com)
* [Technical Support](https://anychart.com/support)

## License
[© AnyChart.com - JavaScript charts](http://www.anychart.com). Released under the [Apache 2.0 License](https://github.com/anychart-integrations/golang-http-mysql-template/blob/master/LICENSE).
