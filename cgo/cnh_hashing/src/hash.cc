#include "cnh_crypto/cn_slow_hash.hpp"
#include "cnh_crypto/cnh_hash.h"

extern "C" {
  void cryptonight_hash(const char* input, char* output, uint32_t len);
}

void cryptonight_hash(const char* input, char* output, uint32_t len) {
  static thread_local cn_pow_hash_v2 ctx2;
  ctx2.hash(input, len, output);
}