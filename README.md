# Code Climate Govet Engine

`codeclimate-govet` is a Code Climate engine that wraps [govet](https://godoc.org/golang.org/x/tools/cmd/vet). You can run it on your command line using the Code Climate CLI, or on our hosted analysis platform.

govet examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers.

### Installation

1. If you haven't already, [install the Code Climate CLI](https://github.com/codeclimate/codeclimate).
2. Run `codeclimate engines:enable govet`. This command both installs the engine and enables it in your `.codeclimate.yml` file.
3. You're ready to analyze! Browse into your project's folder and run `codeclimate analyze`.

### Building

```console
./bin/build
```

This will build a `codeclimate/codeclimate-govet` image locally

### Need help?

For help with Govet, [check out their documentation](https://golang.org/cmd/vet/).

If you're running into a Code Climate issue, first look over this project's [GitHub Issues](https://github.com/codeclimate/codeclimate-govet/issues), as your question may have already been covered. If not, [go ahead and open a support ticket with us](https://codeclimate.com/help).
