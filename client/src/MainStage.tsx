import * as React from 'react'
import { Stage, Layer } from 'react-konva'
import Square from './Square'
// eslint-disable-next-line
import { w3cwebsocket } from 'websocket'
import { Constants } from './constants'

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
  grid: number[][] = new Array(Constants.SQUARE_PER_ROW)
    .fill(0)
    .map(() => new Array(Constants.SQUARE_PER_ROW).fill(0))

  timeout = 250

  public static defaultProps: Partial<IProps> = {
    stageX: 0,
    stageY: 0,
    stageScale: 40
  }

  public state: IState = {
    stageX:
      window.innerWidth / 2 -
      ((Constants.SQUARE_PER_ROW * Constants.SQUARE_SIZE) / 2) *
        Constants.START_SCALE,
    stageY:
      window.innerHeight / 2 -
      ((Constants.SQUARE_PER_ROW * Constants.SQUARE_SIZE) / 2) *
        Constants.START_SCALE,
    stageScale: Constants.START_SCALE,
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
    ws.binaryType = 'arraybuffer'
    const that = this // cache the this
    let connectInterval: NodeJS.Timeout

    // websocket onopen event listener
    ws.onopen = () => {
      console.log('connected websocket main component')

      this.setState({ ws: ws })

      that.timeout = 250 // reset timer to 250 on open of websocket connection
      clearTimeout(connectInterval) // clear Interval on on open of websocket connection
    }

    ws.onmessage = evt => {
      // listen to data sent from the websocket server
      //const message = JSON.parse(evt.data)
      //this.setState({ dataFromServer: message })
      var buf = new Uint8Array(evt.data).buffer
      var data = new DataView(buf)
      const index: number = data.getUint8(0)
      const y: number = data.getUint8(1)
      const x: number = data.getUint8(2)
      console.log('received: ', index, y, x)
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
    for (var i = 0; i < Constants.SQUARE_AMOUNT; i++) {
      const x = (i % Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      const y = Math.floor(i / Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      this.grid[y][x] = Math.floor(Math.random() * (Constants.COLOR_AMOUNT + 1))
      rows.push(
        <Square
          key={i}
          index={i}
          size={Constants.SQUARE_SIZE}
          ws={this.state.ws}
          color={this.grid[y][x]}
          offsetX={x}
          offsetY={y}
        />
      )
    }

    var layers = []
    i = 0
    for (var j = 0; j < Constants.LAYERS_AMOUNT; j++) {
      var layerRows = []
      for (
        var k = 0;
        k < Constants.SQUARE_AMOUNT / Constants.LAYERS_AMOUNT;
        k++
      ) {
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
              e.evt.deltaY > 0
                ? oldScale * Constants.SCALE_BY
                : oldScale / Constants.SCALE_BY

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
