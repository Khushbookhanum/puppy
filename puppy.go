package puppy
import(
	"github.com/Khushbookhanum/dog"
)


func Bark()string{
	return "woow"
}
func Barks()string {
	return "woow""woow""woow"
}
func Bigbark() string{
	return dog.whenGrownUp(Bark())
}
func Bigbarks() string{
	return dog.whenGrownUp(Barks())
}