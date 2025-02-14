## Logger Service

# Project specific folders

- cmd
  - contains the entrypoint of the application, which is the `main.go` file.
  - it's also contain CLI for running consumer.

- internal
  - contains all the business logic inspired by DDD (Domain Driven Design)
  - in this folder and sub-folders, you will only find domain models, domain services.
- migrations
  - contains all db migrations.

# How I Solved the Problem

To address the requirement of logging requests and responses to the database, I implemented a robust solution using middleware and a queue system.

First, I added middleware to the project that intercepts each incoming request. This middleware captures all necessary details of the request, including headers, payloads, and metadata. Instead of directly writing this data to the database—which could introduce latency—I designed the middleware to enqueue the captured data into a message queue system.

On the other side, a consumer service processes the queued data asynchronously. It retrieves the request and response details from the queue and stores them in the database.

This architecture ensures fault tolerance and reliability. Even if the consumer service goes down temporarily, the queue retains all the captured requests. Once the consumer service is back online, it processes the pending data without any loss. By decoupling the logging logic from the request-response cycle, this approach minimizes latency while ensuring every request is reliably logged.

This solution not only meets the immediate requirement but also provides scalability and resilience, accommodating high traffic and potential system outages gracefully.

# Main Packages
- NATS as the queue system.
- To make the logging system reusable across the project, I encapsulated Zap within a custom wrapper.

# How to Test

1- Start the project by running `docker-compose up -d`.

2- Send test requests using the following command:

`for i in {0..10000}; do curl -X GET http://localhost:8080/health; done` or use make file `make data`

3- Verify the data in the database to ensure the requests have been logged correctly.

# Testing the Consumer Failure Scenario
1- Stop the consumer container using `docker-compose down consumer`.

2- Send multiple requests to the service.

3- Check the database to confirm no new data has been logged while the consumer is down.

4- Restart the consumer container with `docker-compose up consumer -d`.

5- Recheck the database to confirm all previously queued data has now been transferred to the database successfully.

# Testing
For test I used mock and testify package.

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
