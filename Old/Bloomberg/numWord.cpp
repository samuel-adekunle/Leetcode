#include <cstdint>
#include <cstdlib>
#include <cstdio>
#include <cassert>
#include <cmath>
#include <string>
#include <sstream>

std::string numword(const uint32_t num) {
  if (num == 0) return "Zero";
  std::stringstream res;
  auto units = [](const uint32_t _num) {
    assert(_num > 0);
    const char* _units[] = { "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen" };

    return _num < 10 ? _units[_num - 1] : _units[_num - 2];
  };

  auto tens = [](const uint32_t _num) {
    assert(_num > 0);
    const char* _tens[] = { "Ten", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety" };
    return _tens[_num - 1];
  };

  auto hundreds = [&](const uint32_t _num) {
    const uint32_t _hundreds_digit = _num / 100;
    const char* _hundreds = _hundreds_digit > 0 ? units(_hundreds_digit) : nullptr;

    const uint32_t _tens_digit = (_num - 100 * _hundreds_digit) / 10;
    const char* _tens = _tens_digit > 0 ? tens(_tens_digit) : nullptr;

    const uint32_t _units_digit = (_num - 100 * _hundreds_digit - 10 * _tens_digit);
    const char* _units = _units_digit > 0 ? units(_tens_digit == 1 ? _units_digit + 10 : _units_digit) : nullptr;

    _tens = _tens_digit == 1 && _units_digit > 0 ? nullptr : _tens;

    if (_num > 0 && !res.str().empty()) res << " ";
    if (_hundreds) res << _hundreds << " Hundred";
    if (_hundreds && _tens) res << " ";
    if (_tens) res << _tens;
    if ((_hundreds || _tens) && _units) res << " ";
    if (_units) res << _units;
  };

  const char* _powers_of_ten[] = { "Billion", "Million", "Thousand" };
  const uint32_t _max_pow_ten = 3;
  uint32_t digits_of_ten[_max_pow_ten] = { 0 };

  auto powers_of_ten = [&](const uint32_t _num, const uint32_t _pow_index) {
    hundreds(_num);
    if (_num > 0) res << " " << _powers_of_ten[_pow_index];
  };

  for (uint32_t i = 0; i <= _max_pow_ten; i++) {
    uint32_t _num = num;
    for (uint32_t j = 0; j < i; j++) {
      _num -= digits_of_ten[j] * std::pow(10, (_max_pow_ten - j) * 3);
    }
    if (i == _max_pow_ten) {
      hundreds(_num);
      break;
    }
    const uint32_t _sub_hundred = _num / std::pow(10, (_max_pow_ten - i) * 3);
    powers_of_ten(_sub_hundred, i);
    digits_of_ten[i] = _sub_hundred;
  }
  return res.str();
}


int main(int argc, char const* argv[])
{
  printf("%s\n", numword(1012345).c_str());
  return 0;
}
