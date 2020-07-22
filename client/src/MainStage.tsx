import * as React from 'react'
import { Stage, Layer } from 'react-konva'
import Square from './Square'
import Konva from 'konva'

const squareAmount = 1600
const squarePerRow = Math.sqrt(squareAmount)
const squareSize = 20
const layersAmount = 2

interface IProps {
  stageX?: number
  stageY?: number
  stageScale?: number
}

interface IState {
  stageX: number
  stageY: number
  stageScale: number
}

class MainStage extends React.Component<IProps, IState> {
  public static defaultProps: Partial<IProps> = {
    stageX: 0,
    stageY: 0,
    stageScale: 40
  }

  public state: IState = {
    stageX: window.innerWidth / 2 - (squarePerRow * squareSize) / 2,
    stageY: window.innerHeight / 2 - (squarePerRow * squareSize) / 2,
    stageScale: 1
  }

  public render() {
    var rows = []
    for (var i = 0; i < squareAmount; i++) {
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

    var layers = []
    i = 0
    for (var j = 0; j < layersAmount; j++) {
      var layerRows = []
      for (var k = 0; k < squareAmount / layersAmount; k++) {
        layerRows.push(rows[i])
        i++
      }
      layers.push(<Layer key={j}>{layerRows}</Layer>)
    }

    return (
      <Stage
        onWheel={e => {
          e.evt.preventDefault()
          const scaleBy = 1.1
          const stage = e.target.getStage()
          if (stage != null) {
            const oldScale = stage?.scaleX()
            const mousePointTo = {
              x:
                stage.getPointerPosition()!.x / oldScale - stage.x() / oldScale,
              y: stage.getPointerPosition()!.y / oldScale - stage.y() / oldScale
            }

            const newScale =
              e.evt.deltaY > 0 ? oldScale * scaleBy : oldScale / scaleBy

            this.setState({
              stageScale: newScale,
              stageX:
                -(mousePointTo.x - stage.getPointerPosition()!.x / newScale) *
                newScale,
              stageY:
                -(mousePointTo.y - stage.getPointerPosition()!.y / newScale) *
                newScale
            })
          }
        }}
        scaleX={this.state.stageScale}
        scaleY={this.state.stageScale}
        draggable={true}
        width={window.innerWidth}
        height={window.innerHeight}
        x={this.state.stageX}
        y={this.state.stageY}
      >
        {layers}
      </Stage>
    )
  }
}

export default MainStage
