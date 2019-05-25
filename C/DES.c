#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define true 1
#define false 0

typedef char bool;
typedef unsigned char byte;

byte pc1Table[] = {0,  57, 49, 41, 33, 25, 17, 9,  1,  58, 50, 42, 34, 26, 18,
                   10, 2,  59, 51, 43, 35, 27, 19, 11, 3,  60, 52, 44, 36, 63,
                   55, 47, 39, 31, 23, 15, 7,  62, 54, 46, 38, 30, 22, 14, 6,
                   61, 53, 45, 37, 29, 21, 13, 5,  28, 20, 12, 4};

byte pc2Table[] = {0,  14, 17, 11, 24, 1,  5,  3,  28, 15, 6,  21, 10,
                   23, 19, 12, 4,  26, 8,  16, 7,  27, 20, 13, 2,  41,
                   52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48, 44, 49,
                   39, 56, 34, 53, 46, 42, 50, 36, 29, 32};

byte ipTable[] = {0,  58, 50, 42, 34, 26, 18, 10, 2,  60, 52, 44, 36,
                  28, 20, 12, 4,  62, 54, 46, 38, 30, 22, 14, 6,  64,
                  56, 48, 40, 32, 24, 16, 8,  57, 49, 41, 33, 25, 17,
                  9,  1,  59, 51, 43, 35, 27, 19, 11, 3,  61, 53, 45,
                  37, 29, 21, 13, 5,  63, 55, 47, 39, 31, 23, 15, 7};

byte ipInverseTable[] = {0,  40, 8,  48, 16, 56, 24, 64, 32, 39, 7,  47, 15,
                         55, 23, 63, 31, 38, 6,  46, 14, 54, 22, 62, 30, 37,
                         5,  45, 13, 53, 21, 61, 29, 36, 4,  44, 12, 52, 20,
                         60, 28, 35, 3,  43, 11, 51, 19, 59, 27, 34, 2,  42,
                         10, 50, 18, 58, 26, 33, 1,  41, 9,  49, 17, 57, 25};

byte sBox[8][4][16] = {
    {
        {14, 4, 13, 1, 2, 15, 11, 8, 3, 10, 6, 12, 5, 9, 0, 7},
        {0, 15, 7, 4, 15, 2, 13, 1, 10, 6, 12, 11, 9, 5, 3, 8},
        {4, 1, 14, 8, 13, 6, 2, 11, 15, 12, 9, 7, 3, 10, 5, 0},
        {15, 12, 8, 2, 4, 9, 1, 7, 5, 11, 3, 14, 10, 0, 6, 13},
    },
    {
        {15, 1, 8, 14, 6, 11, 3, 4, 9, 7, 2, 13, 12, 0, 5, 10},
        {3, 13, 4, 7, 15, 2, 8, 14, 12, 0, 1, 10, 6, 9, 11, 5},
        {0, 14, 7, 11, 10, 4, 13, 1, 5, 8, 12, 6, 9, 3, 2, 15},
        {13, 8, 10, 1, 3, 15, 4, 2, 11, 6, 7, 12, 0, 5, 14, 9},
    },
    {
        {10, 0, 9, 14, 6, 3, 15, 5, 1, 13, 12, 7, 11, 4, 2, 8},
        {13, 7, 0, 9, 3, 4, 6, 10, 2, 8, 5, 14, 12, 11, 15, 1},
        {13, 6, 4, 9, 8, 15, 3, 0, 11, 1, 2, 12, 5, 10, 14, 7},
        {1, 10, 13, 0, 6, 9, 8, 7, 4, 15, 14, 3, 11, 5, 2, 12},
    },
    {
        {7, 13, 14, 3, 0, 6, 9, 10, 1, 2, 8, 5, 11, 12, 4, 15},
        {12, 8, 11, 5, 6, 15, 0, 3, 4, 7, 2, 12, 1, 10, 14, 9},
        {10, 6, 9, 0, 12, 11, 7, 13, 15, 1, 3, 14, 5, 2, 8, 4},
        {3, 15, 0, 6, 10, 1, 13, 8, 9, 4, 5, 11, 12, 7, 2, 14},
    },
    {
        {2, 12, 4, 1, 7, 10, 11, 6, 8, 5, 3, 15, 13, 0, 14, 9},
        {14, 11, 2, 12, 4, 7, 13, 1, 5, 0, 15, 10, 3, 9, 8, 6},
        {4, 2, 1, 11, 10, 13, 7, 8, 15, 9, 12, 5, 6, 3, 0, 14},
        {11, 8, 12, 7, 1, 14, 2, 13, 6, 15, 0, 9, 10, 4, 5, 3},
    },
    {
        {12, 1, 10, 15, 9, 2, 6, 8, 0, 13, 3, 4, 14, 7, 5, 11},
        {10, 15, 4, 2, 7, 12, 9, 5, 6, 1, 13, 14, 0, 11, 3, 8},
        {9, 14, 15, 5, 2, 8, 12, 3, 7, 0, 4, 10, 1, 13, 11, 6},
        {4, 3, 2, 12, 9, 5, 15, 10, 11, 14, 1, 7, 6, 0, 8, 13},
    },
    {
        {4, 11, 2, 14, 15, 0, 8, 13, 3, 12, 9, 7, 5, 10, 6, 1},
        {13, 0, 11, 7, 4, 9, 1, 10, 14, 3, 5, 12, 2, 15, 8, 6},
        {1, 4, 11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5, 9, 2},
        {6, 11, 13, 8, 1, 4, 10, 7, 9, 5, 0, 15, 14, 2, 3, 12},
    },
    {
        {13, 2, 8, 4, 6, 15, 11, 1, 10, 9, 3, 14, 5, 0, 12, 7},
        {1, 15, 13, 8, 10, 3, 7, 4, 12, 5, 6, 11, 0, 14, 9, 2},
        {7, 11, 4, 1, 9, 12, 14, 2, 0, 6, 10, 13, 15, 3, 5, 8},
        {2, 1, 14, 7, 4, 10, 8, 13, 15, 12, 9, 0, 3, 5, 6, 11},
    },
};

byte pTable[] = {0,  16, 7, 20, 21, 29, 12, 28, 17, 1,  15,
                 23, 26, 5, 18, 31, 10, 2,  8,  24, 14, 32,
                 27, 3,  9, 19, 13, 30, 6,  22, 11, 4,  25};

byte eTable[] = {0,  32, 1,  2,  3,  4,  5,  4,  5,  6,  7,  8,  9,
                 8,  9,  10, 11, 12, 13, 12, 13, 14, 15, 16, 17, 16,
                 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25, 24, 25,
                 26, 27, 28, 29, 28, 29, 30, 31, 32, 1};

// to64 src: [0-7]8 to dst: [1-64]64
void to64(byte *src, byte *dst) {
  for (int i = 0; i < 8; i++) {
    byte cur = src[i];
    for (int j = 0; j < 8; j++) {
      dst[(8 - j) + i * 8] = cur % 2;
      cur >>= 1;
    }
  }
}

// to8 src: [1-64]64 to dst: [0-7]8
void to8(byte *src, byte *dst) {
  for (int i = 1; i <= 8; i++) {
    byte b;
    for (int j = 1; j <= 8; j++) {
      b <<= 1;
      if (src[(i - 1) * 8 + j] == 1) b |= 1;
    }
    dst[i - 1] = b;
  }
}

// 置换
void transform(byte *src, int sSize, int dSize, byte *table) {
  byte copy[sSize + 1];
  memcpy(copy, src, (sSize + 1) * sizeof(byte));
  memset(src, 0, dSize * sizeof(byte));
  for (int i = 1; i <= dSize; i++) {
    src[i] = copy[table[i]];
  }
}

// leftShift 循环左移
void leftShift(byte *src) {
  byte hc = src[1], hd = src[29];
  for (int j = 1; j < 28; j++) {
    src[j] = src[j + 1];
    src[j + 28] = src[j + 29];
  }
  src[28] = hc;
  src[56] = hd;
}

// makeKey 生成密钥
void makeKey(byte *key, byte keys[17][49]) {
  transform(key, 64, 56, pc1Table);
  for (int i = 1; i <= 16; i++) {
    if (i != 1 && i != 2 && i != 9 && i != 16) leftShift(key);
    leftShift(key);
    byte copy[65];
    memcpy(copy, key, 65 * sizeof(byte));
    transform(key, 64, 48, pc2Table);
    for (int j = 1; j <= 48; j++) keys[i][j] = key[j];
    memcpy(key, copy, 65 * sizeof(byte));
  }
}

// 补全
void PKCS5Padding(byte *src, byte *dst) {
  int len = strlen(src);
  int count = 8 - len % 8;
  memcpy(dst, src, len * sizeof(byte));
  for (int i = 0; i < count; i++) {
    dst[len + i] = count;
  }
}

// S盒转换
void sBoxTransform(byte *dst, byte *src, int i) {
  int begin = 6 * i - 5, n = src[begin] * 2 + src[begin + 5],
      m = src[begin + 1] * 8 + src[begin + 2] * 4 + src[begin + 3] * 2 +
          src[begin + 4],
      s = sBox[i - 1][n][m];
  for (int j = 0; j < 4; j++) {
    dst[4 - j + (i - 1) * 4] = s % 2;
    s >>= 1;
  }
}

// feistel轮函数
void feistel(byte *dst, byte *src, byte key[49]) {
  // E - 扩展 Test OK
  memcpy(dst, src, 33 * sizeof(byte));
  transform(dst, 48, 48, eTable);
  // 异或
  for (int i = 1; i <= 48; i++) {
    dst[i] ^= key[i];
  }
  byte copy[49];
  memcpy(copy, dst, 49 * sizeof(byte));
  // s-Box
  for (int i = 1; i <= 8; i++) {
    sBoxTransform(dst, copy, i);
  }
  // P - 置换
  transform(dst, 32, 32, pTable);
}

// T - 迭代
void tIteration(byte *src, byte keys[17][49], bool isEncrypt) {
  for (int i = 1; i <= 16; i++) {
    byte newRaw[65], right[33], f[33];
    for (int j = 1; j <= 32; j++) {
      newRaw[j] = src[32 + j];
      right[j] = src[j];
    }
    if (isEncrypt) {
      feistel(f, newRaw, keys[17 - i]);
    } else {
      feistel(f, newRaw, keys[i]);
    }
    for (int j = 1; j <= 32; j++) {
      right[j] ^= f[j];
    }
    for (int j = 1; j <= 32; j++) newRaw[32 + j] = right[j];
    memcpy(src, newRaw, 65 * sizeof(byte));
  }
  // W-置换
  byte copy[65];
  memcpy(copy, src, 65 * sizeof(byte));
  for (int j = 1; j <= 32; j++) {
    src[j] = copy[j + 32];
    src[j + 32] = copy[j];
  }
}

// DES算法主函数
byte *DES(byte *src, byte *key, bool isEncrypt, int *len) {
  if (strlen(key) != 8) {
    printf("Error key length");
    return src;
  }
  if (isEncrypt) {
    *len = strlen(src) + 8 - strlen(src) % 8;
  } else {
    *len = strlen(src);
  }
  byte *dst = malloc(*len * sizeof(byte));
  memset(dst, 0, *len * sizeof(byte));
  // 生成16个子密钥
  byte keys[17][49];
  byte keyBit[65];
  to64(key, keyBit);
  makeKey(keyBit, keys);
  // 补全
  if (isEncrypt)
    PKCS5Padding(src, dst);
  else
    memcpy(dst, src, *len * sizeof(byte));
  int times = strlen(dst) / 8;
  for (int i = 0; i < times; i++) {
    // 分割64位
    byte raw[8];
    for (int j = 0; j < 8; j++) raw[j] = dst[i * 8 + j];
    byte bit[65];
    to64(raw, bit);
    // IP 变换
    transform(bit, 64, 64, ipTable);
    // T-迭代 && W-置换
    tIteration(bit, keys, !isEncrypt);
    // IP逆置换
    transform(bit, 64, 64, ipInverseTable);
    to8(bit, raw);
    for (int j = 0; j < 8; j++) dst[i * 8 + j] = raw[j];
  }
  return dst;
}

// 加密
byte *encrypt(byte *plain, byte *key, int *len) {
  return DES(plain, key, true, len);
}

// 解密
byte *decrypt(byte *cipher, byte *key, int len) {
  int t = 0;
  byte *plain = DES(cipher, key, false, &t);
  // 删除PKCS5填充
  plain[len - plain[len - 1]] = '\0';
  return plain;
}

// 测试
int main() {
  byte key[10] = "megashow",
       plain[100] = "hello,world!";
  int len = 0;
  byte *cipher = encrypt(plain, key, &len);
  byte *mPlain = decrypt(cipher, key, len);
  printf("plain: %s\n", plain);
  printf("cipher: ");
  for (int i = 0; i < len; i++) {
    printf("%x", cipher[i]);
  }
  printf("\n");
  printf("mPlain: %s\n", mPlain);
  free(cipher);
  free(mPlain);
  return 0;
}
