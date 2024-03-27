#include "raindrops.h"
#include <string>

namespace raindrops {
  std::string convert(int n) {
    std::string s{""};
    if (n % 3 == 0) {
      s += "Pling";
    }
    if (n % 5 == 0) {
      s += "Plang";
    }
    if (n % 7 == 0) {
      s += "Plong";
    }

    if (s.length() == 0) {
      return std::to_string(n);
    }

    return s;
  }
} // namespace raindrops
