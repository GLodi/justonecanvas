import React from 'react'
import './App.css'
import MainStage from './MainStage'

function App() {
  return (
    <div className="App">
      <a
        onClick={() => window.open('https://giuliolodi.dev/', '_blank')}
        href="#"
      >
        what is this?
      </a>
      <MainStage />
    </div>
  )
}

export default App
