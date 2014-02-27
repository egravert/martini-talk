martini-talk
============

To run the talk, install the present tool with the following command:
` code.google.com/p/go.talks/pkg/present`

Make sure `$GOPATH/bin` is in your `$PATH` and run `present` from the base directory of this repo.



==To run the example application==

* Create a postgresql database and execute the sql in the `secondbrain/create.sql` file.
* cd to the secondbrain directory
* build the project with 'go build'
* Run the application, replacing the psql and port env variables as necessary
```PORT=8020 PSQL_URI=postgres://localhost/secondbrain?sslmode=disable ./secondbrain```

