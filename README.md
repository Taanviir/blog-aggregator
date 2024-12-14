# Blog Aggregator aka gator

Gator is a Go-based CLI application for aggregating, browsing, and following RSS feeds. It is designed to simplify the consumption of blog and news content for users. This project uses PostgreSQL for database queries and supports user management and feed subscriptions.

## Features

- **User Management**: Create and manage user accounts.
- **RSS Feed Aggregation**: Fetch and store RSS feed content.
- **Feed Browsing**: View available feeds and posts.
- **Follow Feeds**: Users can follow feeds and view personalized content.
- **Efficient Database Interaction**: SQL queries are managed with `sqlc` for type-safe database operations.

---

## Prerequisites

- **Go**: Version 1.18 or higher.
- **PostgreSQL**: A running PostgreSQL instance for the database.
- **sqlc**: To regenerate database code from SQL queries.

---

## Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/Taanviir/blog-aggregator.git
    cd blog-aggregator
    ```

2. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set Up Database**:
    - Create a PostgreSQL database.
    - Apply schema migrations from `sql/schema/`.
      ```bash
      psql -U <username> -d <database_name> -f sql/schema/001_users.sql
      psql -U <username> -d <database_name> -f sql/schema/002_feeds.sql
      # Repeat for remaining files in schema
      ```

4. **Set Up Configuration File**:
    - Create a configuration file in your home directory named ~/.gatorconfig.json.

    ```bash
    {
        "db_url": "url",
        "current_user_name": "user"
    }
    ```
    - Replace db_url with your PostgreSQL connection string and current_user_name with your desired username. The application will automatically look for this configuration file at ~/.gatorconfig.json during runtime.

5. **Install the CLI**:
    ```bash
    go install github.com/Taanviir/blog-aggregator@latest
    ```
    - This will compile and install the gator CLI tool, making it available in your $GOPATH/bin directory.
    - Ensure this directory is in your PATH to use the gator command globally.

---

## Usage

### CLI Commands

To use the application, run the gator executable followed by a command:

```bash
gator <command> [<args>]
```

#### Available Commands

- `login`: Log in as an existing user.

- `register`: Register a new user account.

- `reset`: Reset user password.

- `users`: View all users (admin only).

- `agg`: Aggregate RSS feed data.

- `addfeed`: Add a new RSS feed (requires login).

- `feeds`: List all available RSS feeds.

- `follow`: Follow a specific RSS feed (requires login).

- `following`: View feeds followed by the logged-in user.

- `unfollow`: Unfollow a specific RSS feed (requires login).

- `browse`: Browse posts from followed feeds (requires login).

#### Example Usage

```golang
# Register a new user
gator register <username> <password>

# Log in
gator login <username> <password>

# Add a new feed (requires login)
gator addfeed <feed_url>

# Follow a feed (requires login)
gator follow <feed_id>

# View posts from followed feeds
gator browse
```

---

## Development

### Regenerate SQL Code

If you modify SQL queries, use [`sqlc`](https://sqlc.dev/) to regenerate code:
```bash
sqlc generate
```

---

## Contributing

1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
