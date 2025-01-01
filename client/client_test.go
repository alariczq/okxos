package client

func ExampleClient() {
	client := NewClient("key", "secret", "passphrase",
		WithProjectID("test"),
	)
	_ = client
}
