# Basic Load Balancer with Go

Basic Load Balancer is for coding exercise on Go. That's way of to learn proxy with a very simple example. 

This repo has created from Ahmet Alp Balkan's [YouTube video](https://www.youtube.com/watch?v=QTBZxDgRZM0). You should watch it.

## Usage

```bash
go run .
or
go run load-balancer.go
```

Then start three servers that have different ports

```bash
npx http-server -p 5001
npx http-server -p 5002
npx http-server -p 5003
```

Now, you can send a request to listening address in ```load-balancer.go```
