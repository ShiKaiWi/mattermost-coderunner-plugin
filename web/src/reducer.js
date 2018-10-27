import { combineReducers } from "redux"

import { STATUS_CHANGE } from "./action_types"

const code = (state = false, action) => {
  switch (action.type) {
    case STATUS_CHANGE:
      return action.data

    default:
      return state
  }
}

export default combineReducers({
  code
})
