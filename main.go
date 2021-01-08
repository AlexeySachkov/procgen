package main

import(
  "math/rand"
  "golang.org/x/image/colornames"
  // TODO: this package seems to be a simple wrapper above another one,
  // probably makes sense to get rid of it
  "github.com/h8gi/canvas")

const (
  EMPTY = iota
  HOUSE = iota
  FARM = iota
  MARKET = iota
  STORAGE = iota
  WINDMILL = iota
  CASTLE = iota
)

type MapCell int

type Surroundings struct {
  Empty int
  Houses int
  Farms int
  Markets int
  Storages int
  Windmills int
  Castles int
}

const (
  H = 64
  W = 64
)

type MapType [H][W]MapCell

// TODO: allow changing the seed
func RandomChoice(threshold int) bool {
  probability := rand.Intn(100) + 1;
  return probability >= threshold
}

// TODO: eliminate this map copy
func CalculateSurroundings(Y int, X int, Map MapType) Surroundings {
  dx := []int{-1, 0, 1, 1, 1, 0, -1, -1}
  dy := []int{-1, -1, -1, 0, 1, 1, 1, 0}

  Result := Surroundings{}

  for i := 0; i < 8; i++ {
    x := X + dx[i]
    y := Y + dy[i]
    if x < 0 || x >= W {
      continue
    }
    if y < 0 || y >= H {
      continue
    }
    switch c := Map[y][x]; c {
      case EMPTY: Result.Empty++
      case HOUSE: Result.Houses++
      case FARM: Result.Farms++
      case MARKET: Result.Markets++
      case WINDMILL: Result.Windmills++
      case CASTLE: Result.Castles++
    }
  }

  return Result
}

func Update(Map MapType) MapType {
  NewMap := Map

  for curY := 0; curY < H; curY++ {
    for curX := 0; curX < W; curX++ {
      S := CalculateSurroundings(curY, curX, Map)

      if Map[curY][curX] == EMPTY {
        if S.Houses < 3 && RandomChoice(95 - S.Houses) {
          NewMap[curY][curX] = HOUSE
        } else if S.Farms > 1 && RandomChoice(100 - S.Farms * 2 + S.Windmills + S.Houses - 1) {
          NewMap[curY][curX] = WINDMILL
        } else if S.Houses > 1 && RandomChoice(100 - S.Houses * 10) {
          NewMap[curY][curX] = FARM
        }
      }
    }
  }

  return NewMap
}

func main() {
  c := canvas.NewCanvas(&canvas.CanvasConfig{
    Width: 640,
    Height: 640,
    FrameRate: 30,
    Title: "Procedural generation using cellular automata",
  });

  Map := MapType{}

  tick := 0
  c.Draw(func(ctx *canvas.Context) {
    PW := float64(ctx.Width() / W)
    PH := float64(ctx.Height() / H)
    tick++
    if tick % 30 == 0 {
      Map = Update(Map)
    }

    for y := 0; y < H; y++ {
      for x := 0; x < W; x++ {
        // FIXME: assign colors to the rest of cell types
        if Map[y][x] == EMPTY {
          ctx.SetColor(colornames.White)
        } else if Map[y][x] == HOUSE {
          ctx.SetColor(colornames.Brown)
        } else if Map[y][x] == FARM {
          ctx.SetColor(colornames.Green)
        } else {
          ctx.SetColor(colornames.Red)
        }

        ctx.DrawRectangle(PW * float64(x), float64(ctx.Height()) - PH * float64(y + 1), PW, PH)
        ctx.Fill()
      }
    }
  });

}
