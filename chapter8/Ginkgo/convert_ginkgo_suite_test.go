package main__test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestConvertGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConvertGinkgo Suite")
}
