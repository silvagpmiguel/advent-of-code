import { StringArrayInputParts, Day } from '@advent/utils'

type Config = Record<'start' | 'end', string>
type Node = {
  L?: string
  R?: string
  value: string
}
type NodeMap = {
  instructions: Array<keyof Node>
  nodes: Map<string, Node>
  root: string
}
type MultiRootNodeMap = {
  instructions: Array<keyof Node>
  nodes: Map<string, Node>
  roots: string[]
}

export const day8_2023 = ({ part1, part2 }: StringArrayInputParts, p1Config: Config, p2Config: Config): Day => ({
  part1: (): number => countNodeMapSteps(createNodeMap(part1, p1Config.start), p1Config.end),
  part2: (): number => countMultiRootNodeMapSteps(createMultiRootNodeMap(part2, p2Config.start), p2Config.end),
})

const createNodeMap = (input: string[], start: string): NodeMap => {
  const nodes = new Map()
  const instructions = input[0].split('') as Array<keyof Node>
  input.slice(2).forEach((line) => {
    const [node, targetNodesStr] = line.split(' = ')
    const [L, R] = targetNodesStr.split(', ').map((str, i) => (i == 0 ? str.substring(1) : str.slice(0, -1)))
    nodes.set(node, { L, R, value: node })
  })
  return { instructions, nodes, root: start }
}
const createMultiRootNodeMap = (input: string[], start: string): MultiRootNodeMap => {
  const nodes = new Map()
  const instructions = input[0].split('') as Array<keyof Node>
  const roots: string[] = []
  input.slice(2).forEach((line) => {
    const [node, targetNodesStr] = line.split(' = ')
    const [L, R] = targetNodesStr.split(', ').map((str, i) => (i == 0 ? str.substring(1) : str.slice(0, -1)))
    if (node.endsWith(start)) {
      roots.push(node)
    }
    nodes.set(node, { L, R, value: node })
  })
  return { instructions, nodes, roots }
}
const countNodeMapSteps = (nodeMap: NodeMap, end: string) => {
  let steps = 0
  let currentNode = nodeMap.nodes.get(nodeMap.root)
  while (currentNode?.value != end) {
    nodeMap.instructions.forEach((instruction) => {
      steps++
      currentNode = nodeMap.nodes.get(currentNode![instruction] as string)
    })
  }
  return steps
}
const countMultiRootNodeMapSteps = (multiRootNodeMap: MultiRootNodeMap, end: string) => {
  const currentNodes = multiRootNodeMap.roots.map((root) => multiRootNodeMap.nodes.get(root)!.value)
  let steps = 0
  while (currentNodes.some((node) => !node.endsWith(end))) {
    multiRootNodeMap.instructions.forEach((instruction) => {
      for (let i = 0; i < currentNodes.length; i++) {
        const node = multiRootNodeMap.nodes.get(currentNodes[i]) as Node
        currentNodes[i] = node[instruction]!
      }
    })
    steps += multiRootNodeMap.instructions.length
  }
  return steps
}
