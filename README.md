Streak API for Go
================================

An incomplete Go API wrapping [Streak's API](http://www.streak.com/api/)

Warning: this is totally incomplete. It only exposes enough for me to extract some statistics I wanted. Be prepared to extend the structs in [`streak.go`](streak.go). I'm very willing to accept pull requests!

Example Usage:
-----------------

See [`cmd/streak_example.go`](cmd/streak_example.go) for a complete example. Snippet:

```go
fmt.Println("Pipelines:")
pipelines, err := client.GetPipelines()
if err != nil {
	log.Fatal("Failed to get pipelines: ", err)
}
for _, pipeline := range pipelines {
	fmt.Printf("  %s:\n", pipeline.Name)
	fmt.Printf("  .Key: %s:\n", pipeline.Key)
	fmt.Printf("  .Description: %s:\n\n", pipeline.Description)
}
```
