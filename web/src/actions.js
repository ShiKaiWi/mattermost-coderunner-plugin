import { getConfig } from "mattermost-redux/selectors/entities/general"
import { serverId } from "./manifest"
import { OPEN_ROOT_MODAL, CLOSE_ROOT_MODAL } from "./action_types"

export const openRootModal = result => dispatch => {
  console.log("openRootModal")
  dispatch({
    type: OPEN_ROOT_MODAL,
    data: result
  })
}

export const closeRootModal = () => dispatch => {
  dispatch({
    type: CLOSE_ROOT_MODAL,
    data: null
  })
}

export const getPluginServerRoute = state => {
  const config = getConfig(state)

  let basePath = "/"
  if (config && config.SiteURL) {
    basePath = new URL(config.SiteURL).pathname

    if (basePath && basePath[basePath.length - 1] === "/") {
      basePath = basePath.substr(0, basePath.length - 1)
    }
  }

  return basePath + "plugins/"
}

export const codeResult = data => async (dispatch, getState) => {
  fetch(getPluginServerRoute(getState()) + serverId, {
    method: "post",
    headers: {
      "content-type": "application/json"
    },
    body: JSON.stringify(data)
  })
    .then(res => res.json())
    .then(({ result }) => {
      dispatch(openRootModal(result))
    })
    .catch(e => console.log("codeResultFailed...", e))
}
