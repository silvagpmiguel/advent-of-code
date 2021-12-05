export const sumBetween = (i: number, j: number, arr: number[]): number => {
  let sum = 0
  for (i; i < j; i++) {
    sum += arr[i]
  }
  return sum
}
