Streak API for Go
================================

An incomplete Go API wrapping [Streak's API](http://www.streak.com/api/).

*Warning*: this is totally incomplete. It only exposes enough for me to extract some statistics I wanted. Be prepared to extend the structs in [`streak.go`](streak.go). I'll accept pull requests!

Example:
-----------------

See [`cmd/streak_example.go`](cmd/streak_example.go) for a complete working example program. Code snippet:

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

Example output:

	Pipelines:
		Sales / CRM:
		.Key: agxzfm1haWxmb29nYWVyLgsSDE9yZ2FuaXphdGlvbiINbGVjdG9yaXVzLmNvbQwLEghXb3JrZmxvdxjRDww
		.Description: Use this pipeline to manage your sales process across your organization. Create a box for each customer.
