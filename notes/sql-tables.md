# SQL Tables & PostgreSQL Data Types

```sql
CREATE TABLE table_name(
  field_name TYPE CONSTRAINTS,
  field_name TYPE(args) CONSTRAINTS,
);
```

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE
);
```

## Types & Description

https://www.postgresql.org/docs/current/datatype.html

PostgreSQL supports a wide variety of data types, for this project I will only need to use a handful of them. Below are a few common data types.

Types provide a way to define what kind of data you want to store in a particular column. For example, you might say “I only want to store integers in this column” or “I only want to store strings that are smaller than 140 characters in this column”.

`int` This is used to store integers between -2147483648 and 2147483647.

`serial` This is used to store integers between 1 and 2147483647. The big difference between int and serial is that serial will automatically set a value if you don’t provide one, and the new value will always increase by 1. This is useful for things like the id column, where you want every row to have a unique value and are okay with the database deciding what value to use.

`varchar` This is like a string in Go or other programming languages, except we have to tell the database what the max length of any string we are storing is going to be.

`text` This is a type that is specific to PostgreSQL (and may not be available in all forms of SQL) but it is basically the same as varchar under the hood, but you don’t have to specify a maximum string length when you declare your field.

In Postgres I prefer text over varchar in most cases, but it is important to know that both exist. Not all SQL databases have the same types, and text is one that is Postgres specific.

Used to define my database tables.

```sql
CREATE TABLE users (
  id SERIAL,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT
);
```

## Constraints

Constraints are rules that I can apply to columns in my table. For example, I might want to ensure that every user in my database has a unique `id`, so I could use the `UNIQUE` constraint. I could also verify that an integer column is above or below a specific value.

Like data types, there are many constraints available in Postgres: https://www.postgresql.org/docs/current/ddl-constraints.html

For now I will only be using a few:

`UNIQUE` This ensures that every record in your database has a unique value for the field that is set to unique. For example, you might want every user to have a unique email address. It is important to note that Postgres is case sensitive, so frankCARV@gmail.com is not the same as frankcarv@gmail.com. You will need to account for this on your own when writing data to your database.

`NOT NULL` This ensure that every record in your database has a value for this field. When you don’t provide a value for a field the database will traditionally store null, but this prevents that from being valid.

`PRIMARY KEY` This constraint is similar to combining both UNIQUE and NOT NULL but it can only be used once on each table, and it will automatically result in a the creation of an index for this field. The index is used to make it faster to look up records by this field. For example, I typically set the id to be the primary key, and then I look up users by their id when I need to fetch them from the database.

While it is common to validate data inside our own Go code, many constraints need to also be set at the database layer to ensure they are always enforced. For instance, if two users submit a form at the same time it is possible for your code to process both forms at roughly the same time and, depending on the timing, each might check and verify that a field is unique before either has inserted data into the database. As a result, it is possible to end up with a duplicate on a field that should be unique if validations only occur in your application code. On the other hand, if I set up constraints on my database one of the two submissions would fail when my Go code attempted to create a record with duplicate data.

In short, it is important to use database level constraints.
