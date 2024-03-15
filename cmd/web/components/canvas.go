package components

type CanvasColor struct {
	Color  string
	Border string
}

func CanvasColors() []CanvasColor {
	colors := make([]CanvasColor, 12)
	colors[0] = CanvasColor{Color: "#FF0000"}
	colors[1] = CanvasColor{Color: "#00FF00"}
	colors[2] = CanvasColor{Color: "#0000FF"}
	colors[3] = CanvasColor{Color: "#FFFF00"}
	colors[4] = CanvasColor{Color: "#FFFFFF", Border: "border-top"}
	colors[5] = CanvasColor{Color: "#7C4700"}
	colors[6] = CanvasColor{Color: "#FFC0CB"}
	colors[7] = CanvasColor{Color: "#00AA00"}
	colors[8] = CanvasColor{Color: "#0000AA"}
	colors[9] = CanvasColor{Color: "#FFAA00"}
	colors[10] = CanvasColor{Color: "#000000"}
	colors[11] = CanvasColor{Color: "#999999"}
	return colors
}
