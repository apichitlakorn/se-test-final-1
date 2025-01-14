package unit

import (
	"testing"

	"github.com/apichitlakorn/se-test-final-1/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStudentIDTwo(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("StudentID", func(t *testing.T) {
		user := entity.User{
			StudentID: "B6512767",
			Email:     "a",
			Phone:     "0802199327",
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid format Email"))

	})

	t.Run("Phone must be over 10 number", func(t *testing.T) {
		g := NewGomegaWithT(t)
		user := entity.User{
			StudentID: "B6512767",
			Email:     "a@gmail.com",
			Phone:     "0",
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("not format for phone"))
	})
}
