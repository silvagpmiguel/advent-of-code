import { Day, NumberArrayInputParts } from '@advent/utils'

type StateMap = Map<number, number> // timer, counter
export function day6_2021(inputs: NumberArrayInputParts, afterResetTimer: number, childTimer: number, days: any): Day {
  const p1Input: number[] = inputs.part1
  const p2Input: number[] = inputs.part2
  const incStateMapTimer = (stateMap: StateMap, timer: number, incValue?: number, defCounter?: number) => {
    const inc = incValue ?? 1
    const defaultCounter = defCounter ?? 1
    const actualCounter = Number(stateMap.get(timer))
    stateMap.has(timer) ? stateMap.set(timer, actualCounter + inc) : stateMap.set(timer, defaultCounter)
  }
  const initStateMap = (input: number[]): StateMap => {
    const stateMap: StateMap = new Map()
    for (let timer of input) {
      incStateMapTimer(stateMap, timer)
    }
    return stateMap
  }

  return {
    part1: (): number => {
      const stateMap: StateMap = initStateMap(p1Input)
      for (let i = 0; i < days.part1; i++) {
        let aux = new Map(stateMap)
        aux.forEach((counter, timer) => {
          if (timer == 0) {
            incStateMapTimer(stateMap, afterResetTimer, counter, counter)
            incStateMapTimer(stateMap, childTimer, counter, counter)
          } else {
            incStateMapTimer(stateMap, timer - 1, counter, counter)
          }
          incStateMapTimer(stateMap, timer, -counter)
        })
      }
      return Array.from(stateMap.values()).reduce((a, b) => a + b)
    },
    part2: (): number => {
      const stateMap: StateMap = initStateMap(p2Input)
      for (let i = 0; i < days.part2; i++) {
        let aux = new Map(stateMap)
        aux.forEach((counter, timer) => {
          if (timer == 0) {
            incStateMapTimer(stateMap, afterResetTimer, counter, counter)
            incStateMapTimer(stateMap, childTimer, counter, counter)
          } else {
            incStateMapTimer(stateMap, timer - 1, counter, counter)
          }
          incStateMapTimer(stateMap, timer, -counter)
        })
      }
      return Array.from(stateMap.values()).reduce((a, b) => a + b)
    },
  }
}
