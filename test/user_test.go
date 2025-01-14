package unit

import (
	"testing"

	"github.com/apichitlakorn/se-test-final-1/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestStudentID(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`student_id is required`, func(t *testing.T) {
		user := entity.User{
			StudentID: "",
			Email:     "apichit@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())                              // ต้องเป็น false
		g.Expect(err).NotTo(BeNil())                              // err ต้องไม่เป็น nil
		g.Expect(err.Error()).To(Equal("Student ID is required")) // ข้อความ error ต้องตรงกัน

	})

}
func TestStudentTrue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`StudentID pattern is True`, func(t *testing.T) {
		user := entity.User{
			StudentID: "B6512767",
			Email:     "apichit@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestStudentNotTrue(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`StudentID pattern is NotTrue`, func(t *testing.T) {
		user := entity.User{
			StudentID: "C6512767",
			Email:     "apichit@gmail.com",
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid Student ID format"))

	})
}
