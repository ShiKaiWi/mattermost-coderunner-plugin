import { combineReducers } from "redux"

import { OPEN_ROOT_MODAL, CLOSE_ROOT_MODAL } from "./action_types"

const rootModalVisible = (state = false, action) => {
  switch (action.type) {
    case OPEN_ROOT_MODAL:
      return action.data
    case CLOSE_ROOT_MODAL:
      return action.data
    default:
      return state
  }
}

export default combineReducers({
  rootModalVisible
})
