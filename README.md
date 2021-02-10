# crypto-index

Folder Structure Conventions
============================

> Folder structure options and naming conventions for software projects

### A top-level directory

    .
    ├── dist                   # Compiled files (alternatively `build`)
    ├── api                    # Backend Api source files (alternatively `backend`)
    ├── frontend               # Source files (alternatively `web` or `app`)
    ├── main.go                # Golang init program
    ├── LICENSE
    └── README.md


### Api Golang

The api directory contains the Golang source code.

> This directory could be called a backend depending on the naming convention. Here is the entire Golang code.

    api 
    ├── db                      # Database files
    ├── entity                  # Entities files (alternatively `schema`)
    ├── handlers                # Handlers process requests  (alternatively `controllers` or `delivery`)
    ├── helper                  # Helpers functions files (alternatively `util`)
    ├── repo                    # Repo, manipulate data in the database (alternatively `DAO`)
    ├── storage                 # Storage and location of specific memorized (alternatively `assets`)

## Elm

