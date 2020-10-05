# godot-go-template

This is a template folder of github.com/godot-go/godot-go's project

## Prepare

1. Copy this repository inside your godot's project folder
2. Modify destination folder, filename in magefile.go
3. Rename and move libgodotgo-myscript.tres into your project's favorite place. And fix filenames inside this file as same as step 2.

## Folder structure

* (root): package folder that contains Go logic
* /entrypoint: it includes main package
* /export: it includes CGo related code

## How to edit Scripts

It contains one sample named `mycounter.go`. Create your own struct to extends Godot's API.

If you add your own struct, edit `/export/export.go`.

## How to build

It uses [mage](https://magefile.org/) as a task runner.

You can build this folder like the following:

```sh
$ go run mage.go
```

## License

MIT