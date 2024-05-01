// INFO: Headers from the standard library should be inserted at the top via
// #include <LIBRARY_NAME>

// daily_rate calculates the daily rate given an hourly rate
#include <cmath>
double daily_rate(double hourly_rate) {
  constexpr int daily_hours = 8;
  return hourly_rate * daily_hours;
}

// apply_discount calculates the price after a discount
double apply_discount(double before_discount, double discount) {
  return before_discount * (1 - discount / 100.0);
}

// monthly_rate calculates the monthly rate, given an hourly rate and a discount
// The returned monthly rate is rounded up to the nearest integer.
int monthly_rate(double hourly_rate, double discount) {
  constexpr int days_of_month = 22;
  double rate =
      apply_discount((daily_rate(hourly_rate) * days_of_month), discount);

  return std::ceil(rate);
}

// days_in_budget calculates the number of workdays given a budget, hourly rate,
// and discount The returned number of days is rounded down (take the floor) to
// the next integer.
int days_in_budget(int budget, double hourly_rate, double discount) {
  const double discounted_hourly_rate = apply_discount(hourly_rate, discount);
  const double discounted_daily_rate = daily_rate(discounted_hourly_rate);
  const double days_in_budget = budget / discounted_daily_rate;
  return std::floor(days_in_budget);
}
