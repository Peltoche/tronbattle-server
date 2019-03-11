package game

import (
	"math/rand"

	"github.com/tronbattle/server/model"
)

type Map struct {
	inner [model.MapSize * model.MapSize]model.Element
}

func NewMap() *Map {
	gameMap := Map{
		inner: [model.MapSize * model.MapSize]model.Element{},
	}

	// Add the walls all around the map.
	for i := 0; i < model.MapSize; i++ {
		gameMap.inner[i] = model.Element{Kind: model.Wall}
		gameMap.inner[model.MapSize*(model.MapSize-1)+i] = model.Element{Kind: model.Wall}

		gameMap.inner[i*model.MapSize] = model.Element{Kind: model.Wall}
		gameMap.inner[i*model.MapSize+(model.MapSize-1)] = model.Element{Kind: model.Wall}
	}

	// Add a user for the tests.
	gameMap.inner[10*10] = model.Element{Kind: model.User}

	return &gameMap
}

func (t *Map) SelectRandomStartPoint() model.Position {
	for {
		position := rand.Int() % (model.MapSize * model.MapSize)

		if t.inner[position].Kind == model.Empty {
			return model.Position{
				X: position / model.MapSize,
				Y: position % model.MapSize,
			}
		}
	}
}
