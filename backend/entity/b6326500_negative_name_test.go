package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmpNegativeName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Name cannot be null", func(t *testing.T) {
		emp := Employee{
			Name:       "",
			Email:      "aam@mail.com",
			EmployeeID: "J12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Name cannot be null"))
	})
}
