import { Day, StringArrayInputParts } from '@advent/utils'

export const day1_2023 = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => part1.map(mapFirstLastDigit).reduce((prev, curr) => prev + curr, 0),
  part2: (): number => {
    const lettersDigitMap = new Map([
      ['one', '1'],
      ['two', '2'],
      ['three', '3'],
      ['four', '4'],
      ['five', '5'],
      ['six', '6'],
      ['seven', '7'],
      ['eight', '8'],
      ['nine', '9'],
    ])
    return part2
      .map((str) => mapStringToDigit(lettersDigitMap, str))
      .map(mapFirstLastDigit)
      .reduce((prev, curr) => prev + curr, 0)
  },
})

const mapFirstLastDigit = (str: string): number => {
  const digits = str.replace(/[^\d]/g, '')
  return parseInt(digits[0] + digits[digits.length - 1])
}

const mapStringToDigit = (map: Map<string, string>, str: string): string => {
  let entry = getNextKeyValue(map, str)
  while (entry) {
    // don't replace last character since could clear other valid letters that could use it
    str = str.replace(new RegExp(entry.key.slice(0, -1)), entry.value)
    entry = getNextKeyValue(map, str)
  }
  return str
}

const getNextKeyValue = (map: Map<string, string>, str: string): { key: string; value: string } | undefined => {
  let nextKey = ''
  let nextValue = ''
  let index = str.length
  map.forEach((val, key) => {
    const currIndex = str.indexOf(key)
    if (currIndex > -1 && currIndex < index) {
      nextKey = key
      nextValue = val
      index = currIndex
    }
  })
  return index == str.length ? undefined : { key: nextKey, value: nextValue }
}
