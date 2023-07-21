package m3u8

import "fmt"

type RenditonReport struct{}

func NewRenditionReport() (*RenditonReport, error) {
	return &RenditonReport{}, nil
}

func (re *RenditonReport) String() string {
	return fmt.Sprintf("%s\n", RenditionReportTag)
}
