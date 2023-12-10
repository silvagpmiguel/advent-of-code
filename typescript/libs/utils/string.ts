const NUMBERS = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9']
export const isNumber = (str: string) => NUMBERS.includes(str)
export const isAsterisk = (str: string) => str == '*'
export const hasSymbol = (str: string) => /[^.\w]/g.test(str)
export const hasNumber = (str: string) => /\d+/g.test(str)
export const countOccurrences = (str: string, substr: string): number =>
  (str?.match(new RegExp(substr, 'g')) ?? []).length
