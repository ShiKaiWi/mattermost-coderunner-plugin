import React from "react"

import { codeResult } from "./actions"

const jsCodeSupport = ["js", "javascript"]
const goCodeSupport = ["go", "golang"]

/**
 * get posts message by postId from state
 * @param {string} postId
 * @param {object} state
 * @return {string}
 */
const getMessageById = (postId, state) => {
  const { posts } = state.entities.posts
  const msgObj = posts[postId]
  return msgObj.message
}

/**
 * check if the message is js code
 * @param {string} message
 * @return {boolean}
 */
const isJsSupport = message => {
  const arr = message.split("```")[1].split("\n")
  return jsCodeSupport.includes(arr[0].trim().toLocaleLowerCase())
}

/**
 * check if the message is go code
 * @param {string} message
 * @return {boolean}
 */
const isGoSupport = message => {
  const arr = message.split("```")[1].split("\n")
  return goCodeSupport.includes(arr[0].trim().toLocaleLowerCase())
}

/**
 * @param {string} message
 * @return {boolean}
 */
const isCodeMessage = message => {
  if (!message.startsWith("```")) {
    return false
  }
  return isJsSupport(message) || isGoSupport(message)
}

/**
 * format codeResult params from message
 * @param {string} message
 * @return {object}
 */
const codeMessageRequest = message => {
  const nIndex = message.indexOf("\n")
  const msg = message.slice(nIndex).split("```")[0]
  const isJs = isJsSupport(message)

  return {
    code: msg,
    lang: isJs ? "javascript" : "golang"
  }
}
class CodeRunnerPlugin {
  initialize(registry, store) {
    registry.registerPostDropdownMenuAction(
      "Run",
      id => {
        const meg = getMessageById(id, store.getState())
        store.dispatch(codeResult(codeMessageRequest(meg)))
      },
      postId => {
        const ms = getMessageById(postId, store.getState())
        return isCodeMessage(ms)
      }
    )
  }
}

window.registerPlugin(
  "com.mattermost.web-coderunner-plugin",
  new CodeRunnerPlugin()
)
