import React from 'react'
import { Stage, Layer } from 'react-konva'
import './App.css'
import Square from './Square'

function App() {
  return (
    <div className="App">
      <Stage width={window.innerWidth} height={window.innerHeight}>
        <Layer>
          <Square />
        </Layer>
      </Stage>
    </div>
  )
}

export default App
