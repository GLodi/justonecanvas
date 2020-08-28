import React from 'react'
import './Main.css'
import MainStage from './MainStage'

function Main() {
  return (
    <div className="App">
      <a
        onClick={() => window.open('http://justonecanvas.dev/about', '_blank')}
        href="/#"
      >
        what is this?
      </a>
      <MainStage />
    </div>
  )
}

export default Main
