# Go MongoDB REST

This project is an open-source REST client for MongoDB, designed to facilitate the execution of all CRUD operations and additional features with ease. Users can create, edit, replace, delete, and read documents within collections. The client also supports advanced queries, allows users to count documents within a collection (with or without a query), and provides functionality for paginated data retrieval.

The project is seamlessly executed via Docker Compose, with MongoDB pre-configured and the project set to connect to it effortlessly.

For more detailed information, including commands necessary for execution and usage examples (such as `curl` commands), refer to the Makefile included in the project.

# How to use

Run the application:

```
make run
```

Start the application:

```
make start
```

Stop the application:

```
make stop
```

View the logs from the application:

```
make logs
```

Example of how to list all documents from a collection:

```
make test-list
```

Example of how to list documents from a collection with a find:

```
make test-list-find
```

Example of how to get a document via `_id` from a collection:

(replace **<_ID>** with the _id from document.)

```
make test-show _id=<_ID>
```

Example of inserting documents into a collection:

```
make test-insert
```

Example of replacing all document data from a collection:

```
make test-replace
```

Example of replacing all document data from a collection with a find:

```
make test-replace-find
```

Example of how to update all documents from a collection:

```
make test-update
```

Example of how to update documents from a collection with a find:

```
make test-update-find
```

Example of how to delete all documents from a collection:

```
make test-delete
```

Example of how to delete documents from a collection with a find:

```
make test-delete-find
```

Example of how to perform a search using MongoDB query:

```
make test-query
```

Example of how to count the number of documents in a collection:

```
make test-count
```

Example of how to use pagination:

```
make test-paginate
```
