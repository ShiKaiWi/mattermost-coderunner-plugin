import React from "react"
import PropTypes from "prop-types"

const Icon = () => <i className="icon close" />

const Root = ({ visible, result, close, theme }) => {
  if (!visible) {
    return null
  }

  const style = getStyle(theme)

  return (
    <div style={style.backdrop}>
      <div style={style.modal}>
        <div style={style.header}>
          Your Code Result:
          <span onClick={close}>
            <svg
              style={style.svg}
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
              class="feather feather-x-circle">
              <circle cx="12" cy="12" r="10" />
              <line x1="15" y1="9" x2="9" y2="15" />
              <line x1="9" y1="9" x2="15" y2="15" />
            </svg>
          </span>
        </div>
        <pre style={style.pre}>{visible}</pre>
        <br />
      </div>
    </div>
  )
}

Root.propTypes = {
  visible: PropTypes.bool.isRequired,
  close: PropTypes.func.isRequired,
  theme: PropTypes.object.isRequired
}

const getStyle = theme => ({
  backdrop: {
    position: "absolute",
    display: "flex",
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: "rgba(0, 0, 0, 0.50)",
    zIndex: 2000,
    alignItems: "center",
    justifyContent: "center"
  },
  modal: {
    height: "400px",
    width: "600px",
    color: "white",
    backgroundColor: "white"
  },
  pre: {
    height: "360px",
    overflow: scroll,
    color: "black",

    padding: "10px"
  },
  header: {
    display: "flex",
    justifyContent: "space-between",
    backgroundColor: "#1f7bc1",
    color: "white",
    fontSize: "20px",
    padding: "10px"
  },
  svg: {
    cursor: "pointer"
  }
})

export default Root
