#include <set>
#include <algorithm>
#include <vector>
using namespace std;
#include <stdio.h>
#include <time.h>

vector<float> Bench(vector<int> data) {
  vector<float> times(5);
  set<int> s;
  int start = clock();
  for (int i = 0; i < data.size(); i++) {
    s.insert(data[i]);
  }
  times[0] = (float)(clock() - start) / CLOCKS_PER_SEC;

  start = clock();
  for (int i = 0; i < data.size(); i++) {
    s.insert(data[i]);
  }
  times[1] = (float)(clock() - start) / CLOCKS_PER_SEC;

  start = clock();
  for (int i = 0; i < data.size()/2; i++) {
    s.erase(data[i]);
  }
  times[2] = (float)(clock() - start) / CLOCKS_PER_SEC;

  start = clock();
  for (int i = 0; i < data.size()/2; i++) {
    s.erase(data[i]);
  }
  times[3] = (float)(clock() - start) / CLOCKS_PER_SEC;

  start = clock();
  int count = 0;
  for (int i = 0; i < data.size(); i++) {
    if (s.count(data[i])) {
      count++;
    }
  }
  times[4] = (float)(clock() - start) / CLOCKS_PER_SEC;
  if (count != data.size()/2) {
    printf("Test failed! %d %d\n", count, (int)data.size()/2);
  }
  return times;
}

int main(int argc, char **argv) {
  int N = 100000;
  int R = 5;
  if (argc != 1) {
    if (argc != 3) {
      printf("usage: c_bench [size runs]\n");
      return 0;
    }
    N = atoi(argv[1]);
    R = atoi(argv[2]);
  }

  vector<int> data(N);
  for (int i = 0; i < N; i++) {
    data[i] = i;
  }
  random_shuffle(data.begin(), data.end());
  vector<float> total(5);
  for (int i = 0; i < R; i++) {
    vector<float> times = Bench(data);
    for (int j = 0; j < total.size(); j++) {
      total[j] += times[j];
    }
  }
  for (int i = 0; i < total.size(); i++) {
    total[i] /= R;
  }
  printf("Using input size %d and averaged over %d runs.\n", N, R);
  printf("%3.3f:\t%d\tUnique Inserts\n", total[0], N);
  printf("%3.3f:\t%d\tRepeated Inserts\n", total[1], N);
  printf("%3.3f:\t%d\tUnique Deletes\n", total[2], N/2);
  printf("%3.3f:\t%d\tRepeated Deletes\n", total[3], N/2);
  printf("%3.3f:\t%d\tQueries\n", total[4], N);
}
