# counter-api

## init

```sql
CREATE TABLE counters (
  id SERIAL PRIMARY KEY,
  uuser VARCHAR(64) NOT NULL,
  counter VARCHAR(64) NOT NULL,
  value INT NOT NULL,
  UNIQUE (uuser, counter)
);
```

## usage

endpoint: `https://api.counter.k8s.rileysnyder.dev/<user>/<counter>`

create counter: `POST`
- `curl -X POST localhost:8080/rssnyder/heartbeat`
increment counter: `HEAD`
- `curl -I localhost:8080/rssnyder/heartbeat`
decrement counter: `DELETE`
- `curl -X DELETE localhost:8080/rssnyder/heartbeat`
get counter: `GET`
- `curl -X POST localhost:8080/rssnyder/heartbeat`
