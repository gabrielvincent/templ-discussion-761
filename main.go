package main

import (
	components "children/components"
	"context"
	"os"
)

func main() {
	ctx := context.Background()
	components.Index().Render(ctx, os.Stdout)
}
