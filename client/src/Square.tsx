import * as React from 'react'
import { Rect } from 'react-konva'
import Konva from 'konva'
// eslint-disable-next-line
import { w3cwebsocket } from 'websocket'

interface IProps {
  color?: string
  offsetX?: number
  offsetY?: number
  size?: number
  ws?: WebSocket | null
}

interface IState {
  color: string
  offsetX: number
  offsetY: number
  size: number
}

class Square extends React.Component<IProps, IState> {
  public static defaultProps: Partial<IProps> = {
    color: 'blue',
    offsetX: 0,
    offsetY: 0,
    size: 40,
    ws: null
  }

  public state: IState = {
    color: this.props.color!,
    offsetX: this.props.offsetX!,
    offsetY: this.props.offsetY!,
    size: this.props.size!
  }

  public send(color: string) {
    try {
      const ws = this.props.ws // websocket instance passed as props to the child component.
      if (ws != null) {
        console.log('sto inviando') // catch error
        ws.send(color) //send data to the server
      }
    } catch (error) {
      console.log(error) // catch error
    }
  }

  public changeColor = () => {
    const color: string = Konva.Util.getRandomColor()
    this.send(color)
    this.setState({ color })
    // send new color to ws

    /* const countBy: number = this.props.countBy!
     * const count = this.state.count + countBy
     * this.setState({ count }) */
  }

  public render() {
    return (
      <Rect
        x={this.state.offsetX}
        y={this.state.offsetY}
        width={this.state.size}
        height={this.state.size}
        fill={this.state.color}
        onClick={this.changeColor}
      />
    )
  }
}

export default Square
