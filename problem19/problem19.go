package problem19

func CountingSundays() int {
  day, month, year := 7, 1, 1900
  count := 0
  for year <= 2000 {
    if day == 1 && year > 1900 {
      count += 1
    }
    day += 7

    var days int
    if month == 2 {
      // Leap year
      if year % 4 == 0 && (year % 100 != 0 || year % 400 == 0) {
        days = 29
      } else {
        days = 28
      }
    } else if month == 9 || month == 4 || month == 6 || month == 11 {
      days = 30
    } else {
      days = 31
    }
    if day > days {
      day = day % days
      month += 1
      if month > 12 {
        month = 1
        year += 1
      }
    }
  }
  return count
}
