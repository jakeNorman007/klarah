![logo](./public/logo.png)

### A scaffolding tool for spinning up a quick Golang backend/API for your side projects.

Klarah is a great option to get an http server up and running quickly with little overhead. Let it take
care of the initial load of writing your own server from scratch. With many options of frameworks and
databases to choose from, you can customize it to fit your use case.

As stated in the prior paragraph, customization is a plus when using Klarah to start your project's backend.
Start with the boilerplate that is generated, add more or take away from what's there. The key is that you can
do as much or as little as you want once your project has been generated.

To install Klarah:
```
Install url goes here
```

When installed Klarah creates a binary that attaches directly to your GOPATH, so you can use it whenever you 
need to reach for it.

## Commands

Klarah has a make file that conveniently keeps all terminal commands organized and easy to use.

### General commands:

When you want to start the server after configuration, simply run:
```
make run
```

When needing to run the preset tests, or added tests in ./tests, run:
```
make test
```
## Database related commands:

Klarah uses goose under the hood to run all your database migrations, after configuring your database connection,
to migrate your sql up to your databse run:
```
make up
```

If you need to reverse those migrations run:
```
make down
```

## Klarah's supported frameworks:
- [Standard Library](https://pkg.go.dev/net/http#hdr-Servers)
- [Echo](https://github.com/labstack/echo)
- [Chi](https://github.com/go-chi/chi)
- [Gin](https://github.com/gin-gonic/gin)

## Klarah's supported databases:
- [Postgresql](https://github.com/jackc/pgx)
- [Sqlite](https://github.com/mattn/go-sqlite3)

## Getting started:
> [!NOTE]
> As it is now, Klarah requires a database connection in order to get the server up and running.

Download klarah to your GOPATH using:
```
Install path goes here
```

Once downloaded and ready run the command:
```
klarah new
```

You will then be prompted to create a project, once it's named press enter:
```
project name > your_project_name
```

Next you will see these prompts and asked to pick a framework, followed by a database.

> [!NOTE]
> Use down arrow or j to move cursor down, up arrow or k to move it up.

```
---------------------------------- Frameworks ----------------------------------

->  standard-library: Standard go library for creating http servers

    echo: High performing, extensible framework with minimal overhead

    chi: Lightweight, fast http router that's flexible and powerfull

    gin: Flexible, router for efficiant and scalable applications

Press enter to select and y to confirm your choice, Ctrl-c to exit.

------------------------------- Database drivers -------------------------------

    postgresql: Powerful, open-source relationa database management system

    sqlite: Lightweight, self-contained SQL database engine

Press enter to select and y to confirm your choice, Ctrl-c to exit.
```

After your project generates go into the directory with ```cd your_project_name```

To get the server up and running connect to a database and set up your goose variables in .env file

Postgresql set up example:

```
PORT=8080

DB_URL=postgres://example:example@example.com/example?sslmode=disable

GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://example:example@example.com/example?sslmode=disable
```

Run the command ```make run```

You will get the following output in the console

```
20XX/XX/XX 00:00:00 Connected to database
20XX/XX/XX 00:00:00 Server listening on port :8080
```

Congrats! You are ready to start building!

## Future features and additions:
- Air support for live reloading, https://github.com/air-verse/air
- Option to spin up a databaseless project to make it fit more use cases, to also be more customizable
