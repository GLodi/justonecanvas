import React from 'react'
import { Stage, Layer } from 'react-konva'
import './App.css'
import Square from './Square'
import Konva from 'konva'

const squareAmount = 2500
const squarePerRow = Math.sqrt(squareAmount)
const squareSize = 25

function App() {
  var rows = []
  for (var i = 0; i < squareAmount; i++) {
    // note: we add a key prop here to allow react to uniquely identify each
    // element in this array. see: https://reactjs.org/docs/lists-and-keys.html
    rows.push(
      <Square
        key={i}
        size={squareSize}
        color={Konva.Util.getRandomColor()}
        offsetX={(i % squarePerRow) * squareSize}
        offsetY={Math.floor(i / squarePerRow) * squareSize}
      />
    )
  }
  return (
    <div
      className="App"
      style={{
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center'
      }}
    >
      <Stage
        width={squareSize * squarePerRow}
        height={squareSize * squarePerRow}
      >
        <Layer>{rows}</Layer>
      </Stage>
    </div>
  )
}

export default App
