package method

import (
	"feature"
)

//FirstImage creates the first image
func FirstImage() {
	var x, y float64
	canv := feature.NewCanvas(100, 100)
	red := feature.NewColor(1, 0, 0)
	s := feature.NewSphere()
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, _, _, intersect := s.IntersectWithRay(ray)
			if intersect {
				canv.WritePixel(i, j, red)
			}
		}
	}
	canv.CanvasToPPM("../../output/test1")
}

//SecondImage creates the second image
func SecondImage(fileName string) {
	var x, y float64
	canv := feature.NewCanvas(100, 100)
	red := feature.NewColor(1, 0, 0)
	s := feature.NewSphere()
	//Add Light in here
	l := feature.NewLight()
	l.PointLight(*feature.Point(1, 0, -5), *feature.NewColor(0.5, 1, 0.3))
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, ans1, _, intersect := s.IntersectWithRay(ray)
			if intersect {
				color := DiffuseLight(ans1, l, ray, red, *s)
				canv.WritePixel(i, j, color)
			}
		}
	}
	canv.CanvasToPPM(fileName)
}

//DiffuseLight diffuses Light
func DiffuseLight(intersection float64, l *feature.Light, ray *feature.Ray, originalColor *feature.Color, s feature.Sphere) *feature.Color {
	hitPoint := ray.Position(intersection)
	vectorToLight, _ := l.Position.Subtract(&hitPoint)
	unitVectorToLight, _ := vectorToLight.Normalize()
	normal := s.NormalAt(&hitPoint)
	lightIntensity, _ := normal.DotProduct(&unitVectorToLight)
	colorAtPoint := l.Intensity.ColorMultiply(originalColor)
	color := colorAtPoint.Multiply(lightIntensity)
	return &color
}

//ThirdImage creates the third image
func ThirdImage(fileName string) {
	var x, y float64

	canv := feature.NewCanvas(100, 100)
	s := feature.NewSphere()
	s.Material.Col = *feature.NewColor(1, 0.2, 1)
	//Add Light in here
	l := feature.NewLight()
	*l = l.PointLight(*feature.Point(-10, 10, -10), *feature.NewColor(1, 1, 1))
	for i, row := range canv.Canv {
		y = float64(2) - float64(i)/float64(25)
		for j := range row {
			x = float64(-2) + float64(j)/float64(25)
			ray := feature.NewRay(*feature.Point(x, y, -5), *feature.Vector(0, 0, 1))
			_, ans1, _, intersect := s.IntersectWithRay(ray)
			hitPoint := ray.Position(ans1)
			eye := ray.Direction.Multiply(-1)
			normal := s.NormalAt(&hitPoint)
			if intersect {
				color := s.Material.Lighting(*l, &hitPoint, &normal, &eye)
				canv.WritePixel(i, j, &color)
			}
		}
	}
	canv.CanvasToPPM(fileName)
}
