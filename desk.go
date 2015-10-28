package desk

import (
  "time"
  "github.com/stianeikeland/go-rpio"
  "math"
)

type Direction uint8

const (
  Down Direction = iota
  Up
)

var (
  upPin rpio.Pin
  downPin rpio.Pin
  currentPosition float64
  start time.Time
  direction Direction
  primed bool
)

func IsPrimed() (bool) {
  return primed
}

func Prime() {
  // go all the way up
  // upPin.High()
  //
  // time.Sleep(30 * time.Second)
  // upPin.Low()

  primed = true

  currentPosition = 0.0
}

func Move(dir Direction) {
  // remember start time
  start = time.Now()

  // start moving
  if (dir == Up) {
    upPin.High()
  } else {
    downPin.High()
  }

  direction = dir
}

func MoveTo(pos float64) {
  delta := pos - currentPosition
  distance := time.Duration(math.Abs(delta))

  if (delta > 0.) {
    downPin.High()
  } else {
    upPin.High()
  }

  time.Sleep(distance * time.Second)

  downPin.Low()
  upPin.Low()

  currentPosition = pos
}

func Stop() {
  duration := time.Since(start).Seconds()

  upPin.Low()
  downPin.Low()

  if (direction == Up) {
    duration = duration * -1.0
  }

  currentPosition = currentPosition + duration

  if (currentPosition < 0) {
    currentPosition = 0
  }
}

func GetPosition() (float64) {
  return currentPosition
}

func Init(up int, down int) (err error) {
  err = rpio.Open()

  if (err != nil) {
    return
  }

  primed = true

  upPin = rpio.Pin(up)
  upPin.Output()
  upPin.Low()

  downPin = rpio.Pin(down)
  downPin.Output()
  downPin.Low()

  return nil
}

func Close() {
  rpio.Close()
}
