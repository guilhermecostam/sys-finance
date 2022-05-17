<h1 align="center">
<br>
sys-finance
</h1>

<p align="center">Simple API for transaction maintenance. Developed to study the GO language.</p>

<hr />

## Features

These are the features I used in developing this project.

- **Golang** — Open-source and compiled language.
- [**Goose**](https://github.com/pressly/goose) — Open-source lib for create migrations in GO.
- **Docker Compose** — Docker's container orchestrator.
- **Postgres** — Relational database.

## Getting started

```bash
# Clone this repository
$ git clone https://github.com/guilhermecostam/sys-finance.git

# Create your .env file and overwrite
$ cp .env.example .env

# Run docker compose
$ docker-compose up -d

# Go to migrations path and run the migration
$ cd database/migrations && goose up

# Now, go to server main file and run the application
$ cd .. && cd cmd && go run .
```

## Acces to the API
All requests to the API begin with:

```shell
localhost:8080/transactions
```

For receive all results of transactions, just send this endpoint with the GET method.
For create a transaction, just send this endpoint with the POST method and the required requests.

Now, for calls using a specific transaction, it is necessary to send an ID in the scope of the endpoint:
```shell
localhost:8080/transactions/{$transactionId}
```

For receive a especific transaction, just send this endpoint with the GET method.
For update a especific transaction, just send this endpoint with the PUT method.
For delete a especific transaction, just send this endpoint with the DELETE method.

## How to contribute
Do you want to contribute to the project? Just follow these instructions:

- Fork this repository;
- Clone your repository;
- Create a branch with your feature:
`
git checkout -b my-feature
`;
- Commit your changes:
`
git commit -m 'feat: My new feature'
`;
- Push to your branch:
`
git push origin my-feature
`;
- Come in Pull Requests from the original project and create a pull request with your commit;

After the merge of your pull request is done, you can delete your branch.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/guilhermecostam/sys-finance/blob/master/LICENSE) page for details.

---

Made with :zap: by Guilherme Costa. [Get in touch!](https://www.linkedin.com/in/guilhermecostam/)