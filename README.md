# azure-savelogs

This tool runs the program in its argument and captures stdout and stderr.
They are stored to files in the directory

* `%HOME%\Logfiles\application` .

It is meant for webservices on Azure.

Azure-savelogs is a drop-in replacement of David Ebbo's ConsoleInterceptor
which did not work for me:

* https://github.com/davidebbo-test/ConsoleInterceptor

## Build/Install

Install Go and type in the shell

```shell
go get github.com/freiheit-com/azure-savelogs
```

## Usage

```xml
<httpPlatform processPath="azure-savelogs.exe webservice.exe">
```

Further arguments are passed to the webservice.

Enjoy.
