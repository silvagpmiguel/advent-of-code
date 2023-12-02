import { Day, StringArrayInputParts } from '@advent/utils'

type Numbers = Array<number>
interface Board {
  values: Array<Array<number>>
  marked: Array<Array<boolean>>
  markedRowsCounter: Array<number>
  markedColumnCounter: Array<number>
  score: number
}
type Boards = Array<Board>

export function day4_2021(inputs: StringArrayInputParts, matrixSize: number): Day {
  const p1Input: string[] = inputs.part1
  const p2Input: string[] = inputs.part2
  const initBoard = (): Board => ({
    values: new Array(),
    marked: Array.from(Array(matrixSize), () => Array(matrixSize).fill(false)),
    markedRowsCounter: Array(matrixSize).fill(0),
    markedColumnCounter: Array(matrixSize).fill(0),
    score: 0,
  })
  const getBingoNumbers = (input: string[]): Numbers => input[0].split(',').map(Number)
  const getBingoBoards = (input: string[]): Boards => {
    const boards: Boards = new Array()
    let board: Board = initBoard()
    let i = 0
    input.slice(2, input.length).forEach((line) => {
      const splittedLine = line.trim().split(/ +/g)
      if (splittedLine.length != matrixSize) {
        i++
        boards.push(board)
        board = initBoard()
      } else {
        board.values.push(splittedLine.map(Number))
      }
    })
    return boards
  }
  const isBoardWinner = (board: Board, num: number): boolean => {
    if (board.score != 0) return false
    for (let rowPos = 0; rowPos < matrixSize; rowPos++) {
      const boardLine = board.values[rowPos]
      for (let colPos = 0; colPos < matrixSize; colPos++) {
        if (boardLine[colPos] == num) {
          board.marked[rowPos][colPos] = true
          board.markedRowsCounter[rowPos]++
          board.markedColumnCounter[colPos]++
          if (board.markedRowsCounter[rowPos] == matrixSize || board.markedColumnCounter[colPos] == matrixSize) {
            for (let i = 0; i < matrixSize; i++) {
              for (let j = 0; j < matrixSize; j++) {
                if (!board.marked[i][j]) {
                  board.score += board.values[i][j]
                }
              }
            }
            return true
          }
        }
      }
    }
    return false
  }

  return {
    part1: (): number => {
      const numbers = getBingoNumbers(p1Input)
      const boards = getBingoBoards(p1Input)
      const winners: any = []
      numbers.forEach((num: number) => {
        boards.forEach((board: Board) => {
          if (isBoardWinner(board, num)) {
            winners.push(board.score * num)
          }
        })
      })
      return winners[0] ?? 0
    },
    part2: (): number => {
      const numbers = getBingoNumbers(p2Input)
      const boards = getBingoBoards(p2Input)
      const winners: any = []
      numbers.forEach((num: number) => {
        boards.forEach((board: Board) => {
          if (isBoardWinner(board, num)) {
            winners.push(board.score * num)
          }
        })
      })
      return winners[winners.length - 1] ?? 0
    },
  }
}
