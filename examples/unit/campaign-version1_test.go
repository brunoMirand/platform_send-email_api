
import "testing"

func TestNewCampaign(t *testing.T) {
	//given
	name := "Aprenda PHP em 2024"
	content := "Seja um dev destaque todas as semanas da sprint."
	contacts := []string{"bruno.sm@gmail.com", "bruno.ms@gmail.com"}

	//when
	sut := NewCampaign(name, content, contacts)

	//then
	if sut.Name != "Aprenda PHP em 2024" {
		t.Error("Expected: Aprenda PHP em 2024, but receive:", sut.Name)
	}
}
