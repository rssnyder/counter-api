# counter-api

```sql
CREATE TABLE counters (
  id SERIAL PRIMARY KEY,
  uuser VARCHAR(64) NOT NULL,
  counter VARCHAR(64) NOT NULL,
  value INT NOT NULL,
  UNIQUE (uuser, counter)
);
```
