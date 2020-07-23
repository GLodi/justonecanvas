import * as React from 'react'
import { Stage, Layer } from 'react-konva'
import Square from './Square'
import Konva from 'konva'
// eslint-disable-next-line
import { w3cwebsocket } from 'websocket'

const squareAmount = 1600
const squarePerRow = Math.sqrt(squareAmount)
const squareSize = 1
const layersAmount = 2
const startScale = 20
const scaleBy = 1.05

interface IProps {
  stageX?: number
  stageY?: number
  stageScale?: number
  ws?: WebSocket | null
}

interface IState {
  stageX: number
  stageY: number
  stageScale: number
  ws: WebSocket | null
}

class MainStage extends React.Component<IProps, IState> {
  timeout = 250

  public static defaultProps: Partial<IProps> = {
    stageX: 0,
    stageY: 0,
    stageScale: 40
  }

  public state: IState = {
    stageX:
      window.innerWidth / 2 - ((squarePerRow * squareSize) / 2) * startScale,
    stageY:
      window.innerHeight / 2 - ((squarePerRow * squareSize) / 2) * startScale,
    stageScale: startScale,
    ws: null
  }

  public componentDidMount() {
    this.connect()
  }

  public check() {
    const ws = this.state.ws
    if (!ws || ws.readyState === WebSocket.CLOSED) this.connect() //check if websocket instance is closed, if so call `connect` function.
  }

  public connect() {
    const ws = new WebSocket('ws://localhost:8080/api/v1/canvas/ws')
    const that = this // cache the this
    let connectInterval: NodeJS.Timeout

    // websocket onopen event listener
    ws.onopen = () => {
      console.log('connected websocket main component')

      this.setState({ ws: ws })

      that.timeout = 250 // reset timer to 250 on open of websocket connection
      clearTimeout(connectInterval) // clear Interval on on open of websocket connection
    }

    // websocket onclose event listener
    ws.onclose = e => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
          10000 / 1000,
          (that.timeout + that.timeout) / 1000
        )} second.`,
        e.reason
      )

      that.timeout = that.timeout + that.timeout //increment retry interval
      connectInterval = setTimeout(
        () => this.check(),
        Math.min(10000, that.timeout)
      ) //call check function after timeout
    }

    // websocket onerror event listener
    ws.onerror = err => {
      console.error('Socket encountered error: ', err.target, 'Closing socket')
      ws.close()
    }
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
