export class Constants {
  public static readonly SQUARE_PER_ROW: number = 65

  public static readonly SQUARE_AMOUNT: number = Math.pow(
    Constants.SQUARE_PER_ROW,
    2
  )

  // Amount of pixels for each square
  public static readonly SQUARE_SIZE: number = 1

  // Number of react-konva layers
  public static readonly LAYERS_AMOUNT: number = 4

  // React-konva grid starting scale
  public static readonly START_SCALE: number = 12

  // Mouse-wheel scroll speed
  public static readonly SCALE_BY: number = 1.05

  public static readonly COLOR_AMOUNT: number = 16
}
