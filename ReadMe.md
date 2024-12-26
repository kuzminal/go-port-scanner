# Консольное приложение для сканирования TCP портов
```go
    Usage of port scanner:
        -addr string
            Address to scan (default "scanme.nmap.org")
        -end int
            End with port number (default 1024)
        -start int
            Start from port number greater or equals than 1 (default 1)
        -t duration
            Timeout in seconds (default 5s)
```
Для запуска сканирования можно выполнить команду:
```go
go run main.go -start 1 -end 100 -addr scanme.nmap.org -t 5s
```
    