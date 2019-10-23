package feature

import (
	"math"
)

/*Object type contains all necessary component of an object
 *Object contains material, tuple and matrix*/
type Object struct {
	Mat       Material
	Center    Tuple
	Transform *Matrix
}

/*NewObject gives a default object with default material
 *NewObject returns an object*/
func NewObject() *Object {
	matrix := NewMatrix(4, 4)
	matrix, _ = matrix.GetIdentity()
	o := &Object{
		Mat:       *NewMaterial(),
		Center:    *Point(0, 0, 0),
		Transform: matrix,
	}
	return o
}

/*SetTransform sets the transform matrix
 *SetTransform can only be called by an object
 *SetTransform takes in a matrix
 *SetTransform returns an object*/
func (obj *Object) SetTransform(matrix *Matrix) *Object {
	obj.Transform = matrix
	return obj
}

/*Translate translate the ray with a matrix
 *Translate is a public function
 *Translate takes in three float
 *Translate returns a matrix*/
func Translate(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(3, 0, xInc)
	m = m.Assign(3, 1, yInc)
	m = m.Assign(3, 2, zInc)
	return m
}

/*Scale scales the ray
 *Scale is a public function
 *Scale takes in three float
 *Scale returns a matrix*/
func Scale(xInc, yInc, zInc float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, xInc)
	m = m.Assign(1, 1, yInc)
	m = m.Assign(2, 2, zInc)
	return m
}

/*RotationX rotates the ray around x axis
 *RotationX is a public function
 *RotationX takes in one float
 *RotationX returns a matrix*/
func RotationX(rot float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 1, math.Cos(rot)+3-3)
	m = m.Assign(2, 1, -math.Sin(rot)+3-3)
	m = m.Assign(1, 2, math.Sin(rot)+3-3)
	m = m.Assign(2, 2, math.Cos(rot)+3-3)
	return m
}

/*RotationY rotates the ray around y axis
 *RotationY is a public function
 *RotationY takes in one float
 *RotationY returns a matrix*/
func RotationY(rot float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(rot)+3-3)
	m = m.Assign(2, 0, math.Sin(rot)+3-3)
	m = m.Assign(0, 2, -math.Sin(rot)+3-3)
	m = m.Assign(2, 2, math.Cos(rot)+3-3)
	return m
}

/*RotationZ rotates the ray around z axis
 *RotationZ is a public function
 *RotationZ takes in one float
 *RotationZ returns a matrix*/
func RotationZ(rot float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(0, 0, math.Cos(rot)+3-3)
	m = m.Assign(1, 0, -math.Sin(rot)+3-3)
	m = m.Assign(0, 1, math.Sin(rot)+3-3)
	m = m.Assign(1, 1, math.Cos(rot)+3-3)
	return m
}

/*Shearing makes the straight line slanted
 *Shearing is a public function
 *Shearing takes in six float
 *Shearing returns a matrix*/
func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	m = m.Assign(1, 0, xy)
	m = m.Assign(2, 0, xz)
	m = m.Assign(0, 1, yx)
	m = m.Assign(2, 1, yz)
	m = m.Assign(0, 2, zx)
	m = m.Assign(1, 2, zy)
	return m
}

/*ViewTransformation changes the view orientation
 *ViewTransformation is a public function
 *ViewTransformation takes in three tuple
 *ViewTransformation returns a matrix*/
func ViewTransformation(from, to, up Tuple) *Matrix {
	m := NewMatrix(4, 4)
	m, _ = m.GetIdentity()
	subtract, _ := to.Subtract(&from)
	forward, _ := subtract.Normalize()
	upn, _ := up.Normalize()
	left, _ := forward.CrossProduct(&upn)
	trueUp, _ := left.CrossProduct(&forward)
	m = m.Assign(0, 0, left.X)
	m = m.Assign(1, 0, left.Y)
	m = m.Assign(2, 0, left.Z)
	m = m.Assign(0, 1, trueUp.X)
	m = m.Assign(1, 1, trueUp.Y)
	m = m.Assign(2, 1, trueUp.Z)
	m = m.Assign(0, 2, -forward.X)
	m = m.Assign(1, 2, -forward.Y)
	m = m.Assign(2, 2, -forward.Z)
	trans := Translate(-from.X, -from.Y, -from.Z)
	ans, _ := m.Multiply(trans)
	return ans
}
