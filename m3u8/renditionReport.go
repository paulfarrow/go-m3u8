package m3u8

import "fmt"

type RenditionReport struct{}

func NewRenditionReport() (*RenditionReport, error) {
	return &RenditionReport{}, nil
}

func (re *RenditionReport) String() string {
	return fmt.Sprintf("%s\n", RenditionReportTag)
}
