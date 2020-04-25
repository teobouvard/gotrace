package main

import (
	"math"
	"math/rand"
)

// Camera represents the camera of the scene
type Camera struct {
	origin        Vec3
	horizontal    Vec3
	vertical      Vec3
	corner        Vec3
	u, v, w       Vec3
	lensRadius    float64
	tStart, tStop float64
}

// NewCamera creates a camera
func NewCamera(lookFrom, lookAt, up Vec3, verticalFOV, aspectRatio, aperture, focusDist, tStart, tStop float64) Camera {
	theta := (math.Pi * verticalFOV) / 180.0
	height := math.Tan(theta / 2.0)
	width := aspectRatio * height

	w := lookFrom.Sub(lookAt).Unit()
	u := up.Cross(w).Unit()
	v := w.Cross(u)

	horizontal := u.Scale(2 * width * focusDist)
	vertical := v.Scale(2 * height * focusDist)

	corner := lookFrom.Sub(u.Scale(width * focusDist)).Sub(v.Scale(height * focusDist)).Sub(w.Scale(focusDist))

	return Camera{
		origin:     lookFrom,
		horizontal: horizontal,
		vertical:   vertical,
		corner:     corner,
		u:          u,
		v:          v,
		w:          w,
		lensRadius: aperture / 2.0,
		tStart:     tStart,
		tStop:      tStop,
	}
}

// RayTo returns the Ray when the camera looks at (u, v)
func (c Camera) RayTo(s float64, t float64) Ray {
	rd := RandDisk().Scale(c.lensRadius)
	offset := c.u.Scale(rd.X).Add(c.v.Scale(rd.Y))
	hOffset := c.horizontal.Scale(s)
	vOffset := c.vertical.Scale(t)
	return Ray{
		Origin:    c.origin.Add(offset),
		Direction: c.corner.Add(hOffset).Add(vOffset).Sub(c.origin).Sub(offset),
		Time:      rand.Float64()*(c.tStop-c.tStart) + c.tStart,
	}
}
