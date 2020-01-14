package adapter

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Yemanden/Learning/pkg/patterns/adapter/adapter"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/client"
	"github.com/Yemanden/Learning/pkg/patterns/adapter/service"
)

const (
	testAdapterPass1Name = "Client to Service"
	testAdapterPass1Want = "Service!"

	testAdapterPass2Name = "Service to Client"
	testAdapterPass2Want = "Client!"

	testAdapterServiceInput = "Service!"
	testAdapterClientInput  = "Client!"
)

func TestAdapter(t *testing.T) {
	cl := client.NewClient(testAdapterClientInput)
	serv := service.NewService(testAdapterServiceInput)
	adapt := adapter.NewAdapter(cl, serv)

	t.Run(testAdapterPass1Name, func(t *testing.T) {
		want := testAdapterPass1Want
		got := adapt.Request()

		assert.EqualValues(t, want, got)
	})

	t.Run(testAdapterPass2Name, func(t *testing.T) {
		want := testAdapterPass2Want
		got := adapt.SpecificRequest()

		assert.EqualValues(t, want, got)
	})
}
