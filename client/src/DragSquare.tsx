import * as React from 'react'
import Konva from 'konva'
import { Rect } from 'react-konva'
import { constants } from 'buffer'
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

class DragSquare extends React.Component<IProps, IState> {
  public rect: Konva.Rect | null = null

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

  public boh = () => {
    if (this.rect != null) {
      const x: number = Math.round(this.rect.getPosition().x)
      const y: number = Math.round(this.rect.getPosition().y)
      this.rect.to({
        x: this.state.offsetX,
        y: this.state.offsetY,
        duration: 0
      })

      if (
        x >= 0 &&
        x < Constants.SQUARE_PER_ROW &&
        y >= 0 &&
        y < Constants.SQUARE_PER_ROW
      ) {
        try {
          const ws = this.props.ws
          const data = Uint8Array.from([this.state.color, y, x])
          console.log('sending: ', data)
          if (ws != null) {
            ws.send(data)
          }
        } catch (error) {
          console.log(error)
        }
      }
    }
  }

  public render() {
    return (
      <Rect
        ref={node => {
          this.rect = node
        }}
        x={this.state.offsetX}
        y={this.state.offsetY}
        width={this.state.size}
        height={this.state.size}
        fill={map[this.state.color]}
        onDragEnd={this.boh.bind(this)}
        draggable={true}
      />
    )
  }
}

export default DragSquare
