export class Constants {
  // Total amount of squares in grid. Must be square.
  public static readonly SQUARE_AMOUNT: number = 1600

  public static readonly SQUARE_PER_ROW: number = Math.sqrt(
    Constants.SQUARE_AMOUNT
  )

  // Amount of pixels for each square
  public static readonly SQUARE_SIZE: number = 1

  // Number of react-konva layers
  public static readonly LAYERS_AMOUNT: number = 2

  // React-konva grid starting scale
  public static readonly START_SCALE: number = 20

  // Mouse-wheel scroll speed
  public static readonly SCALE_BY: number = 1.05

  public static readonly COLOR_AMOUNT: number = 16
}
