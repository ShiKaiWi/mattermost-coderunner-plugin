import React from "react"

import { codeResult } from "./actions"

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
 * check message is js or go code
 * @param {string} message
 * @return {boolean}
 */
const isCodeMessage = message => {
  if (!message.startsWith("```")) {
    return false
  }

  const arr = message.split("```")[1].split("\n")
  if (arr[0].includes("js") || arr[0].includes("go")) {
    return true
  }
  return false
}

/**
 * format codeResult params from message
 * @param {string} message
 * @return {object}
 */
const codeMessageRequest = message => {
  const nIndex = message.indexOf("\n")
  const msg = message.slice(nIndex).split("```")[0]
  const isJs = message.slice(0, nIndex).endsWith("js")

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
