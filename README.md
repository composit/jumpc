# JumpC exercise

to run the server binary (for example on port 3000):
```
go build
./jumpc 3000
```
to request a password hash from the server:
```
curl --data "password=hunter2" localhost:3000
```
to shut down the server:
```
curl --data "graceful shutdown" localhost:3000
```
### Happy Hashing!
