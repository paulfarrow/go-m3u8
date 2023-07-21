package test

import (
	"testing"

	"github.com/etherlabsio/go-m3u8/m3u8"
	"github.com/stretchr/testify/assert"
)

func TestRenditionReport_Parse(t *testing.T) {
	re, err := m3u8.NewRenditionReport()
	assert.Nil(t, err)
	assert.Equal(t, m3u8.RenditionReportTag+"\n", re.String())
}
