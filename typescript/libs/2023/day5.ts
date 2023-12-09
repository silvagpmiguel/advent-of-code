import { StringArrayInputParts, Day, hasNumber, applyToRange } from '@advent/utils'

type Range = {
  destStart: number
  sourceStart: number
  len: number
}
type RangeGroup = Range[]
type Almanac = {
  seeds: number[]
  groups: RangeGroup[]
}

export const day5_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => computeLocation(createAlmanac(part1)),
  part2: (): number => computeLocation2(createAlmanac(part2)),
})

const createAlmanac = (input: string[]): Almanac => {
  const seeds = input[0]
    .split(':', 2)[1]
    .trim()
    .split(' ')
    .map((v) => parseInt(v))
  const groups: RangeGroup[] = []
  let group: RangeGroup = []
  input.slice(2).forEach((line) => {
    if (hasNumber(line)) {
      const [dest, source, len] = line.split(' ')
      group.push({ destStart: parseInt(dest), sourceStart: parseInt(source), len: parseInt(len) })
    } else if (group.length != 0) {
      groups.push(group)
      group = []
    }
  })
  if (group.length != 0) {
    groups.push(group)
  }
  return { seeds, groups }
}
const computeLocation2 = ({ seeds, groups }: Almanac): number => {
  let minLocation = Number.MAX_SAFE_INTEGER
  return seeds.reduce((_: number, initialSeed: number, index: number) => {
    if (index % 2 != 0) {
      console.log(`compute location: ${seeds[index - 1]}, length: ${initialSeed}`)
      applyToRange(seeds[index - 1], initialSeed, (seed) => {
        const value = mapSeed(seed, groups)
        if (value < minLocation) {
          minLocation = value
        }
      })
    }
    return minLocation
  }, minLocation)
}
const computeLocation = ({ seeds, groups }: Almanac): number => {
  let minLocation = Number.MAX_SAFE_INTEGER
  return seeds.reduce((_: number, seed: number) => {
    const value = mapSeed(seed, groups)
    if (value < minLocation) {
      minLocation = value
    }
    return minLocation
  }, minLocation)
}
const mapSeed = (seed: number, groups: RangeGroup[]): number =>
  groups.reduce((prev, group) => mapGroup(prev, group), seed)
const mapGroup = (seed: number, group: RangeGroup): number => {
  const possibleRange = group.find(({ sourceStart, len }) => seed >= sourceStart && seed < sourceStart + len)
  return possibleRange ? mapRange(seed, possibleRange) : seed
}
const mapRange = (seed: number, { sourceStart, destStart }: Range): number => destStart + (seed - sourceStart)
