export const DESC_SORT = (a: number, b: number) => b - a
export const ASC_SORT = (a: number, b: number) => a - b
export const range = (from: number | string, count: number | string): number[] => {
  const fromNr = typeof from == 'number' ? from : parseInt(from)
  const countNr = typeof count == 'number' ? count : parseInt(count)
  return [...new Array(countNr).keys()].map((i) => fromNr + i)
}
export const applyToRange = (from: number | string, count: number | string, apply: (val: number) => void): void => {
  const fromNr = typeof from == 'number' ? from : parseInt(from)
  const countNr = typeof count == 'number' ? count : parseInt(count)
  for (let i = fromNr; i < fromNr + countNr; i++) {
    apply(i)
  }
}
