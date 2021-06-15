# todo - Markdown Todo interface

`todo` is a Go based cli tool and REST api that will scan a directory for Markdown files that contain Markdown format TODO items and serve them up to read and manipulate them via cli and eventually a full CRUD REST interface.

![test workflow](https://github.com/jujhars13/todo-markdown-todo-interface/actions/workflows/test.yml/badge.svg)

## Usage

```bash
# will scan this directory for Markdown files
# if it finds todo's it will print them out
todo -d ~/Dropbox/notes
```

## TODO

### Core

- [x] setup project
- [x] read active todos and spit them out to cli
- [ ] setup makefile
- [x] setup github actions for ci unit tests
- [x] publish test badges in readme
- [x] setup cli parsing
- [ ] mark a todo as done via cli
- [ ] investiage ncurses interface
- [ ] Create build and release pipeline
- [ ] Create docker container
- [ ] restructure code like [https://github.com/pokemium/Worldwide]()

### Nice to have

- [ ] create an http interface
- [ ] create GET (active) collection
- [ ] create PUT/PATCH item to mark todo as done
- [ ] create DELETE item
- [ ] create POST to collection
- [ ] Create docker container

## Decision Records

- Following the [project layout](https://github.com/golang-standards/project-layout) Go project layout standard
- Focus on CLI functionality first, then work on REST interface+


## Licence

[MIT](LICENCE)
