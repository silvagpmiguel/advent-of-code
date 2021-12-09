import { NumberArrayInputParts } from '../parse/input'
import { Day } from './Day'

export default function Day7(inputs:NumberArrayInputParts): Day {
  const p1Input: number[] = inputs.part1
  const p2Input: number[] = inputs.part2
  return {
    part1: (): number => {
        let minSum: number = Number.MAX_SAFE_INTEGER
        for(let pos = Math.min(...p1Input); pos <= Math.max(...p1Input); pos++) {
            let sum = 0;
            for(let j = 0; j < p1Input.length; j++) {
                sum += Math.abs(p1Input[j] - pos)
            }
            if(sum < minSum) {
                minSum = sum
            }
        }
        return minSum
    },
    part2: (): number => {
        let minSum: number = Number.MAX_SAFE_INTEGER
        for(let pos = Math.min(...p2Input); pos <= Math.max(...p2Input); pos++) {
            let sum = 0;
            for(let j = 0; j < p2Input.length; j++) {
                const change = Math.abs(p2Input[j] - pos)
                sum += change
                for(let step = 1; step < change; step++) {
                    sum +=step
                }
            }
            if(sum < minSum) {
                minSum = sum
            }
        }
        return minSum
    },
  }
}
