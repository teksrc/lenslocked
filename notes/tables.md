# PSQL Tables

```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NO NULL
);
```

```sql
DROP TABLE IF EXISTS users;
```

```sql
SELECT id, email FROM users;
```

```sql
SELECT * FROM users WHERE email='frankacarv@gmail.com';
```

```sql
SELECT * FROM users WHERE age > 30;
```

```sql
SELECT * FROM users
WHERE age < 40 OR last_name = 'Carv';
```

```sql
UPDATE users
SET first_name = 'Frank', last_name = 'Carvajal'
WHERE id = 1;
```

```sql
UPDATE users
SET first_name = 'Frank';
```

```sql
UPDATE users
SET first_name = 'Anonymous', last_name = 'Man'
WHERE age < 120 AND age > 21;
```

```sql
DELETE FROM users
WHERE id = 1;
```

```
https://selectstarsql.com/ - place to start.
https://mystery.knightlab.com/ - SQL Murder Mystery solve by querying data.
https://www.codecademy.com/learn/learn-sql
http://www.w3schools.com/sql/
https://www.khanacademy.org/computing/computer-programming/sql
https://www.quora.com/How-do-I-learn-SQL - Quora question titled “How do I learn SQL?” that has 100+ answers and a large list of potential resources

PopSQL - Clean UI, native desktop app, team collaboration, and more, but it is a bit pricey. There is a free tier for a single user.

SQLECTRON - Nice free desktop app that is open source. A great alternative if you want something non-commercial with a GUI.
```
