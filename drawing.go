package main

import (
	"canvas"
	"image"
)

/*
Not necessary: Was for if I could get soft body physics working.
ADD THIS TO canvas.go. NECESSARY FOR DRAWING POLYGONS.
func (c *Canvas) BeginPath() {
	c.gc.BeginPath()
}

func (c *Canvas) Close() {
	c.gc.Close()
}
*/

// AnimateSystem takes a slice of Universe objects along with a canvas width
// parameter and a frequency parameter.
// Every frequency steps, it generates a slice of images corresponding to drawing each Universe
// on a canvasWidth x canvasWidth canvas.
// A scaling factor is a final input that is used to scale the stars big enough to see them.
func DrawECM(timePoints []*ECM, canvasWidth, frequency int, scalingFactor float64) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no Universe objects present in AnimateSystem.")
	}

	// for every universe, draw to canvas and grab the image
	for i := range timePoints {
		images = append(images, timePoints[i].DrawToCanvas(canvasWidth, scalingFactor))
	}

	return images
}

// DrawToCanvas generates the image corresponding to a canvas after drawing a ECM
// object's bodies on a square canvas that is canvasWidth pixels x canvasWidth pixels.
func (e *ECM) DrawToCanvas(canvasWidth int, scalingFactor float64) image.Image {
	if e == nil {
		panic("Can't Draw a nil ECM.")
	}

	// set a new square canvas
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	// create a black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// Draw all the fibres
	for _, f := range e.fibres {
		center_x := (f.position.x / ECMwidth) * float64(canvasWidth)
		center_y := (f.position.y / ECMwidth) * float64(canvasWidth)
		direction := f.direction
		magnitude := f.direction.Magnitude()
		direction.x *= 0.5 * f.length / magnitude * float64(canvasWidth) / ECMwidth
		direction.y *= 0.5 * f.length / magnitude * float64(canvasWidth) / ECMwidth

		c.SetLineWidth(f.width / ECMwidth * float64(canvasWidth))
		c.SetStrokeColor(canvas.MakeColor(100, 100, 200))
		c.MoveTo(center_x-direction.x, center_y-direction.y)
		c.LineTo(center_x+direction.x, center_y+direction.y)
		c.Stroke()
		c.FillStroke()
	}

	// range over all the bodies and draw them.
	for _, c1 := range e.cells {
		c.SetFillColor(canvas.MakeColor(200, 150, 200))
		cx := (c1.position.x / ECMwidth) * float64(canvasWidth)
		cy := (c1.position.y / ECMwidth) * float64(canvasWidth)
		r := scalingFactor * (c1.radius / ECMwidth) * float64(canvasWidth)
		c.Circle(cx, cy, r)
		c.Fill()

		/*
			//This is for if soft body works
			c.SetFillColor(canvas.MakeColor(200, 150, 200))
			c.BeginPath()
			for i := range c1.perimeterVertices {
				x := c1.perimeterVertices[i].x / ECMwidth * float64(canvasWidth)
				y := c1.perimeterVertices[i].y / ECMwidth * float64(canvasWidth)
				if i == 0 {
					c.MoveTo(x, y)
				} else {
					c.LineTo(x, y)
				}
			}
			c.LineTo(c1.perimeterVertices[0].x/ECMwidth*float64(canvasWidth),
				c1.perimeterVertices[0].y/ECMwidth*float64(canvasWidth))
			c.Close()
			c.FillStroke()
		*/
	}

	// we want to return an image!
	return c.GetImage()
}
