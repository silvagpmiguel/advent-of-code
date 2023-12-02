import { Day, StringArrayInputParts } from '@advent/utils'

export function day2_2021(inputs: StringArrayInputParts): Day {
  const p1Input: string[] = inputs.part1
  const p2Input: string[] = inputs.part2
  return {
    part1: (): number => {
      const submarine = {
        horizontal_pos: 0,
        depth: 0,
      }
      for (const rawCommand of p1Input) {
        const command: string[] = rawCommand.split(' ')
        const action: string = command[0]
        const value: number = Number(command[1])
        switch (action) {
          case 'forward':
            submarine.horizontal_pos += value
            break
          case 'down':
            submarine.depth += value
            break
          case 'up':
            submarine.depth -= value
            break
          default:
            throw 'Wrong action'
        }
      }
      return submarine.horizontal_pos * submarine.depth
    },
    part2: (): number => {
      const submarine = {
        horizontal_pos: 0,
        depth: 0,
        aim: 0,
      }
      for (const rawCommand of p2Input) {
        const command: string[] = rawCommand.split(' ')
        const action: string = command[0]
        const value: number = Number(command[1])
        switch (action) {
          case 'forward':
            submarine.horizontal_pos += value
            submarine.depth += submarine.aim * value
            break
          case 'down':
            submarine.aim += value
            break
          case 'up':
            submarine.aim -= value
            break
          default:
            throw 'Wrong action'
        }
      }
      return submarine.horizontal_pos * submarine.depth
    },
  }
}
