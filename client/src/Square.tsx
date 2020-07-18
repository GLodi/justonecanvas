import * as React from 'react'
import { Rect } from 'react-konva'

interface IProps {
  countBy?: number
}

interface IState {
  count: number
}

class Square extends React.Component<IProps, IState> {
  public static defaultProps: Partial<IProps> = {
    countBy: 1
  }

  public state: IState = {
    count: 0
  }

  public increase = () => {
    const countBy: number = this.props.countBy!
    const count = this.state.count + countBy
    this.setState({ count })
  }

  public render() {
    return <Rect x={20} y={20} width={20} height={20} fill={'blue'} />
  }
}

export default Square
