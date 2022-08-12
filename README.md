# Go MongoDB REST

...

To know more details, there is the Makefile file with the commands necessary for its execution and example of how to use, including curl commands.

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

```
make test-show
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
