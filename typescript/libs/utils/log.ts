import { appendFileSync } from 'fs'

export const logToFile = (objectOrCallback: string | (() => string), path = './test') => {
  const object = objectOrCallback instanceof Function ? objectOrCallback() : objectOrCallback
  appendFileSync(path, object)
}
