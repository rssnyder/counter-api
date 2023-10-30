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

endpoint: `https://counter.k8s.rileysnyder.dev/<user>/<counter>`

create counter: `POST`
- `curl -X POST https://counter.k8s.rileysnyder.dev/rssnyder/heartbeat`

increment counter: `HEAD`
- `curl -I https://counter.k8s.rileysnyder.dev/rssnyder/heartbeat`

decrement counter: `DELETE`
- `curl -X DELETE https://counter.k8s.rileysnyder.dev/rssnyder/heartbeat`

get counter: `GET`
- `curl https://counter.k8s.rileysnyder.dev/rssnyder/heartbeat`
