package components

type CardProps struct {
	Title       string
	Description string
	Class       string
	MaxWidth    string
}

func NewCardProps() CardProps {
	return CardProps{
		MaxWidth: "",
	}
}
