package util

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func PrereqMet(a []string, c []string) bool {
  if len(a) == 0 {
    return true
  } 

  f := 0

  for _, i := range a {
    for _, j := range c {
      if i == j {
        f++
      }
    }
  }

  return len(a) == f 

}
