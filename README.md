# Revealit

`revealit` is a small binary that helps with the identification of dependencies and their categories.
When you start on a new project, it's always interesting to understand what people have been using.

## Compile

```bash
go build .
```

## Help

```bash
Usage of ./revealit:
  -language string
      the project's main language
```

## Run

```bash
./revealit --language=ruby
```

## Support
### Ruby
Thanks to [The Ruby Toolbox](https://www.ruby-toolbox.com/) for allowing the download of a [JSON](https://github.com/rubytoolbox/catalog) file containing all their data.


## Limitations
It currently supports only Ruby and Gemfile.

As Ruby is my primary language, I've just implemented the parse and identification for it. Feel free to add support for other languages.

