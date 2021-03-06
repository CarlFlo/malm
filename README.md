# Malm

![Tests](https://github.com/CarlFlo/malm/actions/workflows/go.yml/badge.svg)

Malm is a simple to use logger for the [GO](https://golang.org/) programming language.

Test coverage: **93.6%**

## Features
- Detailed. Comprehensive logging showing the message origin file, calling function and line.
- Customizable. Easy to customize and toggle specific messages types on or off.
- Simplistic. No initial configuration required to get started.


## Install

```
go get github.com/CarlFlo/malm
```

The main functionality is working, but I'm always working to improve and fix bugs within the package.

## Usage

The goal of Malm is to make powerful logging simple, and you are able to get started with no configuration required.

By default, Malm will output **all types** of log messages to **os.Stderr**.


Example syntax:
```go
// Arguments are handled in the manner of fmt.Printf, with '\n' automatically added to the end.
malm.Fatal("Fatal log message: '%s'", err) // Will run os.Exit(1)
malm.Error("Error log message: '%s'", err)
malm.Warn("Warning log message")
malm.Info("Info log message")
malm.Debug("Debug log message")

// Custom messages
malm.Custom(os.Stderr, "CUSTOM", "This is a %s message with a custom log tag", "custom")
malm.Custom(os.Stderr, "NETWORK", "Another example with a different log tag")
```

Each of the logging functions above will return **True** if the message could be logged and **False** if it was blocked by a setting. Each of the functions, except **Custom**, can be treated like **log.Printf** and **fmt.printf**. [Click here](https://golang.org/pkg/fmt/) for documentation on formatting.


## Options & Customization

To customize the logging messages that get displayed, the following syntax be used to change the settings.

```go
// To turn on indivudual logging (Default state)
malm.SetLogFatal(true)
malm.SetLogError(true)
malm.SetLogWarning(true)
malm.SetLogInfo(true)
malm.SetLogDebug(true)
malm.SetLogCustom(true)

// To turn off indivudual logging
malm.SetLogFatal(false)
malm.SetLogError(false)
malm.SetLogWarning(false)
malm.SetLogInfo(false)
malm.SetLogDebug(false)
malm.SetLogCustom(false)
```

To customize the logging messages verbosity, the following syntax be used to change the settings.

```go
// To turn on verbose logging (Default state)
malm.SetLogVerboseFatal(true)
malm.SetLogVerboseError(true)
malm.SetLogVerboseWarning(true)
malm.SetLogVerboseInfo(true)
malm.SetLogVerboseDebug(true)
malm.SetLogVerboseCustom(true)

// To turn off verbose logging
malm.SetLogVerboseFatal(false)
malm.SetLogVerboseError(false)
malm.SetLogVerboseWarning(false)
malm.SetLogVerboseInfo(false)
malm.SetLogVerboseDebug(false)
malm.SetLogVerboseCustom(false)
```

```go
// A verbose message looks like this:
<date and time> [<log tag>] <filePath>:<line number>:<caller>() <formatted message>\n

// A non-verbose message looks like this
<date and time> [<log tag>] <formatted message>\n
```
Each of the above functions will return a copy of the updated bitmask (uint8)

---

A bitmask is used to calculate what get outputted:
* fatal = 1
* error = 2
* warning = 4
* info = 8
* debug = 16
* custom = 32

This allows the user to input a prepared value from, e.g. a configuration file to set the desired logging.
```go
malm.SetLogBitmask(63) // Will turn on everything 1+2+4+8+16+32=63
malm.SetLogBitmask(59) // Will turn on everything except warnings (4) 1+2+8+16+32=59
malm.SetLogBitmask(0) // Turns off all logging
```

This allows the user to input a prepared value from, e.g. a configuration file to set the desired logging.
```go
malm.SetLogVerboseBitmask(63) // Will turn on verbosity for everything 1+2+4+8+16+32=63
malm.SetLogVerboseBitmask(59) // Will turn on verbosity for everything except warnings (4) 1+2+8+16+32=59
malm.SetLogVerboseBitmask(0) // Turns off verbosity for all logging messages
```

---

The default output is **os.Stderr**, but this can be changed with:
```go
malm.SetDefaultWriter(newWriter) // Will accept any io.Writer
```

---

It is also possible to change the time format of the logging message.

The default time format is **2006-01-02 15:04:05** but can be changed to any supported string.

```go
malm.SetTimeFormat("2006-01-02 15:04:05")
```
[Click here](https://golang.org/pkg/time/) for documentation. 

---

## Roadmap
- [X] Basic functionality
- [X] Ability to customize which log types gets logged
- [X] Additional message logging types (such as Custom and Fatal)
- [X] Test coverage above at least 80%
- [X] Additional error checking
- [X] Ability to change the time format
- [X] Option for 'verbosity' for each type of log message
