import * as React from 'react'
import { Stage, Layer } from 'react-konva'
import Square from './Square'
import DragSquare from './DragSquare'
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
  grid: number[][]
}

class MainStage extends React.Component<IProps, IState> {
  resturl = ''
  wsurl = ''
  timeout = 250
  references = Array(Constants.SQUARE_AMOUNT)
    .fill(0)
    .map(() => React.createRef<Square>())
  referencesDrag = Array(Constants.COLOR_AMOUNT)
    .fill(0)
    .map(() => React.createRef<DragSquare>())

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
    ws: null,
    grid: this.makegrid()
  }

  private getOrCreateRef(id: number): React.RefObject<Square> {
    if (!this.references.hasOwnProperty(id)) {
      this.references[id] = React.createRef<Square>()
    }
    return this.references[id]
  }

  private getOrCreateRefDrag(id: number): React.RefObject<DragSquare> {
    if (!this.referencesDrag.hasOwnProperty(id)) {
      this.referencesDrag[id] = React.createRef<DragSquare>()
    }
    return this.referencesDrag[id]
  }

  private makegrid(): number[][] {
    const grid: number[][] = new Array(Constants.SQUARE_PER_ROW)
      .fill(0)
      .map(() => new Array(Constants.SQUARE_PER_ROW).fill(0))
    for (var i = 0; i < Constants.SQUARE_AMOUNT; i++) {
      const x = (i % Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      const y = Math.floor(i / Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      /* grid[y][x] = Math.floor(Math.random() * (Constants.COLOR_AMOUNT + 1)) */
      grid[y][x] = 0
    }
    return grid
  }

  private connect() {
    fetch(this.resturl)
      .then(res => res.json())
      .then(data => {
        const grid2: number[][] = new Array(Constants.SQUARE_PER_ROW)
          .fill(0)
          .map(() => new Array(Constants.SQUARE_PER_ROW).fill(0))
        for (var i = 0; i < Constants.SQUARE_AMOUNT; i++) {
          const x = (i % Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
          const y =
            Math.floor(i / Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
          grid2[y][x] = data['cells'][i]['color']
        }
        this.setState({ grid: grid2 })
      })

    const ws = new WebSocket(this.wsurl)
    ws.binaryType = 'arraybuffer'
    const that = this // cache the this
    let connectInterval: NodeJS.Timeout

    // websocket onopen event listener
    ws.onopen = () => {
      console.log('connected websocket main component')

      for (var i = 0; i < Constants.COLOR_AMOUNT; i++) {
        this.referencesDrag[i].current!.setState({
          ws: ws
        })
      }

      that.timeout = 250 // reset timer to 250 on open of websocket connection
      clearTimeout(connectInterval) // clear Interval on on open of websocket connection
    }

    ws.onmessage = evt => {
      const buf = new Uint8Array(evt.data).buffer
      const data = new DataView(buf)
      const color: number = data.getUint8(0)
      const y: number = data.getUint8(1)
      const x: number = data.getUint8(2)
      console.log('received: ', color, y, x)
      // use reference to setstate in children
      // eslint-disable-next-line
      this.state.grid[y][x] = color
      this.references[y * Constants.SQUARE_PER_ROW + x].current!.setState({
        color: color
      })
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

  private check() {
    const ws = this.state.ws
    if (!ws || ws.readyState === WebSocket.CLOSED) this.connect() //check if websocket instance is closed, if so call `connect` function.
  }

  public componentDidMount() {
    if (!process.env.NODE_ENV || process.env.NODE_ENV === 'development') {
      console.log('DEV ENV')
      this.resturl = '/api/v1/canvas'
      this.wsurl = 'ws://localhost:8080/api/v1/canvas/ws'
    } else {
      this.resturl = 'https://justonecanvas.live/api/v1/canvas'
      this.wsurl = 'wss://justonecanvas.live/api/v1/canvas/ws'
    }
    this.connect()
  }

  public render() {
    var rows = []
    for (var i = 0; i < Constants.SQUARE_AMOUNT; i++) {
      const x = (i % Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      const y = Math.floor(i / Constants.SQUARE_PER_ROW) * Constants.SQUARE_SIZE
      rows.push(
        <Square
          ref={this.getOrCreateRef(i)}
          key={i}
          index={i}
          size={Constants.SQUARE_SIZE}
          color={this.state.grid[y][x]}
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

    var colors = []
    for (k = 0; k < Constants.COLOR_AMOUNT; k++) {
      const x = (Constants.SQUARE_PER_ROW + 2) * Constants.SQUARE_SIZE
      const y = k * 2 * Constants.SQUARE_SIZE
      colors.push(
        <DragSquare
          ref={this.getOrCreateRefDrag(k)}
          key={k}
          index={k}
          size={Constants.SQUARE_SIZE}
          color={k}
          offsetX={x}
          offsetY={y}
        />
      )
    }
    layers.push(<Layer key={Constants.LAYERS_AMOUNT}>{colors}</Layer>)

    return (
      <div>
        <Stage
          onWheel={e => {
            e.evt.preventDefault()
            const stage = e.target.getStage()
            if (stage != null) {
              const oldScale = stage?.scaleX()
              const mousePointTo = {
                x:
                  stage.getPointerPosition()!.x / oldScale -
                  stage.x() / oldScale,
                y:
                  stage.getPointerPosition()!.y / oldScale -
                  stage.y() / oldScale
              }

              const newScale =
                e.evt.deltaY > 0
                  ? oldScale * Constants.SCALE_BY
                  : oldScale / Constants.SCALE_BY

              this.setState({
                grid: this.state.grid,
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
          width={window.innerWidth}
          height={window.innerHeight}
          draggable={true}
          x={this.state.stageX}
          y={this.state.stageY}
        >
          {layers}
        </Stage>
      </div>
    )
  }
}

export default MainStage
