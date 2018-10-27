import { id as pluginId } from "./manifest"

const getPluginState = state => state["plugins-" + pluginId] || {}

export const isRootModalVisible = state => {
  return getPluginState(state).rootModalVisible
}
