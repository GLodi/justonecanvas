import * as React from 'react'
import { Rect } from 'react-konva'
import Konva from 'konva'

interface IProps {
  color?: string
  offsetX?: number
  offsetY?: number
  size?: number
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
    size: 40
  }

  public state: IState = {
    color: this.props.color!,
    offsetX: this.props.offsetX!,
    offsetY: this.props.offsetY!,
    size: this.props.size!
  }

  public changeColor = () => {
    const color = Konva.Util.getRandomColor()
    this.setState({ color })
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
