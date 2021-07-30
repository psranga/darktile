package main

import (
	"os"

	"github.com/liamg/darktile/internal/app/darktile/cmd"
)

/**
                                                                        ▒▒▒▒░░
                                                                      ▓▓▓▓▒▒▓▓▓▓░░
                                                                    ██▓▓▓▓▒▒▓▓▓▓▓▓
                                                                  ▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                                                  ██▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                                                  ▓▓▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓
                                                                ▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓████░░
                                                                ▓▓▒▒▒▒▒▒▒▒▒▒▓▓▓▓░░▓▓░░
                                                                ██▓▓▒▒▒▒▒▒▒▒▓▓▓▓░░▒▒░░
                                                                ██▓▓▓▓▓▓▓▓▓▓▓▓██░░░░░░
                                                                ████▓▓▓▓▓▓▓▓██░░░░▒▒░░░░
                                                                ████████▓▓████░░░░░░░░░░
                                                                ██████████████▒▒░░  ░░░░░░░░
                                                                ██▓▓▓▓████████    ░░░░░░░░▒▒░░
                                                                ▓▓▓▓▓▓████████        ░░▒▒▒▒██░░
                                                                ▓▓▓▓▓▓████████              ▒▒
                                                                ▓▓▓▓▓▓████████
                                                                ▓▓▓▓██████████
                                                                ██    ▓▓
                                                                ░░  ░░░░  ░░░░
                                                                ░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                                                ░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                                                ░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
                                                                ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓░░
                                                    ░░░░▒▒░░▒▒░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                            ░░░░░░░░░░▒▒▒▒░░░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                        ░░░░░░░░░░░░░░░░▒▒▒▒░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                                ░░░░░░░░░░░░░░▒▒░░░░▒▒▒▒▒▒▒▒░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
                          ░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒░░▒▒▒▒▒▒▒▒▒▒░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
                  ░░░░░░░░░░▒▒▓▓▓▓▒▒▓▓▓▓▓▓▓▓▒▒▒▒▒▒▒▒░░░░▒▒▒▒░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
            ░░  ░░▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▒▒░░░░░░░░░░░░░░░░░░▒▒░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
      ▓▓██░░░░░░▒▒▓▓▓▓▓▓▓▓▒▒▓▓▓▓░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
  ▓▓  ██▒▒░░▒▒▒▒▓▓▓▓████▓▓░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
  ██  ██▓▓▒▒▓▓▓▓████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
  ██▒▒████▒▒▓▓▓▓████▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒░░▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓██▓▓
  ▓▓▓▓██████▓▓████▒▒░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓████
    ▓▓██████      ░░  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▓▓▓▓██▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓
      ▒▒▓▓▓▓          ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓▓▓▓▓▓██▓▓████▓▓▓▓████████▓▓▓▓██▓▓▒▒
    ▒▒████░░        ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓▓▓▓▓▓████▓▓████████████████▓▓████
        ▓▓░░    ░░      ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▓▓▓▓████▓▓████████████████▓▓▓▓▓▓
                ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓▓▓▓████████████████████████▓▓▓▓
              ░░░░░░░░  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▓▓▓▓████████████████████████▓▓
                        ░░░░░░░░░░░░  ░░░░░░░░░░░░░░░░░░▒▒▒▒██████████████████████████▓▓
                ░░    ░░  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▓▓████████████████████████
                  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▓▓██████████████████████
                    ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░▒▒▒▒▒▒▓▓▓▓████████████████▓▓
                        ░░░░░░░░▒▒░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓██████▒▒
                            ░░░░▒▒▒▒░░▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓▓▓▓▓▒▒
                                  ▒▒░░▓▓▓▓▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓▓▓▓▓░░
                                    ░░▓▓    ░░░░  ▒▒░░▓▓▓▓
                                    ░░▓▓              ▓▓▓▓
                                    ░░▓▓              ▓▓▓▓
                                    ▒▒▒▒              ▒▒▒▒
                                    ░░░░              ▒▒▒▒░░
                                    ░░░░            ▒▒▒▒▒▒▒▒░░░░▒▒▒▒▒▒▒▒
                                      ░░░░              ░░▓▓▒▒░░░░▒▒▒▒▒▒
                                    ░░░░▒▒              ▒▒░░▒▒▒▒▒▒░░░░▒▒
                                      ░░▒▒░░▒▒▒▒▒▒░░▒▒▒▒  ░░▒▒░░
                                      ░░░░▒▒░░░░░░▒▒▒▒▒▒▒▒
                                        ░░░░▒▒▒▒▒▒░░░░░░▒▒
                                          ░░░░▒▒▒▒▒▒▒▒░░░░▓▓

*/

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}