import { StringArrayInputParts, Day } from '@advent/utils'

type TimeDistance = {
  time: number
  distance: number
}
type Document = TimeDistance[]

export const day6_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => createDocument1(part1).reduce((prev, { distance, time }) => prev * getNrWays(distance, time), 1),
  part2: (): number => createDocument2(part2).reduce((prev, { distance, time }) => prev * getNrWays(distance, time), 1),
})

const createDocument1 = (input: string[]): Document => {
  const time = input[0]
    .split(':')[1]
    .trim()
    .split(/\s+/)
    .map((x) => parseInt(x))
  const distance = input[1]
    .split(':')[1]
    .trim()
    .split(/\s+/)
    .map((x) => parseInt(x))
  return time.map((time, index) => ({ time, distance: distance[index] }))
}
const createDocument2 = (input: string[]): Document => {
  const time = parseInt(input[0].split(':')[1].replace(/\s+/g, ''))
  const distance = parseInt(input[1].split(':')[1].replace(/\s+/g, ''))
  return [{ time, distance }]
}
const getNrWays = (distance: number, time: number, holdTime = 1, maxHoldTime = time) => {
  let nrWays = 0
  for (let speed = holdTime; speed < maxHoldTime; speed++) {
    const currentTime = time - speed
    const currentDistance = currentTime * speed
    if (currentDistance > distance) {
      nrWays++
    }
  }
  return nrWays
}
