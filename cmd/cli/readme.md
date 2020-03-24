# Au Suivant CLI

The Command Line version of the app, currently only for Mac.

## Usage

```
suivant [-delay <seconds>]
```

`suivant` reads a list of new line separated text from stdin
and then proceeds to literally call them with a gap of however many
seconds you set.

## Options

- `-delay` the delay in seconds between calling next on the list.
Defaults to 10. 