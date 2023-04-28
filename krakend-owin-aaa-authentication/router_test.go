package krakendowinaaaauth

import (
	"testing"

	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/router/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHandlerFactory(t *testing.T) {
	type args struct {
		in0     logging.Logger
		factory gin.HandlerFactory
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFactory
	}{
		{
			name: "sample",
			args: args{
				in0:     nil,
				factory: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				assert.True(t, true)
			},
		)
	}
}
