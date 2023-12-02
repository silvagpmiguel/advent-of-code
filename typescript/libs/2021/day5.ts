import { Day, StringArrayInputParts } from '@advent/utils'

interface Coord {
  x: number
  y: number
}
type Counter = number
type Diagram = Map<string, Counter>

export function day5_2021(inputs: StringArrayInputParts): Day {
  const p1Input: string[] = inputs.part1
  const p2Input: string[] = inputs.part2
  const getLineSegment = (first: Coord, second: Coord): Array<Coord> => {
    const lineSegment: Array<Coord> = []
    const incX = first.x == second.x ? 0 : first.x > second.x ? -1 : 1
    const incY = first.y == second.y ? 0 : first.y > second.y ? -1 : 1
    let coordX = first.x,
      coordY = first.y
    for (; coordX != second.x || coordY != second.y; coordX += incX, coordY += incY) {
      lineSegment.push({ x: coordX, y: coordY })
    }
    lineSegment.push({ x: coordX, y: coordY })
    return lineSegment
  }

  return {
    part1: (): number => {
      const diagram: Diagram = new Map()
      let overlapCounter = 0
      for (let line of p1Input) {
        const splittedLine = line.split(' -> ')
        const [x1, y1] = splittedLine[0].split(',').map(Number)
        const [x2, y2] = splittedLine[1].split(',').map(Number)
        if (x1 == x2 || y1 == y2) {
          for (let entry of getLineSegment({ x: x1, y: y1 }, { x: x2, y: y2 })) {
            const key = entry.x + ',' + entry.y
            const value = Number(diagram.get(key))
            if (diagram.has(key)) {
              diagram.set(key, value + 1)
              if (value == 1) {
                overlapCounter++
              }
            } else {
              diagram.set(key, 1)
            }
          }
        }
      }
      return overlapCounter
    },
    part2: (): number => {
      const diagram: Diagram = new Map()
      let overlapCounter = 0
      for (let line of p2Input) {
        const splittedLine = line.split(' -> ')
        const [x1, y1] = splittedLine[0].split(',').map(Number)
        const [x2, y2] = splittedLine[1].split(',').map(Number)
        for (let entry of getLineSegment({ x: x1, y: y1 }, { x: x2, y: y2 })) {
          const key = entry.x + ',' + entry.y
          const value = Number(diagram.get(key))
          if (diagram.has(key)) {
            diagram.set(key, value + 1)
            if (value == 1) {
              overlapCounter++
            }
          } else {
            diagram.set(key, 1)
          }
        }
      }
      return overlapCounter
    },
  }
}
