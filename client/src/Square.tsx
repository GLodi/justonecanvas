import * as React from 'react'
import { Rect } from 'react-konva'
// eslint-disable-next-line
import { w3cwebsocket } from 'websocket'
import { Constants } from './constants'

const map = [
  'white',
  'blue',
  'lime',
  'red',
  'orange',
  'yellow',
  'purple',
  'fuchsia',
  'black',
  'teal',
  'cyan',
  'gray',
  'green',
  'pink',
  'navy',
  'chocolate'
]

interface IProps {
  color?: number
  offsetX?: number
  offsetY?: number
  size?: number
  ws?: WebSocket | null
  index?: number
}

interface IState {
  color: number
  offsetX: number
  offsetY: number
  size: number
  index: number
}

class Square extends React.Component<IProps, IState> {
  public static defaultProps: Partial<IProps> = {
    color: 0,
    offsetX: 0,
    offsetY: 0,
    size: 40,
    ws: null,
    index: 0
  }

  public state: IState = {
    color: this.props.color!,
    offsetX: this.props.offsetX!,
    offsetY: this.props.offsetY!,
    size: this.props.size!,
    index: this.props.index!
  }

  public componentWillReceiveProps(nextProps: IProps) {
    this.setState({ color: nextProps.color! })
  }

  public send(colorIndex: number) {
    try {
      const ws = this.props.ws
      const data = Uint8Array.from([
        colorIndex,
        Math.floor(this.state.index / Constants.SQUARE_PER_ROW),
        this.state.index % Constants.SQUARE_PER_ROW
      ])
      console.log('sending: ', data)
      if (ws != null) {
        ws.send(data)
      }
    } catch (error) {
      console.log(error)
    }
  }

  public changeColor = () => {
    const color: number = Math.floor(
      Math.random() * (Constants.COLOR_AMOUNT + 1)
    )
    this.send(color)
    /* this.setState({ color }) */
  }

  public render() {
    return (
      <Rect
        x={this.state.offsetX}
        y={this.state.offsetY}
        width={this.state.size}
        height={this.state.size}
        fill={map[this.state.color]}
        onClick={this.changeColor}
      />
    )
  }
}

export default Square
