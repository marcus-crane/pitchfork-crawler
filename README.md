# pitchfork-crawler

Nothing fancy, a small application written in Go that takes an artist name as a flag (`-artist`) and spits out album scores for that artist in the form of JSON

## Example usage

```go
go run main.go -artist mf doom
[
 {
  "album": "Czarface Meets Metal Face",
  "artist": "CzarfaceMF DOOM",
  "score": "6.4"
 },
 ...
]
```
