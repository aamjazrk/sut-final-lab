package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestEmpNegativeEmployee(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Employee not match fomat", func(t *testing.T) {
		emp := Employee{
			Name:       "AAmjazrk",
			Email:      "aam@mail.com",
			EmployeeID: "J123456",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Employee not match fomat"))
	})
	t.Run("Employee not match fomat", func(t *testing.T) {
		emp := Employee{
			Name:       "AAmjazrk",
			Email:      "aam@mail.com",
			EmployeeID: "Y12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Employee not match fomat"))
	})
	t.Run("Employee not match fomat", func(t *testing.T) {
		emp := Employee{
			Name:       "AAmjazrk",
			Email:      "aam@mail.com",
			EmployeeID: "z02345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).ToNot(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("Employee not match fomat"))
	})
}
