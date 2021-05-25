# todo - Markdown Todo interface

`todo` is a Go based cli tool and REST api that will scan in a directory for Markdown files that contain Markdown format TODO items and serve them up to read and manipulate them via cli and eventually a full CRUD REST interface.

![test workflow](https://github.com/jujhars13/todo-markdown-todo-interface/actions/workflows/test.yml/badge.svg)

## TODO

### Core

- [x] setup project
- [x] read active todos and spit them out to cli
- [ ] setup makefile
- [x] setup github actions for ci unit tests
- [x] publish test badges in readme
- [ ] mark a todo as done via cli
- [ ] Create build and release pipeline
- [ ] Create docker container

### Nice to have

- [ ] create an http interface
- [ ] create GET (active) collection
- [ ] create PUT/PATCH item to mark todo as done
- [ ] create DELETE item
- [ ] create POST to collection
- [ ] Create docker container

## Decision Records

- Following the [project layout](https://github.com/golang-standards/project-layout) go standard
- Focus on CLI functionality first, then work on REST interface+


## Licence

[MIT](LICENCE)
