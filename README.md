# Summary

Can simply run locally: `cd main && go run main.go`.
And `curl -k 'localhost:5050/about'`.

Run with docker:
```bash
docker build -t gosimpleserver .
docker run -p 5050:5050 -tid gosimpleserver
```