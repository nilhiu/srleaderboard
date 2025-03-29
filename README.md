# SRLeaderboard

SRLeaderboard is a small (single-page) web application that keeps track of
completion times of speedruns. It's written with the following technologies:

- For the frontend: Templ, HTMX, Alpine.js and Tailwind CSS.
- For the backend: Go, PostgreSQL and Redis.

## Building

To build SRLeaderboard you will need to have installed `go`, `make`, `templ`,
`docker` and `node`. Create a `.env` file by copying the `.env.example` file:

```bash
cp .env.example .env
```

Afterward, edit the file and change the secrets, for security purposes do **NOT**
keep the defaults provided in the `.env.example`. Finally, run:

```bash
make
```

And let Docker download and setup all the images necessary.

## Customization

If you want to change the color scheme of the frontend, you can change the
Tailwind CSS theme configuration located at `internal/app/static/tailwind.css`,
and then run `make restart` to rebuild and restart SRLeaderboard

## API Routes

The API routes provided by SRLeaderboard, which are defined in
`internal/app/routes.go` (routes prefixed with `/api`), are the following:

- `GET /api/runs`
- `GET /api/runs/{user}`
- `POST /api/auth/login`
- `POST /api/auth/register`
- `POST /api/logout`
- `POST /api/runs`

You can view the documentation in [API.md](API.md) or in the
[OpenAPI spec](openapi.yaml).
