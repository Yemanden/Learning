package adapter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	testAdapterPass1Name = "Client to Service"
	testAdapterPass1Want = "Service!"
	testAdapterPass2Name = "Service to Client"
	testAdapterPass2Want = "Client!"
)

func TestAdapter(t *testing.T) {
	client := NewClient()
	service := NewService()
	adapter := NewAdapter(client, service)

	t.Run(testAdapterPass1Name, func(t *testing.T) {
		got := adapter.Request()
		want := testAdapterPass1Want

		assert.EqualValues(t, got, want)
	})

	t.Run(testAdapterPass2Name, func(t *testing.T) {
		got := adapter.SpecificRequest()
		want := testAdapterPass2Want

		assert.EqualValues(t, got, want)
	})
}
