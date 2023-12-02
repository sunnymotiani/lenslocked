# Creating SQL tabels

```sql
    CREATE TABLE table_name(
        field_name TYPE CONTRAINTS,
        field_name TYPE(args) CONTRAINTS
    );
```
EXAMPLE SQL:

```sql
    CREATE TABLE users(
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE ,

    )

```

