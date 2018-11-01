#include <bits/stdc++.h>
using namespace std;

int main() {
  string S;
  cin >> S;

  int plus = 1;
  int minus = 0;
  for (int i = 0; i < S.size(); i++) {
    if ((i % 2) == 1) {
      if (S.at(i) == '-')
        minus++;
      else
        plus++;
    }
  }
  cout << plus - minus << endl;
}
