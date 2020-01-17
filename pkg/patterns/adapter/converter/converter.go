package converter

type (
	decimalData = []byte
	text        = string
)

// Converter ...
type Converter interface {
	Convert(decimalData) text
}

type converter struct{}

func (c *converter) Convert(decimalData decimalData) (text text) {
	for i := 0; i < len(decimalData); i++ {
		switch decimalData[i] {
		case 1:
			{
				text += "H"
			}
		case 2:
			{
				text += "e"
			}
		case 3:
			{
				text += "l"
			}
		case 4:
			{
				text += "o"
			}
		default:
			{
				text += "*"
			}
		}
	}
	return
}

// NewConverter ...
func NewConverter() Converter {
	return &converter{}
}
