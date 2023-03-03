package main

import (
	"github.com/sk1t0n/php-pds-skeleton/internal/creator"
)

func main() {
	creator := creator.New()
	creator.CreateProjectStructure()
}
