import * as React from 'react'
import { Rect } from 'react-konva'

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

  public render() {
    return (
      <Rect
        x={this.state.offsetX}
        y={this.state.offsetY}
        width={this.state.size}
        height={this.state.size}
        fill={map[this.state.color]}
        strokeWidth={0.001}
        stroke="black"
      />
    )
  }
}

export default Square
