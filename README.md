## The task

**Requirements**:
- You need to implement a solution to record the **access logs** for the service.
- All **incoming requests** should be written into a table in the database, as well as **outgoing responses**.
- The tables and their structure must be added as migrations in the `./migrations` directory (see [Database migrations](#database-migrations)).
  - It's up to you to decide how you want to design the database tables, their columns, and types.
- Add a markdown file called `./task.md` explaining your thoughts and paradigms used to implement the solution.

**Things to consider**:
- Keep in mind this code might be extended in the future by other developers.
  - For example, multiple endpoints could be added, and those should also record the access logs.
- Find a way to somehow ensure your solution is **working as expected** following the above listed requirements.
- Consider this is a user-facing API, so we need to reduce the latency impact on endpoints.

**Other points**:
- The task should take **no longer than `4 hours`** to complete.
- You have a week to complete the task. If you want a time extension, or you got stuck solving it, do not hesitate to email us.

### Submitting the task
- The changes should be committed under a **new branch** called `task`.
- Create a Pull Request with the title `Task`.
- Send us a word once you’ve completed the task, and we will contact you after reviewing it.

## Development

### Requirements
- [Docker](https://docs.docker.com/compose/install/)
- [pressly/goose](https://github.com/pressly/goose#install)

### Run
To run the application, simply type `docker compose up` in the root folder.

Migrations will be applied and the HTTP server will bind to port `8080` together with a Postgres database which can be accessed on port `5432`.

To check if everything is going fine:
```bash
curl -v http://localhost:8080/health
```
You should get as response:
```
ok!
```

### Database migrations
To manage the database migrations, we are using [pressly/goose](https://github.com/pressly/goose).

You do not need to install the tool locally as we provide the option to run it via docker via the provided makefile commands:

- Check migration status: `make migration_status`
- Create a new migration: `make migration name=my_migration_name_here`
- Execute migrations: `make migrate`
- Rollback the last migration: `make migrate_down`

If you do want to run it locally the relevant credentials and settings can be found in `.env`
