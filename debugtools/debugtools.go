package debugtools

import (
  "log"
)

func CheckError(err error) error {
    if err != nil {
        log.Fatalf("Error: %s", err)
    }
    return err
}
