# Golang Webserver Template
This is a template project that I commonly use to start various Go projects.

## Features
- Routing and middleware with [`chi`](https://github.com/go-chi/chi)
  - Asset routes
  - Public routes (e.g. `/` index page)
  - Private routes (e.g. `/dashboard` with basic authentication)
- Templates
  - `static/html/pages`
  - `static/html/components`
  - Reloaded every 3 seconds in debug mode, for quick live development
- Database using [`go-sqlite3`](https://github.com/mattn/go-sqlite3) (with [WAL optimization](https://www.sqlite.org/wal.html) enabled)
- HTTPS (TLS/SSL) using [`certmagic`](https://github.com/caddyserver/certmagic) (no reverse proxy required)

Running the application with `-production` will make templates reload less frequently and enable SSL for a specific domain. 

## Structure
- `database/` golang database interaction, as well as different structures (e.g. `User` struct, which may match a table)
- `handlers/` http routing handlers
- `static/`
    - `html/`
        - `pages/` contains all html pages like `index.html`
        - `components/` contains html components like `footer.html`
    - `assets/`
        - `css/` stylesheets
        - `js/` scripts
        - `...` and whatever else you want

## Build
Clone or download this repository.  
- `go run .` to run normally
- `go run . -production` to run in production mode

## FAQ
- **Why is the `sql.DB` handle not protected by a mutex?**
    - > It's safe for concurrent use by multiple goroutines. [[docs]](https://pkg.go.dev/database/sql#DB)
- **Why are assets and public routes in different router groups?**
    - There are several reasons for this. Let's say you want something like age verification on your site. In this case it'd probably be okay to request stylesheets and scripts, but not the different HTML pages before the age has been verified. Also, if you want to rate limit requests, it can be hard to do this when all is grouped together. You may want to rate limit page requests to something like `10 requests/min`, but assuming you have multiple assets (stylesheets, scripts, fonts, images, etc.) you might want to rate limit this differently, say `60 requests/min`.
- **Is basic authentication safe? Why not use JWT instead?**
    - SSL basic authentication is more than good enough for most purposes. Obviously the username and password shouldn't be easy to guess, and if wanted you could add several security measures like limiting language requests, user-agents, make your own 2FA like algorithm where the password auto-updates every minute, hour, day, etc... Basic authentication also works especially well for something like admin pages, as all modern browsers will allow you to access the grouped admin routes, without having to login again, until the session / browser is closed. JWT while modern and often recommended, can also be over-complicated in certain cases. However, if you are interested in JWT, `chi` has middleware for this called [`jwtauth`](https://github.com/go-chi/jwtauth).