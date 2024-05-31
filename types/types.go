package types

import (
  "net/http"
)

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error
