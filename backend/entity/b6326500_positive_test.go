package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name       string `valid:"required~Nam cannot be null"`
	Email      string
	EmployeeID string `valid:"matches([JMS][0-9]{8}$)"`
}

func TestEmpPass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	t.Run("Employee Pass Test", func(t *testing.T) {
		emp := Employee{
			Name:       "Sirinya",
			Email:      "aam@mail.com",
			EmployeeID: "J12345678",
		}
		ok, err := govalidator.ValidateStruct(emp)
		g.Expect(ok).To(gomega.BeTrue())
		g.Expect(err).To(gomega.BeNil())
	})
}
