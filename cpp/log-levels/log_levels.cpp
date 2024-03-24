#include <string>

namespace log_line {
  std::string message(std::string line) {
    size_t colonPos{line.find(":")};
    return line.substr(colonPos + 2);
  }

  std::string log_level(std::string line) {
    size_t leftBracketPos{line.find("[")};
    size_t rightBracketPos{line.find("]")};
    return line.substr(leftBracketPos + 1,
                       rightBracketPos - leftBracketPos - 1);
  }

  std::string reformat(std::string line) {
    return message(line) + " (" + log_level(line) + ")";
  }
} // namespace log_line
