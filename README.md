# 📰 Gator — Go Blog Aggregator

A CLI-based RSS feed aggregator built in Go with a PostgreSQL backend. Started as a guided [Boot.dev](https://boot.dev) project, Gator lets you register users, add RSS feeds, follow feeds others have added, and run a background scraper that continuously pulls the latest posts into your local database for browsing.

This is an actively evolving project — see the [Roadmap](#roadmap) below for planned features.

---

## Prerequisites

Before you can run Gator, make sure you have the following installed:

- **Go** (1.22+) — https://go.dev/doc/install
- **PostgreSQL** (v15+) — https://www.postgresql.org/download/
- **Goose** (database migration tool) — install with:
```bash
  go install github.com/pressly/goose/v3/cmd/goose@latest
```

---

## Installation

You can install the `gator` CLI directly using `go install`:

```bash
go install github.com/AggroSec/Go-Blog-Aggregator@latest
```

Make sure your Go binary path is in your `$PATH`. If it isn't, add this to your `~/.bashrc` or `~/.zshrc`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## Database Setup

Even when installing via `go install`, you'll need a properly configured PostgreSQL database before running Gator.

### 1. Create the database in PostgreSQL

Connect to your PostgreSQL instance:

```bash
# Linux / WSL
sudo -u postgres psql

# macOS (Homebrew)
psql postgres
```

Then create the database:

```sql
CREATE DATABASE gator;
```

You can verify it was created with `\l`, then exit with `\q`.

### 2. Run the migrations with Goose

Gator uses Goose to manage the database schema. The migration files live in `sql/schema/`. Clone the repo to get the schema files:

```bash
git clone https://github.com/AggroSec/Go-Blog-Aggregator.git
cd Go-Blog-Aggregator
```

Then run the migrations against your database:

```bash
goose postgres "postgres://postgres:[password]@[host]:[port]/gator" -dir sql/schema up
```

Replace `[password]`, `[host]`, and `[port]` with your actual PostgreSQL credentials. For a default local install this typically looks like:

```bash
goose postgres "postgres://postgres:yourpassword@localhost:5432/gator" -dir sql/schema up
```

Connect to the `gator` database and run `\dt` to confirm the tables were created successfully.

---

## Configuration

Gator reads its config from a JSON file in your home directory named `.gatorconfig.json`. You must create this file manually before running the program.

**Location:** `~/.gatorconfig.json`

**Contents:**

```json
{
  "db_url": "postgres://postgres:[password]@[url_of_database]:[port]/gator?sslmode=disable"
}
```

Replace `[password]`, `[url_of_database]`, and `[port]` with your actual PostgreSQL connection details. For a default local installation it will typically look like:

```json
{
  "db_url": "postgres://postgres:yourpassword@localhost:5432/gator?sslmode=disable"
}
```

---

## Running Gator

Once installed, configured, and migrated, run any command with:

```bash
Go-Blog_Aggregator <command> [arguments]
```

---

## Commands

### User Management

| Command | Description |
|---|---|
| `Go-Blog_Aggregator register <name>` | Create a new user and set them as the current user |
| `Go-Blog_Aggregator login <name>` | Switch the current active user |
| `Go-Blog_Aggregator users` | List all registered users |
| `Go-Blog_Aggregator reset` | ⚠️ Delete all users from the database (destructive) |

### Feed Management

| Command | Description |
|---|---|
| `Go-Blog_Aggregator addfeed <name> <url>` | Add a new RSS feed and automatically follow it |
| `Go-Blog_Aggregator feeds` | List all feeds in the database |
| `Go-Blog_Aggregator follow <url>` | Follow an existing feed by URL |
| `Go-Blog_Aggregator following` | List all feeds the current user is following |
| `Go-Blog_Aggregator unfollow <url>` | Unfollow a feed |

### Aggregation & Browsing

| Command | Description |
|---|---|
| `Go-Blog_Aggregator agg <interval>` | Start the feed scraper. Fetches new posts at the given interval (e.g. `30s`, `1m`, `5m`). Runs continuously until stopped with `Ctrl+C` |
| `Go-Blog_Aggregator browse [limit]` | Display the latest posts from your followed feeds. Defaults to 2 posts; pass a number to show more (e.g. `gator browse 10`) |

### Example Workflow

```bash
# Register a user
Go-Blog_Aggregator register alice

# Add a couple of RSS feeds
Go-Blog_Aggregator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
Go-Blog_Aggregator addfeed "Hacker News" https://hnrss.org/frontpage

# Start scraping every 30 seconds (run in a separate terminal or in the background)
Go-Blog_Aggregator agg 30s

# Browse the latest 5 posts
Go-Blog_Aggregator browse 5
```

---

## Project File Structure

```text
Go-Blog-Aggregator/
├── main.go                  # Entry point — wires up state, registers commands, dispatches CLI args
├── commands.go              # Command registry and dispatch logic
├── state.go                 # App state struct (DB connection, config)
├── middleware.go            # Auth middleware — resolves current user before protected commands
├── cmdRegister.go           # register command
├── cmdLogin.go              # login command
├── cmdUsers.go              # users command
├── cmdReset.go              # reset command
├── cmdAddFeed.go            # addfeed command
├── cmdFeeds.go              # feeds command
├── cmdfollow.go             # follow / unfollow / following commands
├── cmdAgg.go                # agg command — scraper loop and RSS fetching logic
├── cmdBrowse.go             # browse command — displays latest posts for current user
├── sqlc.yaml                # sqlc configuration for code generation
├── go.mod
├── go.sum
├── sql/
│   ├── schema/              # Goose migration files (run these to set up the DB)
│   └── queries/             # SQL queries used by sqlc to generate Go code
└── internal/
├── database/            # Auto-generated sqlc database layer (models, queries)
└── RSS/                 # Internal RSS feed fetching and parsing package
```

---

## Roadmap

Gator started as a Boot.dev guided project but is planned to grow into a more fully-featured, remotely-usable tool. Here's what's on the horizon:

### 🔜 Near Future
- [ ] **Post pagination** — Page through posts rather than retrieving a flat list with a hard limit, making browsing large feeds much more practical
- [ ] **Terminal UI (TUI)** — A proper in-terminal interface (likely using a library like [Bubble Tea](https://github.com/charmbracelet/bubbletea)) to browse, read, and navigate posts without leaving the command line

### 🔭 Further Down the Road
- [ ] **API authentication & remote hosting** — Add proper auth (API keys or token-based) so Gator can be hosted on a server and accessed by multiple users remotely, rather than being limited to a local setup
- [ ] **Multi-user remote access** — Once auth is in place, expose an HTTP API so users can interact with the service from anywhere
- [ ] **And more...** — Feed categorization, search, bookmarking, and other quality-of-life improvements as the project evolves

---

## Tech Stack

- **Language:** Go
- **Database:** PostgreSQL
- **DB Migrations:** [Goose](https://github.com/pressly/goose)
- **SQL Code Generation:** [sqlc](https://sqlc.dev/)
- **RSS Parsing:** Custom internal RSS package

---

## Acknowledgements

Built following the [Boot.dev](https://boot.dev) guided project curriculum — a great hands-on platform for learning backend development with Go.

---

*Licensed under MIT*
