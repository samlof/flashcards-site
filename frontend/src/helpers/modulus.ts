export const modulus = (n: number, mod: number): number => {
  while (n < 0) n += mod;
  while (n > mod - 1) n -= mod;
  return n;
};
