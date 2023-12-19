package main

func CalculatePositionDelta(pos1 int32, pos2 int32) int32 {
  return (pos2 - pos1)
}

func CalculateTimeDelta(time1 int32, time2 int32) int32 {
  return (time2 - time1)
}

func Speed(posDelta int32, timeDelta int32) int32 {
  return (posDelta / timeDelta)
}
